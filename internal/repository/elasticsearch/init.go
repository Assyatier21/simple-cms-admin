package postgres

import (
	"github.com/olivere/elastic/v7"
)

type ElasticHandler interface {
}

type repository struct {
	es *elastic.Client
}

func NewElasticRepository(es *elastic.Client) ElasticHandler {
	return &repository{
		es: es,
	}
}
