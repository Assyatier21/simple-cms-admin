package postgres

import (
	m "cms-admin/models"
	"context"
	"database/sql"
)

type Repository interface {
	GetArticles(ctx context.Context, limit int, offset int) ([]m.ResArticle, error)
	GetArticleDetails(ctx context.Context, id int) (m.ResArticle, error)
	GetCategoryTree(ctx context.Context) ([]m.Category, error)
	GetCategoryByID(ctx context.Context, id int) (m.Category, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
