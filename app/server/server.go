package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
	csrfHeader := middleware.CSRF()

	r.Static("/", "static")
	r.Renderer = templateRenderer("templates/*.html", true)
	r.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("APP_SECRET")))))
	{
		account := accountHandler{s: s.Account}
		r.POST("/token/refresh", account.refreshToken, s.jwtConfig(fallback))
		g := r.Group("/accounts")
		g.GET("/login", account.loginPage, csrfForm, account.restricted)
		g.POST("/login", account.loginPerform, csrfForm, account.restricted)
		g.POST("/setting/new", account.registerAccount, s.jwtConfig(fallback), getToken)
		g.POST("/setting/list", account.accountTable, s.jwtConfig(fallback), getToken)
		g.POST("/setting/delete", account.deleteAccount, s.jwtConfig(fallback), getToken)
	}
	{
		hub := NewHub()
		go hub.Run()
		dashboard := dashboardHandler{s: s.Account, logger: s.Logger, hub: hub}
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
		g.GET("/ws", dashboard.engine)
		g.GET("/board/trello", dashboard.boardTrelloPage)
		g.GET("/setting/details", dashboard.settingDetails)
		g.GET("/setting/users", dashboard.settingUsers, csrfHeader, getToken)
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

		return handlerFunc(c)
	}
}

func (s *Server) Start(addr, port string) error {
	return s.router.Start(fmt.Sprintf("%s:%s", addr, port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}
