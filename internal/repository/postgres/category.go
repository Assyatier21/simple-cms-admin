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
		log.Println("[Repository][GetCategoryTree] can't get list of categories, err:", err.Error())
		return nil, err
	}

	for rows.Next() {
		var temp = m.Category{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[Repository][GetCategoryTree] failed to scan category, err :", err.Error())
			return nil, err
		}
		categories = append(categories, temp)
	}

	if len(categories) == 0 {
		return []m.Category{}, nil
	}

	return categories, nil
}
func (r *repository) GetCategoryDetails(ctx context.Context, id int) (m.Category, error) {
	var (
		category m.Category
		err      error
	)

	err = r.db.QueryRow(database.GetCategoryDetails, id).Scan(&category.Id, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.Println("[Repository][GetCategoryDetails] failed to scan category, err:", err.Error())
		return m.Category{}, err
	}

	return category, nil
}
func (r *repository) InsertCategory(ctx context.Context, category m.Category) (m.Category, error) {
	err := r.db.QueryRow(database.InsertCategory, category.Title, category.Slug, category.CreatedAt, category.UpdatedAt).Scan(&category.Id)
	if err != nil {
		log.Println("[Repository][InsertCategory] can't insert category, err:", err.Error())
		return m.Category{}, err
	}

	return category, nil
}
func (r *repository) UpdateCategory(ctx context.Context, category m.Category) (m.Category, error) {
	rows, err := r.db.Exec(database.UpdateCategory, &category.Title, &category.Slug, &category.UpdatedAt, &category.Id)
	if err != nil {
		log.Println("[Repository][UpdateCategory] can't update category, err:", err.Error())
		return m.Category{}, err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return m.Category{}, nil
	}

	return category, nil
}
func (r *repository) DeleteCategory(ctx context.Context, id int) error {
	rows, err := r.db.Exec(database.DeleteCategory, id)
	if err != nil {
		log.Println("[Repository][DeleteCategory] can't delete category, err:", err.Error())
		return err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected == 0 {
		return utils.NoRowsAffected
	}

	return nil
}
