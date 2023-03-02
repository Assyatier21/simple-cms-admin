package helper

import (
	m "cms-admin/models"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GenerateUUIDString() string {
	id := uuid.New()
	stringID := id.String()
	uuid := strings.Replace(stringID, "-", "", -1)
	return uuid
}

func ValidateSortBy(sort_by string) string {
	if sort_by == "title" {
		sort_by = "title.keyword"
	} else if sort_by == "slug" {
		sort_by = "slug.keyword"
	} else if sort_by == "html_content" {
		sort_by = "html_content.keyword"
	} else {
		sort_by = "updated_at"
	}

	return sort_by
}

func ValidateOrderBy(order_by string) bool {
	var order_by_bool bool

	order_by_bool = false
	if order_by == "asc" {
		order_by_bool = true
	} else if order_by == "desc" {
		order_by_bool = false
	}

	return order_by_bool
}

func WriteResponse(ctx echo.Context, statusCode int, status string, message string, data []interface{}) error {
	if statusCode > 299 {
		res := m.SetError(statusCode, status, message)
		return ctx.JSON(statusCode, res)
	} else {
		res := m.SetResponse(statusCode, status, message, data)
		return ctx.JSON(statusCode, res)
	}
}
