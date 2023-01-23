package postgres

import (
	m "cms-admin/models"
	"context"
	"database/sql"
)

type RepositoryHandler interface {
	GetArticles(ctx context.Context, limit int, offset int) ([]m.ResArticle, error)
	GetArticleDetails(ctx context.Context, id int) (m.ResArticle, error)
	InsertArticle(ctx context.Context, article m.Article) (m.ResArticle, error)
	UpdateArticle(ctx context.Context, article m.Article) (m.ResArticle, error)
	DeleteArticle(ctx context.Context, id int) error

	GetCategoryTree(ctx context.Context) ([]m.Category, error)
	GetCategoryDetails(ctx context.Context, id int) (m.Category, error)
	InsertCategory(ctx context.Context, category m.Category) (m.Category, error)
	UpdateCategory(ctx context.Context, category m.Category) (m.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) RepositoryHandler {
	return &repository{
		db: db,
	}
}
