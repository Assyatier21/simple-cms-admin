package postgres

import (
	"cms-admin/config"
	m "cms-admin/models"
	"context"
	"encoding/json"
	"log"
	"strconv"

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
				log.Println("[Elastic][GetCategoryDetails] failed to unmarshal category, err: ", err)
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
			log.Println("[Elastic][GetCategoryDetails] failed to unmarshal category, err: ", err)
			return category, err
		}
	}

	return category, err
}
func (r *elasticRepository) InsertCategory(ctx context.Context, category m.Category) error {
	var (
		categoryJSON []byte
		category_id  string
		body         string
		err          error
	)

	category_id = strconv.Itoa(category.Id)
	categoryJSON, err = json.Marshal(category)
	if err != nil {
		log.Println("[Elastic][InsertCategory] failed to marshal category, err: ", err)
		return err
	}

	body = string(categoryJSON)
	_, err = r.es.Index().
		Index(config.ES_INDEX_CATEGORY).
		Id(category_id).
		BodyJson(body).
		Do(ctx)

	if err != nil {
		log.Println("[Elastic][InsertCategory] failed to insert category, err: ", err)
		return err
	}

	return nil
}
func (r *elasticRepository) UpdateCategory(ctx context.Context, category m.Category) error {
	var (
		category_id string
		err         error
	)

	category_id = strconv.Itoa(category.Id)
	_, err = r.es.Update().
		Index(config.ES_INDEX_CATEGORY).
		Id(category_id).
		Doc(category).
		Do(ctx)

	if err != nil {
		log.Println("[Elastic][UpdateCategory] failed to update category, err: ", err)
		return err
	}

	return nil
}
func (r *elasticRepository) DeleteCategory(ctx context.Context, id int) error {
	var (
		category_id string
		err         error
	)

	category_id = strconv.Itoa(id)
	_, err = r.es.Delete().
		Index(config.ES_INDEX_CATEGORY).
		Id(category_id).
		Do(ctx)
	if err != nil {
		log.Println("[Elastic][DeleteCategory] failed to delete category, err: ", err)
		return err
	}

	return nil
}
