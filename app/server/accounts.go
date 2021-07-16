package server

import (
	"github.com/gorilla/sessions"
	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/accounts"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type accountHandler struct {
	s accounts.Service

	logger *log.Entry
}

func (h *accountHandler) loginPage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "login.html", nil)
}

func (h *accountHandler) loginPerform(ctx echo.Context) error {
	m := &app.LoginInput{
		Username: ctx.Request().PostFormValue("username"),
		Password: ctx.Request().PostFormValue("password"),
	}
	checkbox := ctx.Request().PostFormValue("remember")

	if m.Validate() == false {
		return ctx.Render(http.StatusBadRequest, "login.html", m)
	}

	account, token, err := h.s.AuthLogin(ctx.Request().Context(), m)
	if err != nil {
		m.Errors["Content"] = err.Error()
		return ctx.Render(http.StatusBadRequest, "login.html", m)
	}

	var age int
	var remember bool
	if checkbox == "on" {
		remember = true
		age = 86400 * 7 // one weeks
	} else {
		age = 21600 // 6 hours
	}

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:     "/dashboard",
		MaxAge:   age,
		HttpOnly: true,
	}
	sess.Values["user_id"] = account.ID
	sess.Values["username"] = account.Username
	sess.Values["name"] = account.Name
	sess.Values["remember"] = remember
	err = sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		h.logger.Error(err)
	}

	h.writeCookie(ctx, "token", token)

	return ctx.Redirect(http.StatusFound, "/dashboard")
}

func (h *accountHandler) writeCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(6 * time.Hour)
	c.SetCookie(cookie)
}

func (h *accountHandler) restricted(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := ctx.Cookie("token")
		if err != nil {
			return next(ctx)
		}

		return ctx.Redirect(http.StatusFound, "/dashboard")
	}
}

type dashboardHandler struct {
	s accounts.Service

	logger *log.Entry
}

func (h *dashboardHandler) dashboardPage(ctx echo.Context) error {
	data := &app.DashboardContent{
		User: ctx.Get("context"),
	}

	return ctx.Render(http.StatusOK, "dashboard.html", data)
}

func (h *dashboardHandler) restricted(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := ctx.Cookie("session")
		if err != nil {
			cookie := new(http.Cookie)
			cookie.Name = "token"
			cookie.Value = ""
			cookie.Path = "/"
			cookie.Expires = time.Now().Add(0)
			ctx.SetCookie(cookie)
			return ctx.Redirect(http.StatusFound, "/accounts/login")
		}

		sess, err := session.Get("session", ctx)
		if err != nil {
			return ctx.Redirect(http.StatusFound, "/accounts/login")
		}

		ctx.Set("context", sess.Values)
		return next(ctx)
	}
}
