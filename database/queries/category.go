package database

const (
	GET_CATEGORY_TREE = `SELECT * FROM cms_category 
                     ORDER BY id`

	GET_CATEGORY_DETAILS = `SELECT * FROM cms_category 
                        WHERE id = $1`

	INSERT_CATEGORY = `INSERT INTO cms_category (title, slug, created_at, updated_at)
                   VALUES ($1, $2, $3, $4) RETURNING id`

	UPDATE_CATEGORY = `UPDATE cms_category SET title = $1, slug = $2, updated_at = $3
                   WHERE id = $4`

	DELETE_CATEGORY = `DELETE FROM cms_category 
                   WHERE id = $1`
)
