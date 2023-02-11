package postgres

import (
	"cms-admin/config"
	m "cms-admin/models"
	"context"
	"encoding/json"
	"log"
)

func (r *elasticRepository) InsertArticle(ctx context.Context, article m.ResArticle) error {
	var (
		articleJSON []byte
		body        string
		err         error
	)

	articleJSON, err = json.Marshal(article)
	body = string(articleJSON)
	_, err = r.es.Index().
		Index(config.ES_INDEX_ARTICLE).
		BodyJson(body).
		Do(ctx)

	if err != nil {
		log.Println("[Elastic][InsertArticle] can't insert article, err: ", err.Error())
		return err
	}
	return nil
}
