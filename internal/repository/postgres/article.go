package postgres

import (
	database "cms-admin/database/queries"
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
	rows, err = r.db.Query(database.GetArticles, limit, offset)
	if err != nil {
		log.Println("[Repository][GetArticles] failed to get list of articles, err: ", err)
		return nil, err
	}

	for rows.Next() {
		var (
			temp         m.ResArticle
			byteMetadata []byte
		)

		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.HtmlContent, &temp.ResCategory.Id, &temp.ResCategory.Title, &temp.ResCategory.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[Repository][GetArticles] failed to scan article, err :", err)
			return nil, err
		}

		err = r.db.QueryRow(database.GetMetaData, temp.Id).Scan(&byteMetadata)
		if err != nil {
			log.Println("[Repository][GetArticles] failed to scan metadata, err :", err)
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
func (r *repository) GetArticleDetails(ctx context.Context, id string) (m.ResArticle, error) {
	var (
		article      m.ResArticle
		err          error
		byteMetadata []byte
	)

	err = r.db.QueryRow(database.GetArticleDetails, id).Scan(&article.Id, &article.Title, &article.Slug, &article.HtmlContent, &article.ResCategory.Id, &article.ResCategory.Title, &article.ResCategory.Slug, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		log.Println("[Repository][GetArticleDetails] failed to scan article, err: ", err)
		return m.ResArticle{}, err
	}

	err = r.db.QueryRow(database.GetMetaData, article.Id).Scan(&byteMetadata)
	if err != nil {
		log.Println("[Repository][GetArticleDetails] failed to scan metadata, err :", err)
		return m.ResArticle{}, err
	}
	json.Unmarshal(byteMetadata, &article.MetaData)

	return article, nil
}
func (r *repository) InsertArticle(ctx context.Context, article m.Article) (m.ResArticle, error) {
	var (
		resArticle m.ResArticle
		err        error
	)

	marshaled_metadata, err := json.Marshal(article.MetaData)
	if err != nil {
		log.Println("[Repository][InsertArticle] failed to insert article, err: ", err)
		return m.ResArticle{}, err
	}

	_, err = r.db.Exec(database.InsertArticle, article.Id, article.Title, article.Slug, article.HtmlContent, article.CategoryID, marshaled_metadata, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		log.Println("[Repository][InsertArticle] failed to insert article, err: ", err)
		return m.ResArticle{}, err
	}

	resArticle, err = r.GetArticleDetails(context.Background(), article.Id)
	if err != nil {
		log.Println("[Repository][InsertArticle] failed to get article details response, err: ", err)
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

	marshaled_metadata, err := json.Marshal(article.MetaData)
	if err != nil {
		log.Println("[Repository][UpdateArticle] failed to update article, err: ", err)
		return m.ResArticle{}, err
	}

	rows, err = r.db.Exec(database.UpdateArticle, &article.Title, &article.Slug, &article.HtmlContent, &article.CategoryID, marshaled_metadata, &article.UpdatedAt, &article.Id)
	if err != nil {
		log.Println("[Repository][UpdateArticle] failed to update article, err: ", err)
		return m.ResArticle{}, err
	}

	resArticle, err = r.GetArticleDetails(context.Background(), article.Id)
	if err != nil {
		log.Println("[Repository][UpdateArticle] failed to get article details response, err: ", err)
		return m.ResArticle{}, err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return m.ResArticle{}, nil
	}

	return resArticle, nil
}
func (r *repository) DeleteArticle(ctx context.Context, id string) error {
	rows, err := r.db.Exec(database.DeleteArticle, id)
	if err != nil {
		log.Println("[Repository][DeleteArticle] failed to delete article, err: ", err)
		return err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return msg.ERROR_NO_ROWS_AFFECTED
	}

	return nil
}
