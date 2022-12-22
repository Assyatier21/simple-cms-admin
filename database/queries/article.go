package database

const (
	GetArticles = `SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at 
						FROM cms_article a JOIN cms_category c ON a.category_id = c.id 
						ORDER BY a.id LIMIT $1 OFFSET $2`

	GetArticleDetails = `SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at 
							FROM cms_article a JOIN cms_category c ON a.category_id = c.id 
							WHERE a.id = $1`

	GetMetaData = `SELECT metadata 
						FROM cms_article WHERE id = $1`

	InsertArticles = `INSERT INTO cms_articles (id, title, slug, html_content, category_id, metadata, created_at, updated_at) 
							VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	UpdateArticles = `UPDATE cms_articles SET title = $1, slug = $2, html_content = $3, category_id = $4, metadata = $5, updated_at = $6 
						WHERE id = $7`

	DeleteArticles = `DELETE FROM cms_articles 
						WHERE id = $1`
)
