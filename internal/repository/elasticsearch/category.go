package postgres

import (
	"cms-admin/config"
	m "cms-admin/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/olivere/elastic/v7"
)

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
		log.Println("[Elastic][InsertCategory] can't insert category, err:", err.Error())
		return err
	}
	return nil
}
func (r *elasticRepository) GetCategoryDetails(ctx context.Context, id int) (m.Category, error) {
	var (
		category = m.Category{}
		res      *elastic.SearchResult
		err      error
	)
	strId := strconv.Itoa(id)

	query := elastic.NewMatchQuery("id", strId)
	res, err = r.es.Search().Index(config.ES_INDEX_CATEGORY).Query(query).Do(ctx)
	if err != nil {
		return category, err
	}

	if res.Hits.TotalHits.Value > 0 {
		err = json.Unmarshal(res.Hits.Hits[0].Source, &category)
		fmt.Println(category)
		if err != nil {
			panic(err)
		}
	}

	return category, err
}
