package models

type Category struct {
	Id        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Slug      string `json:"slug" form:"slug"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}

type ResCategory struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}
