package api

import (
	"cms-admin/internal/usecase"

	"github.com/labstack/echo/v4"
)

type DeliveryHandler interface {
	GetArticles(ctx echo.Context) (err error)
	GetArticleDetails(ctx echo.Context) (err error)
	InsertArticle(ctx echo.Context) (err error)
	UpdateArticle(ctx echo.Context) (err error)
	DeleteArticle(ctx echo.Context) (err error)

	GetCategoryTree(ctx echo.Context) (err error)
	GetCategoryDetails(ctx echo.Context) (err error)
	InsertCategory(ctx echo.Context) (err error)
	UpdateCategory(ctx echo.Context) (err error)
	DeleteCategory(ctx echo.Context) (err error)
}

type handler struct {
	usecase usecase.UsecaseHandler
}

func NewHandler(usecase usecase.UsecaseHandler) DeliveryHandler {
	return &handler{
		usecase: usecase,
	}
}
