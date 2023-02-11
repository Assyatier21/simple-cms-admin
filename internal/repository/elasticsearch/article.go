package postgres

import (
	"cms-admin/config"
	m "cms-admin/models"
	"context"
	"encoding/json"
	"log"

	"github.com/olivere/elastic/v7"
)

func (r *elasticRepository) GetArticles(ctx context.Context, limit int, offset int) ([]m.ResArticle, error) {
	var (
		articles = []m.ResArticle{}
		res      *elastic.SearchResult
		err      error
	)

	res, err = r.es.Search().Index(config.ES_INDEX_ARTICLE).From(offset).Size(limit).Do(ctx)
	if err != nil {
		return articles, err
	}

	if res.Hits.TotalHits.Value > 0 {
		for _, hit := range res.Hits.Hits {
			var article m.ResArticle
			err = json.Unmarshal(hit.Source, &article)
			if err != nil {
				panic(err)
			}
			articles = append(articles, article)
		}
	}

	return articles, err
}
func (r *elasticRepository) GetArticleDetails(ctx context.Context, query elastic.Query) (m.ResArticle, error) {
	var (
		article = m.ResArticle{}
		res     *elastic.SearchResult
		err     error
	)

	res, err = r.es.Search().Index(config.ES_INDEX_ARTICLE).Query(query).Do(ctx)
	if err != nil {
		return article, err
	}

	if res.Hits.TotalHits.Value > 0 {
		err = json.Unmarshal(res.Hits.Hits[0].Source, &article)
		if err != nil {
			panic(err)
		}
	}

	return article, nil
}
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

	return err
}
