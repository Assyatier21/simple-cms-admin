package postgres

import (
	m "cms-admin/models"
	"context"
	"database/sql"
)

type RepositoryHandler interface {
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
