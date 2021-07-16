package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/hinha/PAM-Trello/app/accounts"
)

type Server struct {
	Account accounts.Service

	Logger *log.Entry

	router *echo.Echo
}

func New(account accounts.Service, logger *log.Entry) *Server {
	s := &Server{
		Account: account,

		Logger: logger,
	}

	r := echo.New()
	r.Static("/", "static")
	r.Renderer = templateRenderer("templates/*.html", true)
	r.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("APP_SECRET")))))
	{
		account := accountHandler{s: s.Account}
		g := r.Group("/accounts")
		g.GET("/login", account.loginPage)
		g.POST("/login", account.loginPerform)
	}
	{
		account := accountHandler{}
		g := r.Group("/accounts")
		g.GET("/login2", account.loginPage)
	}
	{
		g := r.Group("/dashboard")
		g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(os.Getenv("JWT_SECRET")),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				if errors.Is(err, middleware.ErrJWTMissing) || errors.Is(err, middleware.ErrJWTInvalid) {
					return c.Redirect(http.StatusPermanentRedirect, "/accounts/login")
				}
				return nil
			},
		}))
		g.GET("/home", func(ctx echo.Context) error {
			return ctx.Render(http.StatusOK, "index.html", nil)
		}, func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				fmt.Println(c.Response().Status)
				return handlerFunc(c)
			}
		})
	}

	r.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "index.html", nil)
	})

	s.router = r
	return s
}

func (s *Server) Start(addr, port string) error {
	return s.router.Start(fmt.Sprintf("%s:%s", addr, port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}
