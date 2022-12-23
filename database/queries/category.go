package database

const (
	GetCategoryTree = `SELECT * FROM cms_category 
							ORDER BY id`

	GetCategoryDetails = `SELECT * FROM cms_category 
							WHERE id = $1`

	InsertCategory = `INSERT INTO categories (id, title, slug, created_at, updated_at)
							VALUES ($1, $2, $3, $4, $5)`

	UpdateCategory = `UPDATE categories SET title = $1, slug = $2, created_at = $3, updated_at = $4
							WHERE id = $5`

	DeleteCategory = `DELETE FROM categories 
							WHERE id = $1`
)
