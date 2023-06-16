package database

const (
	GET_ARTICLES = `SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at 
                FROM cms_article a JOIN cms_category c ON a.category_id = c.id 
                ORDER BY a.id LIMIT $1 OFFSET $2`

	GET_ARTICLE_DETAILS = `SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at 
                       FROM cms_article a JOIN cms_category c ON a.category_id = c.id 
                       WHERE a.id = $1`

	GET_METADATA = `SELECT metadata 
                FROM cms_article WHERE id = $1`

	INSERT_ARTICLE = `INSERT INTO cms_article (title, slug, html_content, category_id, metadata, created_at, updated_at) 
                  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	UPDATE_ARTICLE = `UPDATE cms_article SET title = $1, slug = $2, html_content = $3, category_id = $4, metadata = $5, updated_at = $6 
                  WHERE id = $7`

	DELETE_ARTICLE = `DELETE FROM cms_article 
                  WHERE id = $1`
)
