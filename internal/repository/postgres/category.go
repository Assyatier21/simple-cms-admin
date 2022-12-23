package postgres

import (
	database "cms-admin/database/queries"
	m "cms-admin/models"
	"cms-admin/utils"
	"context"
	"database/sql"
	"log"
)

func (r *repository) GetCategoryTree(ctx context.Context) ([]m.Category, error) {
	var (
		categories []m.Category
		rows       *sql.Rows
		err        error
	)

	rows, err = r.db.Query(database.GetCategoryTree)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrNotFound
		} else {
			log.Println("[GetCategoryTree] can't get list of categories, err:", err.Error())
			return nil, err
		}
	}

	for rows.Next() {
		var temp = m.Category{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[GetCategoryTree] failed to scan category, err :", err.Error())
			return nil, err
		}
		temp.CreatedAt = utils.FormattedTime(temp.CreatedAt)
		temp.UpdatedAt = utils.FormattedTime(temp.UpdatedAt)
		categories = append(categories, temp)
	}

	if len(categories) > 0 {
		return categories, nil
	} else {
		return []m.Category{}, nil
	}
}
func (r *repository) GetCategoryDetails(ctx context.Context, id int) (m.Category, error) {
	var (
		category m.Category
		err      error
	)

	err = r.db.QueryRow(database.GetCategoryDetails, id).Scan(&category.Id, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return m.Category{}, utils.ErrNotFound
		} else {
			log.Println("[GetCategoryDetails] failed to scan category, err:", err.Error())
			return m.Category{}, err
		}
	}
	category.CreatedAt = utils.FormattedTime(category.CreatedAt)
	category.UpdatedAt = utils.FormattedTime(category.UpdatedAt)

	return category, nil
}
func (r *repository) InsertCategory(ctx context.Context, category m.Category) (m.Category, error) {
	return m.Category{}, nil
}
func (r *repository) UpdateCategory(ctx context.Context, category m.Category) (m.Category, error) {
	return m.Category{}, nil
}
func (r *repository) DeleteCategory(ctx context.Context, id int) error {
	return nil
}
