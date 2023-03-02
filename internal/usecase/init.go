package usecase

import (
	elastic "cms-admin/internal/repository/elasticsearch"
	"cms-admin/internal/repository/postgres"
	"context"
)

type UsecaseHandler interface {
	GetArticles(ctx context.Context, limit int, offset int, sort_by string, order_by string) ([]interface{}, error)
	GetArticleDetails(ctx context.Context, id string) ([]interface{}, error)
	InsertArticle(ctx context.Context, title string, slug string, htmlcontent string, categoryid int, metadata string) ([]interface{}, error)
	UpdateArticle(ctx context.Context, id string, title string, slug string, htmlcontent string, categoryid int, metadata string) ([]interface{}, error)
	DeleteArticle(ctx context.Context, id string) error

	GetCategoryTree(ctx context.Context) ([]interface{}, error)
	GetCategoryDetails(ctx context.Context, id int) ([]interface{}, error)
	InsertCategory(ctx context.Context, title string, slug string) ([]interface{}, error)
	UpdateCategory(ctx context.Context, id int, title string, slug string) ([]interface{}, error)
	DeleteCategory(ctx context.Context, id int) error
}

type usecase struct {
	repository postgres.RepositoryHandler
	es         elastic.ElasticHandler
}

func NewUsecase(repository postgres.RepositoryHandler, es elastic.ElasticHandler) UsecaseHandler {
	return &usecase{
		repository: repository,
		es:         es,
	}
}
