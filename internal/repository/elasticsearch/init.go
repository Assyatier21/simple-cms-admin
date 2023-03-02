package postgres

import (
	m "cms-admin/models"
	"context"

	"github.com/olivere/elastic/v7"
)

type ElasticHandler interface {
	GetArticles(ctx context.Context, limit int, offset int, sort_by string, order_by bool) ([]m.ResArticle, error)
	GetArticleDetails(ctx context.Context, query elastic.Query) (m.ResArticle, error)
	InsertArticle(ctx context.Context, article m.ResArticle) error
	UpdateArticle(ctx context.Context, article m.ResArticle) error
	DeleteArticle(ctx context.Context, id string) error

	GetCategoryTree(ctx context.Context, limit int, offset int) ([]m.Category, error)
	GetCategoryDetails(ctx context.Context, query elastic.Query) (m.Category, error)
	InsertCategory(ctx context.Context, category m.Category) error
	UpdateCategory(ctx context.Context, category m.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type elasticRepository struct {
	es *elastic.Client
}

func NewElasticRepository(es *elastic.Client) ElasticHandler {
	return &elasticRepository{
		es: es,
	}
}
