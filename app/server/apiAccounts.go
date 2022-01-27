package server

import (
	"github.com/gorilla/sessions"
	"github.com/hinha/PAM-Trello/app"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hinha/PAM-Trello/app/accounts"
)

type apiAccountHandler struct {
	s accounts.Service

	logger *log.Entry
}

func (h *apiAccountHandler) formToken(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("csrf").(string))
}

func (h *apiAccountHandler) loginPerform(ctx echo.Context) error {
	m := &app.LoginInput{
		Username: ctx.Request().PostFormValue("username"),
		Password: ctx.Request().PostFormValue("password"),
		Token:    ctx.Request().PostFormValue("csrf"),
	}
	checkbox := ctx.Request().PostFormValue("remember")
	if strings.TrimSpace(checkbox) == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "request failed"})
	}

	remeberMe, err := strconv.ParseBool(checkbox)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if m.Validate() == false {
		return ctx.JSON(http.StatusBadRequest, m.Errors)
	}

	var age int
	if remeberMe {
		age = 86400 * 7 // one weeks
		m.LongToken = true
	} else {
		age = 21600 // 6 hours
	}
	account, token, err := h.s.AuthLogin(ctx.Request().Context(), m)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: age,
	}
	sess.Values["user_id"] = account.ID
	sess.Values["username"] = account.Username
	sess.Values["name"] = account.Name
	sess.Values["remember"] = remeberMe
	err = sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		h.logger.Error(err)
	}

	h.writeCookie(ctx, "token", token)

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *apiAccountHandler) profileData(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	account, arn, secret, err := h.s.GetProfile(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"name":        account.Name,
		"username":    account.Username,
		"credentials": secret,
		"arn":         arn,
	})
}

func (h *apiAccountHandler) writeCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(6 * time.Hour)
	c.SetCookie(cookie)
}

func (h *apiAccountHandler) refreshToken(ctx echo.Context) error {

	oldToken := ctx.Get("jwt").(string)
	input := &app.LoginInput{}
	token, err := input.RefreshJwt(oldToken)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	h.writeCookie(ctx, "token", token)
	return ctx.JSON(http.StatusOK, map[string]interface{}{"new_token": token})
}
