package postgres

import (
	"cms-admin/config"
	m "cms-admin/models"
	"context"
	"encoding/json"
	"log"
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
