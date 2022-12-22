package api

import (
	"cms-admin/internal/repository/postgres"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetArticles(ctx echo.Context) (err error)
	GetArticleDetails(ctx echo.Context) (err error)
	GetCategoryTree(ctx echo.Context) (err error)
	GetCategoryByID(ctx echo.Context) (err error)
}

type handler struct {
	repository postgres.Repository
}

func New(repository postgres.Repository) Handler {
	return &handler{
		repository: repository,
	}
}
