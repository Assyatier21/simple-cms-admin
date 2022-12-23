package postgres

import (
	database "cms-admin/database/queries"
	m "cms-admin/models"
	"cms-admin/utils"
	"context"
	"database/sql"
	"encoding/json"
	"log"
)

func (r *repository) GetArticles(ctx context.Context, limit int, offset int) ([]m.ResArticle, error) {
	var (
		articles     []m.ResArticle
		metadata     m.MetaData
		rows         *sql.Rows
		tempMetaData []byte
		err          error
	)

	rows, err = r.db.Query(database.GetArticles, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrNotFound
		} else {
			log.Println("[GetArticles] can't get list of articles, err:", err.Error())
			return nil, err
		}
	}

	for rows.Next() {
		var temp = m.ResArticle{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.HtmlContent, &temp.ResCategory.Id, &temp.ResCategory.Title, &temp.ResCategory.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[GetArticles] failed to scan article, err :", err.Error())
			return nil, err
		}

		err = r.db.QueryRow(database.GetMetaData, temp.Id).Scan(&tempMetaData)
		if err != nil {
			log.Println("[GetArticleDetails] failed to scan metadata, err :", err.Error())
			return nil, err
		}

		json.Unmarshal(tempMetaData, &metadata)
		temp.MetaData = metadata

		temp.CreatedAt = utils.FormattedTime(temp.CreatedAt)
		temp.UpdatedAt = utils.FormattedTime(temp.UpdatedAt)
		articles = append(articles, temp)
	}

	if len(articles) > 0 {
		return articles, nil
	} else {
		return []m.ResArticle{}, nil
	}
}
func (r *repository) GetArticleDetails(ctx context.Context, id int) (m.ResArticle, error) {
	var (
		article      m.ResArticle
		err          error
		tempMetaData []byte
		metadata     m.MetaData
	)

	err = r.db.QueryRow(database.GetArticleDetails, id).Scan(&article.Id, &article.Title, &article.Slug, &article.HtmlContent, &article.ResCategory.Id, &article.ResCategory.Title, &article.ResCategory.Slug, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return m.ResArticle{}, utils.ErrNotFound
		} else {
			log.Println("[GetArticleDetails] failed to scan article, err:", err.Error())
			return m.ResArticle{}, err
		}
	}

	err = r.db.QueryRow(database.GetMetaData, article.Id).Scan(&tempMetaData)
	if err != nil {
		log.Println("[GetArticleDetails] failed to scan metadata, err :", err.Error())
		return m.ResArticle{}, err
	}

	json.Unmarshal(tempMetaData, &metadata)
	article.MetaData = metadata

	article.CreatedAt = utils.FormattedTime(article.CreatedAt)
	article.UpdatedAt = utils.FormattedTime(article.UpdatedAt)

	return article, nil
}
func (r *repository) InsertArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	return m.ResArticle{}, nil
}
func (r *repository) UpdateArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	return m.ResArticle{}, nil
}
func (r *repository) DeleteArticle(ctx context.Context, id int) error {
	return nil
}
