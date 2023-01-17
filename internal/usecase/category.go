package usecase

import "github.com/labstack/echo/v4"

func (h *handler) GetCategoryTree(ctx echo.Context) ([]interface{}, error) {
	var categories []interface{}
	return categories, nil
}
func (h *handler) GetCategoryDetails(ctx echo.Context, id int) ([]interface{}, error) {
	var categories []interface{}
	return categories, nil
}
func (h *handler) InsertCategory(ctx echo.Context, title string, slug string) ([]interface{}, error) {
	var categories []interface{}
	return categories, nil
}
func (h *handler) UpdateCategory(ctx echo.Context, id int, title string, slug string) ([]interface{}, error) {
	var categories []interface{}
	return categories, nil
}
func (h *handler) DeleteCategory(ctx echo.Context, id int) error {
	return nil
}
