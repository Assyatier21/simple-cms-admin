package usecase

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"log"

	"github.com/labstack/echo/v4"
)

func (u *usecase) GetCategoryTree(ctx echo.Context) ([]interface{}, error) {
	var (
		categories []interface{}
	)

	resData, err := u.repository.GetCategoryTree(ctx.Request().Context())
	if err != nil {
		log.Println("[Usecase][GetCategoryTree] can't get list of categories, err:", err.Error())
		return categories, err
	}

	categories = make([]interface{}, len(resData))
	for i, v := range resData {
		utils.FormatTimeResCategory(&v)
		categories[i] = v
	}

	return categories, nil
}
func (u *usecase) GetCategoryDetails(ctx echo.Context, id int) ([]interface{}, error) {
	var category []interface{}

	resData, err := u.repository.GetCategoryDetails(ctx.Request().Context(), id)
	if err != nil {
		log.Println("[Usecase][GetCategoryDetails] can't get category details, err:", err.Error())
		return category, err
	}
	utils.FormatTimeResCategory(&resData)

	category = append(category, resData)
	return category, nil
}
func (u *usecase) InsertCategory(ctx echo.Context, title string, slug string) ([]interface{}, error) {
	var (
		category []interface{}
	)

	categoryData := m.Category{
		Title:     title,
		Slug:      slug,
		CreatedAt: utils.TimeNow,
		UpdatedAt: utils.TimeNow,
	}

	resData, err := u.repository.InsertCategory(ctx.Request().Context(), categoryData)
	if err != nil {
		log.Println("[Delivery][InsertCategory] can't insert category, err:", err.Error())
		return category, err
	}
	utils.FormatTimeResCategory(&resData)

	category = append(category, resData)
	return category, nil
}
func (u *usecase) UpdateCategory(ctx echo.Context, id int, title string, slug string) ([]interface{}, error) {
	var (
		category []interface{}
	)

	categoryData, _ := u.repository.GetCategoryDetails(ctx.Request().Context(), id)

	if title != "" {
		categoryData.Title = title
	}

	if slug != "" {
		categoryData.Slug = slug
	}
	categoryData.CreatedAt = utils.FormattedTime(categoryData.CreatedAt)
	categoryData.UpdatedAt = utils.FormattedTime(utils.TimeNow)

	categoryData, err := u.repository.UpdateCategory(ctx.Request().Context(), categoryData)
	if err != nil {
		log.Println("[Usecase][UpdateCategory] can't update category, err:", err.Error())
		return category, err
	}

	category = append(category, categoryData)
	return category, nil
}
func (u *usecase) DeleteCategory(ctx echo.Context, id int) error {
	err := u.repository.DeleteCategory(ctx.Request().Context(), id)
	if err != nil {
		log.Println("[Delivery][DeleteCategory] can't delete category, err:", err.Error())
		return err
	}

	return nil
}
