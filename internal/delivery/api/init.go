package api

import (
	"cms-admin/internal/usecase"

	"github.com/labstack/echo/v4"
)

type DeliveryHandler interface {
	GetCategoryTree(ctx echo.Context) (err error)
	GetCategoryDetails(ctx echo.Context) (err error)
	InsertCategory(ctx echo.Context) (err error)
	UpdateCategory(ctx echo.Context) (err error)
	DeleteCategory(ctx echo.Context) (err error)
}

type handler struct {
	usecase usecase.UsecaseHandler
}

func New(usecase usecase.UsecaseHandler) DeliveryHandler {
	return &handler{
		usecase: usecase,
	}
}
