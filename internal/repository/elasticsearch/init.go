package postgres

import (
	m "cms-admin/models"
	"context"

	"github.com/olivere/elastic/v7"
)

type ElasticHandler interface {
	InsertCategory(ctx context.Context, category m.Category) error
	GetCategoryDetails(ctx context.Context, id int) (m.Category, error)
}

type elasticRepository struct {
	es *elastic.Client
}

func NewElasticRepository(es *elastic.Client) ElasticHandler {
	return &elasticRepository{
		es: es,
	}
}
