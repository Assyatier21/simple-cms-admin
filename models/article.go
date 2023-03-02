package models

type Article struct {
	Id          string   `json:"id" form:"id"`
	Title       string   `json:"title" form:"title"`
	Slug        string   `json:"slug" form:"slug"`
	HtmlContent string   `json:"html_content" form:"html_content"`
	CategoryID  int      `json:"category_id" form:"category_id"`
	MetaData    MetaData `json:"metadata" form:"metadata"`
	CreatedAt   string   `json:"created_at" form:"created_at"`
	UpdatedAt   string   `json:"updated_at" form:"updated_at"`
}

type ResArticle struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Slug        string      `json:"slug"`
	HtmlContent string      `json:"html_content"`
	ResCategory ResCategory `json:"category"`
	MetaData    MetaData    `json:"metadata"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

type MetaData struct {
	Title       string   `json:"meta_title"`
	Description string   `json:"meta_description"`
	Author      string   `json:"meta_author"`
	Keywords    []string `json:"meta_keywords"`
	Robots      []string `json:"meta_robots"`
}
