package api

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) LoginUser(ctx echo.Context) (err error) {
	var (
		req = m.GetUserReq{}
	)

	err = ctx.Bind(&req)
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, "bad request")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if req.Phone == "" || req.Password == "" {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, "bad request")
		return ctx.JSON(http.StatusBadRequest, res)
	}

	UserJWT, err := h.usecase.Login(ctx.Request().Context(), req.Phone, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no user found", []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		}
		log.Println("[Delivery][LoginUser] can't get user details, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponseLogin(http.StatusOK, utils.STATUS_SUCCESS, "logged in successfully", UserJWT)
	return ctx.JSON(http.StatusOK, res)
}
