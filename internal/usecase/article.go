package usecase

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"encoding/json"

	"log"

	"github.com/labstack/echo/v4"
)

func (u *usecase) GetArticles(ctx echo.Context, limit int, offset int) ([]interface{}, error) {
	var (
		articles []interface{}
	)

	resData, err := u.repository.GetArticles(ctx.Request().Context())
	if err != nil {
		log.Println("[Usecase][GetArticles] can't get list of articles, err:", err.Error())
		return articles, err
	}

	articles = make([]interface{}, len(resData))
	for i, v := range resData {
		utils.FormatTimeResArticle(&v)
		articles[i] = v
	}

	return articles, nil
}
func (u *usecase) GetArticleDetails(ctx echo.Context, id int) ([]interface{}, error) {
	var (
		article []interface{}
	)
	resData, err := u.repository.GetArticleDetails(ctx.Request().Context(), id)
	if err != nil {
		log.Println("[Usecase][GetArticleDetails] can't get article details, err:", err.Error())
		return article, err
	}
	utils.FormatTimeResArticle(&resData)

	article = append(article, resData)
	return article, nil
}
func (u *usecase) InsertArticle(ctx echo.Context, title string, slug string, html_content string, category_id int, metadata string) ([]interface{}, error) {
	var (
		article     []interface{}
		articleData m.Article
	)

	articleData = m.Article{
		Title:       title,
		Slug:        slug,
		HtmlContent: html_content,
		CategoryID:  category_id,
	}

	err := json.Unmarshal([]byte(metadata), &articleData.MetaData)
	if err != nil {
		log.Println("[Usecase][InsertArticle] can't insert article, err:", err.Error())
		return article, nil
	}

	resData, err := u.repository.InsertArticle(ctx.Request().Context(), articleData)
	if err != nil {
		log.Println("[Delivery][InsertArticle] can't insert category, err:", err.Error())
		return article, err
	}

	utils.FormatTimeResArticle(&resData)

	article = append(article, resData)
	return article, nil
}
func (u *usecase) UpdateArticle(ctx echo.Context, id int, title string, slug string, html_content string, category_id int, metadata string) ([]interface{}, error) {
	var (
		article        []interface{}
		resArticleData m.ResArticle
		articleData    m.Article
	)

	resArticleData, _ = u.repository.GetArticleDetails(ctx.Request().Context(), id)
	if title != "" {
		resArticleData.Title = title
	}

	if slug != "" {
		resArticleData.Slug = slug
	}

	if html_content != "" {
		resArticleData.HtmlContent = html_content
	}

	if category_id != 0 {
		resArticleData.ResCategory.Id = category_id
	}

	if metadata != "" {
		err := json.Unmarshal([]byte(metadata), &resArticleData.MetaData)
		if err != nil {
			log.Println("[Usecase][UpdateArticle] can't update article, err:", err.Error())
			return article, nil
		}
	}

	articleData = m.Article{
		Id:          id,
		Title:       resArticleData.Title,
		Slug:        resArticleData.Slug,
		HtmlContent: resArticleData.Slug,
		CategoryID:  resArticleData.ResCategory.Id,
		MetaData:    resArticleData.MetaData,
		CreatedAt:   resArticleData.CreatedAt,
	}

	articleData.UpdatedAt = utils.FormattedTime(utils.TimeNow)

	resArticleData, err := u.repository.UpdateArticle(ctx.Request().Context(), articleData)
	if err != nil {
		log.Println("[Usecase][UpdateArticle] can't update article, err:", err.Error())
		return article, err
	}

	resArticleData.CreatedAt = utils.FormattedTime(resArticleData.CreatedAt)
	resArticleData.UpdatedAt = utils.FormattedTime(utils.TimeNow)

	article = append(article, resArticleData)
	return article, nil
}
func (u *usecase) DeleteArticle(ctx echo.Context, id int) error {
	err := u.repository.DeleteArticle(ctx.Request().Context(), id)
	if err != nil {
		log.Println("[Usecase][DeleteArticle] can't delete article, err:", err.Error())
		return err
	}

	return nil
}
