package database

const (
	GetCategoryTree = `SELECT * FROM cms_category 
							ORDER BY id`

	GetCategoryDetails = `SELECT * FROM cms_category 
							WHERE id = $1`

	InsertCategory = `INSERT INTO categories (title, slug, created_at, updated_at)
							VALUES ($1, $2, $3, $4)`

	UpdateCategory = `UPDATE categories SET title = $1, slug = $2, updated_at = $3
							WHERE id = $4`

	DeleteCategory = `DELETE FROM categories 
							WHERE id = $1`
)
