package postgres

import (
	"cms-admin/config"
	m "cms-admin/models"
	"context"
	"encoding/json"
	"log"

	"github.com/olivere/elastic/v7"
)

func (r *elasticRepository) GetCategoryTree(ctx context.Context, limit int, offset int) ([]m.Category, error) {
	var (
		categories = []m.Category{}
		res        *elastic.SearchResult
		err        error
	)

	res, err = r.es.Search().Index(config.ES_INDEX_CATEGORY).From(offset).Size(limit).Do(ctx)
	if err != nil {
		return categories, err
	}

	if res.Hits.TotalHits.Value > 0 {
		for _, hit := range res.Hits.Hits {
			var category m.Category
			err = json.Unmarshal(hit.Source, &category)
			if err != nil {
				log.Println("[Elastic][GetCategoryDetails] failed to unmarshall category, err: ", err.Error())
				return categories, err
			}
			categories = append(categories, category)
		}
	}

	return categories, err
}
func (r *elasticRepository) GetCategoryDetails(ctx context.Context, query elastic.Query) (m.Category, error) {
	var (
		category = m.Category{}
		res      *elastic.SearchResult
		err      error
	)
	res, err = r.es.Search().Index(config.ES_INDEX_CATEGORY).Query(query).Do(ctx)
	if err != nil {
		return category, err
	}

	if res.Hits.TotalHits.Value > 0 {
		err = json.Unmarshal(res.Hits.Hits[0].Source, &category)
		if err != nil {
			log.Println("[Elastic][GetCategoryDetails] failed to unmarshall category, err: ", err.Error())
			return category, err
		}
	}

	return category, err
}
func (r *elasticRepository) InsertCategory(ctx context.Context, category m.Category) error {
	var (
		categoryJSON []byte
		body         string
		err          error
	)

	categoryJSON, err = json.Marshal(category)
	body = string(categoryJSON)
	_, err = r.es.Index().
		Index(config.ES_INDEX_CATEGORY).
		BodyJson(body).
		Do(ctx)

	if err != nil {
		log.Println("[Elastic][InsertCategory] can't insert category, err: ", err.Error())
		return err
	}
	return nil
}
