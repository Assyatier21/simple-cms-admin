package usecase

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"context"
	"errors"
	"log"
	"strconv"
	"sync"

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
		log.Println("[Usecase][GetCategoryTree] failed to get list of categories, err: ", err)
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

	strId := strconv.Itoa(id)
	query = elastic.NewMatchQuery("id", strId)

	resData, err := u.es.GetCategoryDetails(ctx, query)
	if err != nil {
		log.Println("[Usecase][GetCategoryDetails] failed to get category details, err: ", err)
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
		log.Println("[Usecase][InsertCategory] failed to insert category, err: ", err)
		return category, err
	}

	err = u.es.InsertCategory(ctx, resData)
	if err != nil {
		log.Println("[Usecase][InsertCategory] failed to insert category to elastic, err: ", err)
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

	resData, err := u.repository.UpdateCategory(ctx, categoryData)
	if err != nil {
		log.Println("[Usecase][UpdateCategory] failed to update category, err: ", err)
		return category, err
	}

	err = u.es.UpdateCategory(ctx, categoryData)
	if err != nil {
		log.Println("[Usecase][InsertCategory] failed to update category to elastic, err: ", err)
		return category, err
	}

	resData.CreatedAt = utils.FormattedTime(resData.CreatedAt)
	resData.UpdatedAt = utils.FormattedTime(utils.TimeNow)

	category = append(category, resData)
	return category, nil
}
func (u *usecase) DeleteCategory(ctx context.Context, id int) error {
	var (
		categoryDeleted bool
		elasticDeleted  bool
		err             error
		wg              sync.WaitGroup
	)

	wg = sync.WaitGroup{}
	wg.Add(2)

	go func() {
		err = u.repository.DeleteCategory(ctx, id)
		if err != nil {
			log.Println("[Usecase][DeleteCategory] failed to delete category, err: ", err)
		} else {
			categoryDeleted = true
		}
		wg.Done()
	}()

	go func() {
		err = u.es.DeleteCategory(ctx, id)
		if err != nil {
			log.Println("[Usecase][DeleteCategory] failed to delete category from elastic, err: ", err)
		} else {
			elasticDeleted = true
		}
		wg.Done()
	}()

	wg.Wait()

	if !categoryDeleted && !elasticDeleted {
		err = errors.New("category not found")
		return err
	}

	return nil
}
