package helper

import (
	m "cms-admin/models"

	"github.com/labstack/echo/v4"
)

func WriteResponse(ctx echo.Context, statusCode int, status string, message string, data []interface{}) error {
	if statusCode > 299 {
		res := m.SetError(statusCode, status, message)
		return ctx.JSON(statusCode, res)
	} else {
		res := m.SetResponse(statusCode, status, message, data)
		return ctx.JSON(statusCode, res)
	}
}
