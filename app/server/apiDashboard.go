package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hinha/PAM-Trello/app"
	"github.com/hinha/PAM-Trello/app/accounts"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hinha/PAM-Trello/app/trello"
)

type apiDashboardHandler struct {
	trello  trello.Service
	account accounts.Service

	logger *log.Entry
}

func (h *apiDashboardHandler) performance(ctx echo.Context) error {

	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}
	claim := verify.(map[string]interface{})

	performance, err := h.trello.Performance(claim["id"].(string))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, performance)
}

func (h *apiDashboardHandler) analyticTrelloCard(ctx echo.Context) error {
	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}

	cards, err := h.trello.CardList()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": cards,
	})
}

func (h *apiDashboardHandler) kMethodsTrelloCard(ctx echo.Context) error {
	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}

	year := ctx.QueryParam("year")
	if strings.TrimSpace(year) == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "param cannot be empty"})
	}

	cardOut, clusters, average, scatterPlot, err := h.trello.GetClusters(ctx.Request().Context(), year)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	var cardOutResponse interface{}
	if cardOut == nil {
		cardOutResponse = []string{}
	} else {
		cardOutResponse = cardOut
	}

	var clusterResponse interface{}
	if clusters == nil {
		clusterResponse = []string{}
	} else {
		clusterResponse = clusters
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"card":              cardOutResponse,
		"cluster":           clusterResponse,
		"average":           average,
		"scatterClustering": scatterPlot,
	})
}

func (h *apiDashboardHandler) kMethodsData(ctx echo.Context) error {
	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}

	year := ctx.QueryParam("year")
	if strings.TrimSpace(year) == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "param cannot be empty"})
	}

	performance, err := h.trello.GetTotalTrello(year)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, performance)
}

func (h *apiDashboardHandler) userSetting(ctx echo.Context) error {
	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}
	claim := verify.(map[string]interface{})

	access, _ := h.account.GetAccessList(ctx.Request().Context())

	trelloList, err := h.trello.TrelloList(claim["id"].(string))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	listAccount, err := h.account.ListAccount(ctx.Request().Context(), claim["id"].(string), claim["role"].(string))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"trello":  trelloList,
		"access":  access,
		"account": listAccount,
	})
}

func (h *apiDashboardHandler) assignRoleAccount(ctx echo.Context) error {

	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}
	claim := verify.(map[string]interface{})

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

	err := h.account.NewAccessControlList(ctx.Request().Context(), claim["id"].(string), claim["role"].(string), m)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "ok",
	})
}

func (h *apiDashboardHandler) addUserSetting(ctx echo.Context) error {

	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}

	var bodyJSON app.RegisterInput

	// Read the Body content
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	// Restore the io.ReadCloser to its original state
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(ctx.Request().Body).Decode(&bodyJSON); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
	}

	if !bodyJSON.ValidateAPI() {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": bodyJSON.Errors})
	}

	claim := verify.(map[string]interface{})
	result, err := h.account.NewAccount(ctx.Request().Context(), claim["id"].(string), claim["role"].(string), &bodyJSON)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "ok",
		"data":    result,
	})
}

func (h *apiDashboardHandler) deleteUserSetting(ctx echo.Context) error {
	user := ctx.QueryParam("s")
	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}
	claim := verify.(map[string]interface{})

	separator := strings.Split(user, ",")
	if len(separator) != 2 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Errorf("can't allocate request").Error()})
	}

	err := h.account.DeleteAccount(ctx.Request().Context(), claim["id"].(string), claim["role"].(string), separator[0], separator[1])
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "ok",
	})
}

func (h *apiDashboardHandler) updateUserSetting(ctx echo.Context) error {
	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}
	claim := verify.(map[string]interface{})

	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	var body app.UpdateAccount
	// Restore the io.ReadCloser to its original state
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(ctx.Request().Body).Decode(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
		})
	}

	if !body.Validate() {
		return ctx.JSON(http.StatusBadRequest, body.Errors)
	}

	if err := h.account.UpdateAccount(ctx.Request().Context(), claim["id"].(string), claim["role"].(string), body); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "ok",
	})
}

func (h *apiDashboardHandler) trelloUserSetting(ctx echo.Context) error {
	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}

	var bodyJSON app.TrelloAddMember
	// Read the Body content
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	// Restore the io.ReadCloser to its original state
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	if err := json.NewDecoder(ctx.Request().Body).Decode(&bodyJSON); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
	}

	if !bodyJSON.Validate() {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": bodyJSON.Errors})
	}

	claim := verify.(map[string]interface{})
	result, err := h.trello.AddMember(claim["id"].(string), bodyJSON)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "ok",
		"data":    result,
	})
}

func (h *apiDashboardHandler) verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		credential := ctx.QueryParam("key")

		// will representation: email, id, role
		data, err := h.trello.Authorize(credential)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
		}

		ctx.Set("verify", data)

		return next(ctx)
	}
}
