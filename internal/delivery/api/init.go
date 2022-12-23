package api

import (
	"cms-admin/internal/repository/postgres"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	// Article Handler
	GetArticles(ctx echo.Context) (err error)
	GetArticleDetails(ctx echo.Context) (err error)
	InsertArticle(ctx echo.Context) (err error)
	UpdateArticle(ctx echo.Context) (err error)
	DeleteArticle(ctx echo.Context) (err error)

	// Category Handler
	GetCategoryTree(ctx echo.Context) (err error)
	GetCategoryDetails(ctx echo.Context) (err error)
	InsertCategory(ctx echo.Context) (err error)
	UpdateCategory(ctx echo.Context) (err error)
	DeleteCategory(ctx echo.Context) (err error)
}

type handler struct {
	repository postgres.Repository
}

func New(repository postgres.Repository) Handler {
	return &handler{
		repository: repository,
	}
}
