package usecase

import (
	m "cms-admin/models"

	"github.com/labstack/echo/v4"
)

func (u *usecase) GetArticles(ctx echo.Context) ([]interface{}, error) {
	return []interface{}{}, nil
}
func (u *usecase) GetArticleDetails(ctx echo.Context, id int) ([]interface{}, error) {
	return []interface{}{}, nil
}
func (u *usecase) InsertArticle(ctx echo.Context, title string, slug string, html_content string, category_id int, metadata m.MetaData) ([]interface{}, error) {
	return []interface{}{}, nil
}
func (u *usecase) UpdateArticle(ctx echo.Context, id int, title string, slug string, html_content string, category_id int, metadata m.MetaData) ([]interface{}, error) {
	return []interface{}{}, nil
}
func (u *usecase) DeleteArticle(ctx echo.Context, id int) error {
	return nil
}
