package server

import (
	"fmt"
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

	if m.Validate() == false {
		return ctx.Render(http.StatusBadRequest, "login.html", m)
	}

	account, token, err := h.s.AuthLogin(ctx.Request().Context(), m)
	if err != nil {
		m.Errors["Content"] = err.Error()
		return ctx.Render(http.StatusBadRequest, "login.html", m)
	}
	fmt.Println(token)

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:     "/dashboard",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["user_id"] = account.ID
	sess.Values["username"] = account.Username
	err = sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		h.logger.Error(err)
	}

	h.writeCookie(ctx, "token", token)

	return ctx.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *accountHandler) writeCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(6 * time.Hour)
	c.SetCookie(cookie)
}
