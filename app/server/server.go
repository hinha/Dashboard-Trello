package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/hinha/PAM-Trello/app"
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
		g := r.Group("/accounts", account.restricted)
		g.GET("/login", account.loginPage)
		g.POST("/login", account.loginPerform)
		g.POST("/token/refresh", account.refreshToken, s.jwtConfig(func(err error, c echo.Context) error {
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
		}))
	}
	{
		dashboard := dashboardHandler{s: s.Account}
		g := r.Group("/dashboard", dashboard.restricted)
		g.Use(s.jwtConfig(func(err error, c echo.Context) error {

			if errors.Is(err, middleware.ErrJWTMissing) || errors.Is(err, middleware.ErrJWTInvalid) {
				return c.Redirect(http.StatusPermanentRedirect, "/accounts/login")
			}

			if err.Error() == "Token is expired" || err.Error() == "signature is invalid" {
				c.SetCookie(app.DeleteCookie)
				return c.Redirect(http.StatusPermanentRedirect, "/accounts/login")
			}
			return nil
		}))
		g.GET("", dashboard.dashboardPage)
	}

	r.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "index.html", nil)
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

func (s *Server) Start(addr, port string) error {
	return s.router.Start(fmt.Sprintf("%s:%s", addr, port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}
