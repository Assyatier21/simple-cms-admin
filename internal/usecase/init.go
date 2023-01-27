package usecase

import (
	"cms-admin/internal/repository/postgres"

	"github.com/labstack/echo/v4"
)

type UsecaseHandler interface {
	GetArticles(ctx echo.Context, limit int, offset int) ([]interface{}, error)
	GetArticleDetails(ctx echo.Context, id int) ([]interface{}, error)
	InsertArticle(ctx echo.Context, title string, slug string, html_content string, category_id int, metadata string) ([]interface{}, error)
	UpdateArticle(ctx echo.Context, id int, title string, slug string, html_content string, category_id int, metadata string) ([]interface{}, error)
	DeleteArticle(ctx echo.Context, id int) error

	GetCategoryTree(ctx echo.Context) ([]interface{}, error)
	GetCategoryDetails(ctx echo.Context, id int) ([]interface{}, error)
	InsertCategory(ctx echo.Context, title string, slug string) ([]interface{}, error)
	UpdateCategory(ctx echo.Context, id int, title string, slug string) ([]interface{}, error)
	DeleteCategory(ctx echo.Context, id int) error
}

type usecase struct {
	repository postgres.RepositoryHandler
}

func NewUsecase(repository postgres.RepositoryHandler) UsecaseHandler {
	return &usecase{
		repository: repository,
	}
}
