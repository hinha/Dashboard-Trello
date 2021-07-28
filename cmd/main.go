package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	gorm_logger "gorm.io/gorm/logger"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/accounts"
	"github.com/hinha/PAM-Trello/app/handling"
	"github.com/hinha/PAM-Trello/app/repository"
	"github.com/hinha/PAM-Trello/app/server"
	cron_server "github.com/hinha/PAM-Trello/app/server/cron"
	"github.com/hinha/PAM-Trello/app/trello"
	"github.com/hinha/PAM-Trello/app/util/authority"
)

var rootCmd = &cobra.Command{
	Use:     "run",
	Short:   "Main Dashboard Partitioning Around Medoids",
	Example: "run",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

const defaultPort = "8080"

func main() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	// handle err
	time.Local = loc

	var (
		appSecret = envString("APP_SECRET", "secret")
		dbUser    = envString("DB_USER", "user")
		dbPass    = envString("DB_PASS", "password")
		dbName    = envString("DB_NAME", "admin")
		dbHost    = envString("DB_HOST", "127.0.0.1")
	)

	logger := log.New()
	logger.Formatter = new(log.TextFormatter)
	logger.Formatter.(*log.TextFormatter).DisableColors = true
	logger.Formatter.(*log.TextFormatter).FullTimestamp = true
	logger.Level = log.TraceLevel
	logger.WithField("ts", logger.WithTime(time.Now()))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gorm_logger.Default.LogMode(gorm_logger.Silent),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	var cmdMigrate = &cobra.Command{
		Use:     "migrate",
		Short:   "Migrate Database",
		Example: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Need refactor

			// initiate authority
			err := db.AutoMigrate(&app.Accounts{}, &app.AccountDetail{}, &app.Trello{}, &app.TrelloUserCard{})
			if err != nil {
				logger.Fatal(err)
			}

			password := []byte(fmt.Sprintf("%s:%s", appSecret, "admin"))
			// Hashing the password with the default cost of 10
			hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
			if err != nil {
				logger.Fatal(err)
			}
			t := time.Date(2021, 7, 15, 15, 00, 00, 00, time.UTC)
			UserID := fmt.Sprintf("%d%d", int(t.Unix()/100)%math.MaxInt64, len(hashedPassword))

			// Do nothing on conflict
			db.Clauses(clause.OnConflict{DoNothing: true}).
				Create(&app.Accounts{
					ID:             UserID,
					Username:       "admin",
					Password:       string(hashedPassword),
					SecretPassword: appSecret,
					CreatedAt:      t,
					LastLogin:      t,
				})

			auth := authority.New(authority.Options{
				TablesPrefix: "authority_",
				DB:           db,
			})
			auth.CreatePermission(app.PermReadWrite)
			auth.CreatePermission(app.PermResetPassword)
			auth.CreatePermission(app.PermListEmployee)
			auth.CreatePermission(app.PermAttendance)
			auth.CreatePermission(app.PermUserDetails)

			auth.CreateRole("admin")
			auth.CreatePermission(app.PermUserGroup)
			auth.CreatePermission(app.PermUserProperties)
			auth.CreatePermission(app.PermUserManage)
			auth.CreatePermission(app.PermUserUpdatePassword)
			auth.CreatePermission(app.PermServiceRequest)
			auth.AssignPermissions("admin", []string{
				app.PermUserGroup,
				app.PermUserProperties,
				app.PermUserManage,
				app.PermUserUpdatePassword,
				app.PermServiceRequest,
			})
			auth.AssignRole(UserID, "admin")
			auth.CreateRole("developer")
			auth.AssignPermissions("developer", []string{
				app.PermReadWrite,
				app.PermResetPassword,
				app.PermAttendance,
				app.PermUserDetails,
			})
			auth.CreateRole("c_level")
			auth.AssignPermissions("c_level", []string{
				app.PermReadWrite,
				app.PermResetPassword,
				app.PermAttendance,
				app.PermUserDetails,
				app.PermListEmployee,
			})
		},
	}

	var cmdRun = &cobra.Command{
		Use:     "start",
		Short:   "Running dashboard",
		Example: "start",
		Run: func(cmd *cobra.Command, args []string) {
			// Setup repository
			var authRepo app.AuthRepository
			authRepo = repository.NewAuthRepository(db)

			var accountRepo app.AccountRepository
			accountRepo = repository.NewAccountRepository(db)

			var tR app.TrelloRepository
			tR = repository.NewTrelloRepository(db)

			var ts trello.Service
			ts = trello.New(tR)
			ts = trello.NewLoggingService(logger.WithField("component", "trelloClient"), ts)

			// Handling
			var inbox handling.ServiceInbox
			inbox = handling.NewHandlingInbox(accountRepo)
			go inbox.Run()

			var as accounts.Service
			as = accounts.NewService(authRepo, accountRepo)
			as = accounts.NewLoggingService(logger.WithField("component", "accounts"), as)

			srv := server.New(as, ts, inbox, logger.WithField("component", "http"))

			go func() {
				logger.WithFields(log.Fields{"transport": "http", "address": defaultPort}).Info("listening")
				if err := srv.Start("0.0.0.0", defaultPort); err != nil {
					logger.Fatal(err)
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)
			<-quit

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				logger.Fatal(err)
			}

			logger.Info("terminated")
		},
	}

	var cmdJob = &cobra.Command{
		Use:     "job",
		Short:   "Running scheduler dashboard",
		Example: "job",
		Run: func(cmd *cobra.Command, args []string) {
			var tR app.TrelloRepository
			tR = repository.NewTrelloRepository(db)

			var ts trello.Service
			ts = trello.New(tR)
			ts = trello.NewLoggingService(logger.WithField("component", "trelloClient"), ts)

			worker := cron_server.New(ts, logger.WithField("component", "worker"))
			worker.Start()

			logger.Info("worker is now running.  Press CTRL-C to exit.")
			sc := make(chan os.Signal, 1)
			signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
			<-sc

			worker.Stop()
			logger.Info("Gracefully shutdown the job...")
		},
	}

	rootCmd.AddCommand(cmdMigrate, cmdRun, cmdJob)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
