package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/accounts"
)

type accountHandler struct {
	s accounts.Service

	logger *log.Entry
}

func (h *accountHandler) loginPage(ctx echo.Context) error {
	data := map[string]interface{}{
		"Token": ctx.Get("csrf"),
	}
	return ctx.Render(http.StatusOK, "login.html", data)
}

func (h *accountHandler) loginPerform(ctx echo.Context) error {
	m := &app.LoginInput{
		Username: ctx.Request().PostFormValue("username"),
		Password: ctx.Request().PostFormValue("password"),
		Token:    ctx.Request().PostFormValue("csrf"),
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

func (h *accountHandler) registerAccount(ctx echo.Context) error {

	// Read the Body content
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	m := new(app.RegisterInput)
	// Restore the io.ReadCloser to its original state
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(ctx.Request().Body).Decode(m); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
		})

	}

	if !m.ValidateAPI() {
		return ctx.JSON(http.StatusBadRequest, m.Errors)
	}

	roleName := ctx.Get("authorize").(string)
	adminID := ctx.QueryParam("key")
	err := h.s.NewAccount(ctx.Request().Context(), adminID, roleName, m)
	if err != nil {
		m.Errors["message"] = err.Error()
		return ctx.JSON(http.StatusBadRequest, m.Errors)
	}

	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *accountHandler) accountTable(ctx echo.Context) error {
	roleName := ctx.Get("authorize").(string)
	adminID := ctx.QueryParam("key")

	listAccount, err := h.s.ListAccount(ctx.Request().Context(), adminID, roleName)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": listAccount,
	})
}

func (h *accountHandler) deleteAccount(ctx echo.Context) error {
	roleName := ctx.Get("authorize").(string)
	adminID := ctx.QueryParam("key")
	user := ctx.QueryParam("s")

	separator := strings.Split(user, ",")
	if len(separator) != 2 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Errorf("can't allocate request").Error()})
	}

	err := h.s.DeleteAccount(ctx.Request().Context(), adminID, roleName, separator[0], separator[1])
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

func (h *accountHandler) assignRoleAccount(ctx echo.Context) error {
	adminID := ctx.Request().Header.Get("Admin-Id")
	if adminID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}

	// Read the Body content
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	m := new(app.AssignRole)
	// Restore the io.ReadCloser to its original state
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(ctx.Request().Body).Decode(m); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
		})

	}

	if !m.Validate() {
		return ctx.JSON(http.StatusBadRequest, m.Errors)
	}

	roleName := ctx.Get("authorize").(string)
	err := h.s.NewAccessControlList(ctx.Request().Context(), adminID, roleName, m)
	if err != nil {
		m.Errors["message"] = err.Error()
		return ctx.JSON(http.StatusBadRequest, m.Errors)
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

func (h *accountHandler) refreshToken(ctx echo.Context) error {

	c, err := ctx.Cookie("token")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "token error"})
	}

	input := &app.LoginInput{}
	token, err := input.RefreshJwt(c.Value)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
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

	hub    *Hub
	logger *log.Entry
}

var upgraded = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *dashboardHandler) dashboardPage(ctx echo.Context) error {
	data := &app.DashboardContent{
		User: ctx.Get("context"),
		Any:  make(map[string]string),
		Page: make(map[string]int),
	}

	data.Any["Location"] = "/token/refresh"
	data.Page["Menu"] = int(app.HomeMenu)

	return ctx.Render(http.StatusOK, "dashboard.html", data)
}

func (h *dashboardHandler) boardTrelloPage(ctx echo.Context) error {
	data := &app.DashboardContent{
		User:  ctx.Get("context"),
		Token: ctx.Get("csrf").(string),
		Any:   make(map[string]string),
		Page:  make(map[string]int),
	}

	data.Any["Location"] = "/token/refresh"
	data.Page["Menu"] = int(app.TrelloMenu)

	return ctx.Render(http.StatusOK, "dashboard.html", data)
}

func (h *dashboardHandler) engine(ctx echo.Context) error {
	get := ctx.Get("context")
	username := get.(map[interface{}]interface{})["username"].(string)
	name := get.(map[interface{}]interface{})["name"].(string)
	userID := get.(map[interface{}]interface{})["user_id"].(string)

	// Upgrading the HTTP connection socket connection
	connection, err := upgraded.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		h.logger.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	CreateNewSocketUser(h.hub, connection, userID, username, name)

	return nil
}

func (h *dashboardHandler) restricted(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_, err := ctx.Cookie("session")
		if err != nil {
			ctx.SetCookie(app.DeleteCookie)
			return ctx.Redirect(http.StatusFound, "/accounts/login")
		}

		sess, err := session.Get("session", ctx)
		if err != nil {
			ctx.SetCookie(app.DeleteCookie)
			return ctx.Redirect(http.StatusFound, "/accounts/login")
		}

		ctx.Set("context", sess.Values)
		return next(ctx)
	}
}

func (h *dashboardHandler) settingDetails(ctx echo.Context) error {
	data := &app.DashboardContent{
		User: ctx.Get("context"),
		Any:  make(map[string]string),
		Page: make(map[string]int),
	}

	data.Any["Location"] = "/token/refresh"
	data.Page["Menu"] = int(app.SettingDetails)

	return ctx.Render(http.StatusOK, "dashboard.html", data)
}

func (h *dashboardHandler) settingUsers(ctx echo.Context) error {
	data := &app.DashboardContent{
		User:  ctx.Get("context"),
		Token: ctx.Get("csrf").(string),
		Any:   make(map[string]string),
		Page:  make(map[string]int),
		Data:  make(map[string]interface{}),
	}

	mmap := ctx.Get("context").(map[interface{}]interface{})
	data.Any["Location"] = "/token/refresh"
	data.Any["Key"] = mmap["user_id"].(string)
	data.Page["Menu"] = int(app.SettingUsers)

	control, err := h.s.GetAccessList(ctx.Request().Context())
	if err != nil {
		h.logger.Error(err)
	}
	data.Data["AccessControl"] = control

	return ctx.Render(http.StatusOK, "dashboard.html", data)
}
