package usecase

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"context"
	"log"
	"strconv"

	"github.com/olivere/elastic/v7"
)

func (u *usecase) GetCategoryTree(ctx context.Context) ([]interface{}, error) {
	var (
		categories []interface{}
		limit      int
		offset     int
	)

	limit = 100
	offset = 0

	resData, err := u.es.GetCategoryTree(ctx, limit, offset)
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
func (u *usecase) GetCategoryDetails(ctx context.Context, id int) ([]interface{}, error) {
	var (
		category []interface{}
		query    elastic.Query
	)
	// titleQuery := elastic.NewMatchQuery("title", "Sambo")
	// slugQuery := elastic.NewMatchQuery("slug", "new-sambo")

	// query = elastic.NewBoolQuery().Must(titleQuery, slugQuery)

	strId := strconv.Itoa(id)
	query = elastic.NewMatchQuery("id", strId)

	resData, err := u.es.GetCategoryDetails(ctx, query)
	if err != nil {
		log.Println("[Usecase][GetCategoryDetails] can't get category details, err:", err.Error())
		return category, err
	}

	if resData.Id == 0 {
		return category, nil
	}

	utils.FormatTimeResCategory(&resData)

	category = append(category, resData)
	return category, nil
}
func (u *usecase) InsertCategory(ctx context.Context, title string, slug string) ([]interface{}, error) {
	var (
		category     []interface{}
		categoryData m.Category
	)

	categoryData = m.Category{
		Title:     title,
		Slug:      slug,
		CreatedAt: utils.TimeNow,
		UpdatedAt: utils.TimeNow,
	}

	resData, err := u.repository.InsertCategory(ctx, categoryData)
	if err != nil {
		log.Println("[Usecase][InsertCategory] can't insert category, err:", err.Error())
		return category, err
	}

	err = u.es.InsertCategory(ctx, resData)
	if err != nil {
		log.Println("[Usecase][InsertCategory] can't insert category, err:", err.Error())
		return category, err
	}

	utils.FormatTimeResCategory(&resData)

	category = append(category, resData)
	return category, nil
}
func (u *usecase) UpdateCategory(ctx context.Context, id int, title string, slug string) ([]interface{}, error) {
	var (
		category []interface{}
		query    elastic.Query
	)

	strId := strconv.Itoa(id)
	query = elastic.NewMatchQuery("id", strId)

	categoryData, _ := u.es.GetCategoryDetails(ctx, query)

	if title != "" {
		categoryData.Title = title
	}

	if slug != "" {
		categoryData.Slug = slug
	}
	categoryData.CreatedAt = utils.FormattedTime(categoryData.CreatedAt)
	categoryData.UpdatedAt = utils.FormattedTime(utils.TimeNow)

	category = append(category, categoryData)
	return category, nil
}
func (u *usecase) DeleteCategory(ctx context.Context, id int) error {
	err := u.repository.DeleteCategory(ctx, id)
	if err != nil {
		log.Println("[Usecase][DeleteCategory] can't delete category, err:", err.Error())
		return err
	}

	return nil
}
