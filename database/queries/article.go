package database

const (
	GetArticles = `SELECT a.id, a.title, a.slug, a.htmlcontent, c.id, c.title , c.slug, a.created_at, a.updated_at 
						FROM cms_article a JOIN cms_category c ON a.categoryid = c.id 
						ORDER BY a.id LIMIT $1 OFFSET $2`

	GetArticleDetails = `SELECT a.id, a.title, a.slug, a.htmlcontent, c.id, c.title , c.slug, a.created_at, a.updated_at 
							FROM cms_article a JOIN cms_category c ON a.categoryid = c.id 
							WHERE a.id = $1`

	GetMetaData = `SELECT metadata 
						FROM cms_article WHERE id = $1`

	InsertArticle = `INSERT INTO cms_article (title, slug, htmlcontent, categoryid, metadata, created_at, updated_at) 
							VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	UpdateArticle = `UPDATE cms_article SET title = $1, slug = $2, htmlcontent = $3, categoryid = $4, metadata = $5, updated_at = $6 
						WHERE id = $7`

	DeleteArticle = `DELETE FROM cms_article 
						WHERE id = $1`
)
