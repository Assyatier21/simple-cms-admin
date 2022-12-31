package postgres

import (
	database "cms-admin/database/queries"
	m "cms-admin/models"
	"cms-admin/utils"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"reflect"
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
		utils.FormatTimeResArticle(&temp)
		json.Unmarshal(tempMetaData, &metadata)
		temp.MetaData = metadata
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
	utils.FormatTimeResArticle(&article)
	json.Unmarshal(tempMetaData, &metadata)
	article.MetaData = metadata

	return article, nil
}
func (r *repository) InsertArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	var (
		lastId     int
		resArticle m.ResArticle
	)

	marshalMetadata, _ := json.Marshal(article.MetaData)
	err := r.db.QueryRow(database.InsertArticle, article.Title, article.Slug, article.HtmlContent, article.CategoryID, marshalMetadata, article.CreatedAt, article.UpdatedAt).Scan(&lastId)
	if err != nil {
		log.Println("[InsertArticle] can't insert article, err:", err.Error())
		return m.ResArticle{}, err
	}

	resArticle, err = r.GetArticleDetails(context.Background(), lastId)
	if err != nil {
		log.Println("[InsertArticle][GetArticleDetails] can't get article details response, err:", err.Error())
		return m.ResArticle{}, err
	}

	return resArticle, nil
}
func (r *repository) UpdateArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	var (
		resArticle      m.ResArticle
		err             error
		marshalMetadata []byte
	)

	resArticle, err = r.GetArticleDetails(context.Background(), int(article.Id))
	if err != nil {
		log.Println("[UpdateArticle][GetArticleDetails] can't get article details, err:", err.Error())
		return m.ResArticle{}, err
	}

	if article.Title == "" {
		article.Title = resArticle.Title
	}
	if article.Slug == "" {
		article.Slug = resArticle.Slug
	}
	if article.HtmlContent == "" {
		article.HtmlContent = resArticle.HtmlContent
	}
	if article.CategoryID == 0 {
		article.CategoryID = resArticle.ResCategory.Id
	}
	if reflect.DeepEqual(article.MetaData, m.MetaData{}) {
		article.MetaData = resArticle.MetaData
	} else {
		marshalMetadata, _ = json.Marshal(article.MetaData)
	}

	rows, err := r.db.Exec(database.UpdateArticle, &article.Title, &article.Slug, &article.HtmlContent, &article.CategoryID, marshalMetadata, &article.UpdatedAt, &article.Id)
	if err != nil {
		log.Println("[UpdateArticle] can't update article, err:", err.Error())
		return m.ResArticle{}, err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		return resArticle, nil
	} else {
		log.Println("[UpdateArticle] err:", utils.NoRowsAffected)
		return m.ResArticle{}, utils.NoRowsAffected
	}
}
func (r *repository) DeleteArticle(ctx context.Context, id int) error {
	rows, err := r.db.Exec(database.DeleteArticle, id)
	if err != nil {
		log.Println("[DeleteArticle] can't delete article, err:", err.Error())
		return err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		return nil
	} else {
		return utils.NoRowsAffected
	}
}
