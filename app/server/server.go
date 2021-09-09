package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/accounts"
	"github.com/hinha/PAM-Trello/app/trello"
)

type Server struct {
	Account accounts.Service
	Trello  trello.Service
	//Inbox   handling.ServiceInbox

	Logger *log.Entry

	router *echo.Echo
}

func New(account accounts.Service, trello trello.Service, logger *log.Entry) *Server {
	s := &Server{
		Account: account,
		Trello:  trello,
		//Inbox:   handlingInbox,

		Logger: logger,
	}

	r := echo.New()
	echo.NotFoundHandler = func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "404.html", nil)
	}
	r.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		errorPage := fmt.Sprintf("%d.html", code)
		if err := c.File(errorPage); err != nil {
			c.Logger().Error(err)
		}
		c.Logger().Error(err)
	}
	csrfForm := middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	})
	csrfHeader := middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength:    32,
		TokenLookup:    "header:" + echo.HeaderXCSRFToken,
		ContextKey:     "csrf",
		CookieName:     "_csrf",
		CookieSecure:   true,
		CookieMaxAge:   86400,
		CookieSameSite: http.SameSiteDefaultMode,
	})

	r.Static("/", "static")
	r.Renderer = templateRenderer("templates/*.html", true)
	{
		{
			account := apiAccountHandler{s: s.Account, logger: s.Logger}
			api := r.Group("/api")

			apiAccount := api.Group("/accounts")
			apiAccount.GET("/login", account.formToken, csrfHeader)
			apiAccount.POST("/login", account.loginPerform)

			gAccountSub1 := apiAccount.Group("/data", s.jwtConfigHeader(fallback), getToken)
			gAccountSub1.GET("", account.profileData)
			gAccountSub1.POST("/refresh", account.refreshToken)

			dashboard := apiDashboardHandler{trello: s.Trello, account: s.Account, logger: s.Logger}
			apiDashboard := api.Group("/dashboard", s.jwtConfigHeader(fallback), getToken, dashboard.verify)
			apiDashboard.GET("/performance", dashboard.performance)
			apiDashboard.GET("/analytic/trello", dashboard.analyticTrelloCard)
			apiDashboard.GET("/settings/user", dashboard.userSetting)
			apiDashboard.POST("/settings/user", dashboard.addUserSetting)
			apiDashboard.PATCH("/settings/user", dashboard.updateUserSetting)
			apiDashboard.DELETE("/settings/user", dashboard.deleteUserSetting)

			apiDashboard.POST("/settings/user/trello", dashboard.trelloUserSetting)
			apiDashboard.POST("/settings/user/role", dashboard.assignRoleAccount)
		}
		{
			account := accountHandler{s: s.Account}
			r.POST("/token/refresh", account.refreshToken, s.jwtConfig(fallback))
			g := r.Group("/accounts")
			g.GET("/login", account.loginPage, csrfForm, account.restricted)
			g.POST("/login", account.loginPerform, csrfForm, account.restricted)
			g.POST("/setting/new", account.registerAccount, s.jwtConfig(fallback), getToken)
			g.POST("/setting/list", account.accountTable, s.jwtConfig(fallback), getToken)
			g.POST("/setting/delete", account.deleteAccount, s.jwtConfig(fallback), getToken)
			g.POST("/setting/role/new", account.assignRoleAccount, s.jwtConfig(fallback), getToken)
		}
	}
	{
		// Deprecated
		//hub := NewHub()
		//go hub.Run()
		//dashboard := dashboardHandler{trello: trello.Account, socket: trello.Inbox, logger: trello.Logger, hub: hub}
		//g := r.Group("/dashboard")
		//g.Use(trello.jwtConfig(func(err error, c echo.Context) error {
		//
		//	if errors.Is(err, middleware.ErrJWTMissing) || errors.Is(err, middleware.ErrJWTInvalid) {
		//		return c.Redirect(http.StatusPermanentRedirect, "/accounts/login")
		//	}
		//
		//	if err.Error() == "Token is expired" || err.Error() == "signature is invalid" {
		//		c.SetCookie(app.DeleteCookie)
		//		return c.Redirect(http.StatusPermanentRedirect, "/accounts/login")
		//	}
		//	return nil
		//}))

		//g.GET("", dashboard.dashboardPage)
		//g.GET("/inbox/ws", dashboard.inboxSocket, dashboard.verify)
		//g.GET("/board/trello", dashboard.boardTrelloPage)
		//g.GET("/setting/details", dashboard.settingDetails)
		//g.GET("/setting/users", dashboard.settingUsers, csrfHeader, getToken)
	}

	r.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{})
	})

	s.router = r
	return s
}

func (s *Server) jwtConfig(callback middleware.JWTErrorHandlerWithContext) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:              []byte(os.Getenv("JWT_SECRET")),
		TokenLookup:             "cookie:token",
		ErrorHandlerWithContext: callback,
	})
}

func (s *Server) jwtConfigHeader(callback middleware.JWTErrorHandlerWithContext) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:              []byte(os.Getenv("JWT_SECRET")),
		ErrorHandlerWithContext: callback,
	})
}

func fallback(err error, c echo.Context) error {
	output := c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	if errors.Is(err, middleware.ErrJWTMissing) || errors.Is(err, middleware.ErrJWTInvalid) {
		c.SetCookie(app.DeleteCookie)
		return output
	}
	if err.Error() == "Token is expired" || err.Error() == "signature is invalid" {
		c.SetCookie(app.DeleteCookie)
		return output
	}

	return nil
}

func getToken(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		c.Set("authorize", claims["status"])
		c.Set("user_id", claims["id"])
		c.Set("jwt", user.Raw)

		return handlerFunc(c)
	}
}

func (s *Server) Start(addr, port string) error {
	s.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowHeaders: []string{
			echo.HeaderAuthorization,
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderXRequestID,
			echo.HeaderXXSSProtection, // start security
			echo.HeaderXFrameOptions,
			echo.HeaderContentSecurityPolicy,
			echo.HeaderContentSecurityPolicyReportOnly,
			echo.HeaderXCSRFToken,
			echo.HeaderReferrerPolicy, // end security
		},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}), session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("APP_SECRET")))))
	return s.router.Start(fmt.Sprintf("%s:%s", addr, port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}
