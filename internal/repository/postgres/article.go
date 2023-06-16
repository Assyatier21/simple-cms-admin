package postgres

import (
	DB_QUERY "cms-admin/database/queries"
	m "cms-admin/models"
	msg "cms-admin/models/lib"
	"context"
	"database/sql"
	"encoding/json"
	"log"
)

func (r *repository) GetArticles(ctx context.Context, limit int, offset int) ([]m.ResArticle, error) {
	var (
		articles []m.ResArticle
		rows     *sql.Rows
		err      error
	)
	rows, err = r.db.Query(DB_QUERY.GET_ARTICLES, limit, offset)
	if err != nil {
		log.Println("[Repository][GetArticles] can't get list of articles, err:", err.Error())
		return nil, err
	}

	for rows.Next() {
		var (
			temp         m.ResArticle
			byteMetadata []byte
		)

		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.HtmlContent, &temp.ResCategory.Id, &temp.ResCategory.Title, &temp.ResCategory.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[Repository][GetArticles] failed to scan article, err :", err.Error())
			return nil, err
		}

		err = r.db.QueryRow(DB_QUERY.GET_METADATA, temp.Id).Scan(&byteMetadata)
		if err != nil {
			log.Println("[Repository][GetArticles] failed to scan metadata, err :", err.Error())
			return nil, err
		}

		json.Unmarshal(byteMetadata, &temp.MetaData)
		articles = append(articles, temp)
	}

	if len(articles) == 0 {
		return []m.ResArticle{}, nil
	}

	return articles, nil
}
func (r *repository) GetArticleDetails(ctx context.Context, id int) (m.ResArticle, error) {
	var (
		article      m.ResArticle
		err          error
		byteMetadata []byte
	)

	err = r.db.QueryRow(DB_QUERY.GET_ARTICLE_DETAILS, id).Scan(&article.Id, &article.Title, &article.Slug, &article.HtmlContent, &article.ResCategory.Id, &article.ResCategory.Title, &article.ResCategory.Slug, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		log.Println("[Repository][GetArticleDetails] failed to scan article, err:", err.Error())
		return m.ResArticle{}, err
	}

	err = r.db.QueryRow(DB_QUERY.GET_METADATA, article.Id).Scan(&byteMetadata)
	if err != nil {
		log.Println("[Repository][GetArticleDetails] failed to scan metadata, err :", err.Error())
		return m.ResArticle{}, err
	}
	json.Unmarshal(byteMetadata, &article.MetaData)

	return article, nil
}
func (r *repository) InsertArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	var (
		lastId     int
		resArticle m.ResArticle
		err        error
	)

	marshalled_metadata, err := json.Marshal(article.MetaData)
	if err != nil {
		log.Println("[Repository][InsertArticle] can't insert article, err:", err.Error())
		return m.ResArticle{}, err
	}

	err = r.db.QueryRow(DB_QUERY.INSERT_ARTICLE, article.Title, article.Slug, article.HtmlContent, article.CategoryID, marshalled_metadata, article.CreatedAt, article.UpdatedAt).Scan(&lastId)
	if err != nil {
		log.Println("[Repository][InsertArticle] can't insert article, err:", err.Error())
		return m.ResArticle{}, err
	}

	resArticle, err = r.GetArticleDetails(context.Background(), lastId)
	if err != nil {
		log.Println("[Repository][InsertArticle] can't get article details response, err:", err.Error())
		return m.ResArticle{}, err
	}

	return resArticle, nil
}
func (r *repository) UpdateArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	var (
		resArticle m.ResArticle
		rows       sql.Result
		err        error
	)

	marshalled_metadata, err := json.Marshal(article.MetaData)
	if err != nil {
		log.Println("[Repository][UpdateArticle] can't update article, err:", err.Error())
		return m.ResArticle{}, err
	}

	rows, err = r.db.Exec(DB_QUERY.UPDATE_ARTICLE, &article.Title, &article.Slug, &article.HtmlContent, &article.CategoryID, marshalled_metadata, &article.UpdatedAt, &article.Id)
	if err != nil {
		log.Println("[Repository][UpdateArticle] can't update article, err:", err.Error())
		return m.ResArticle{}, err
	}

	resArticle, err = r.GetArticleDetails(context.Background(), article.Id)
	if err != nil {
		log.Println("[Repository][UpdateArticle] can't get article details response, err:", err.Error())
		return m.ResArticle{}, err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return m.ResArticle{}, nil
	}

	return resArticle, nil
}
func (r *repository) DeleteArticle(ctx context.Context, id int) error {
	rows, err := r.db.Exec(DB_QUERY.DELETE_ARTICLE, id)
	if err != nil {
		log.Println("[Repository][DeleteArticle] can't delete article, err:", err.Error())
		return err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return msg.ERROR_NO_ROWS_AFFECTED
	}

	return nil
}
