package usecase

import (
	"cms-admin/internal/repository/postgres"
	"context"
)

type UsecaseHandler interface {
	GetArticles(ctx context.Context, limit int, offset int) ([]interface{}, error)
	GetArticleDetails(ctx context.Context, id int) ([]interface{}, error)
	InsertArticle(ctx context.Context, title string, slug string, html_content string, category_id int, metadata string) ([]interface{}, error)
	UpdateArticle(ctx context.Context, id int, title string, slug string, html_content string, category_id int, metadata string) ([]interface{}, error)
	DeleteArticle(ctx context.Context, id int) error

	GetCategoryTree(ctx context.Context) ([]interface{}, error)
	GetCategoryDetails(ctx context.Context, id int) ([]interface{}, error)
	InsertCategory(ctx context.Context, title string, slug string) ([]interface{}, error)
	UpdateCategory(ctx context.Context, id int, title string, slug string) ([]interface{}, error)
	DeleteCategory(ctx context.Context, id int) error
}

type usecase struct {
	repository postgres.RepositoryHandler
}

func NewUsecase(repository postgres.RepositoryHandler) UsecaseHandler {
	return &usecase{
		repository: repository,
	}
}
