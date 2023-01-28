package usecase

import (
	m "cms-admin/models"
	"cms-admin/utils"
	"context"
	"encoding/json"

	"log"
)

func (u *usecase) GetArticles(ctx context.Context, limit int, offset int) ([]interface{}, error) {
	var (
		articles []interface{}
	)

	resData, err := u.repository.GetArticles(ctx, limit, offset)
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
func (u *usecase) GetArticleDetails(ctx context.Context, id int) ([]interface{}, error) {
	var (
		article []interface{}
	)
	resData, err := u.repository.GetArticleDetails(ctx, id)
	if err != nil {
		log.Println("[Usecase][GetArticleDetails] can't get article details, err:", err.Error())
		return article, err
	}
	utils.FormatTimeResArticle(&resData)

	article = append(article, resData)
	return article, nil
}
func (u *usecase) InsertArticle(ctx context.Context, title string, slug string, htmlcontent string, categoryid int, metadata string) ([]interface{}, error) {
	var (
		article     []interface{}
		articleData m.Article
	)

	articleData = m.Article{
		Title:       title,
		Slug:        slug,
		HtmlContent: htmlcontent,
		CategoryID:  categoryid,
		CreatedAt:   utils.TimeNow,
		UpdatedAt:   utils.TimeNow,
	}

	err := json.Unmarshal([]byte(metadata), &articleData.MetaData)
	if err != nil {
		log.Println("[Usecase][InsertArticle] can't insert article, err:", err.Error())
		return article, nil
	}

	resData, err := u.repository.InsertArticle(ctx, articleData)
	if err != nil {
		log.Println("[Usecase][InsertArticle] can't insert category, err:", err.Error())
		return article, err
	}

	utils.FormatTimeResArticle(&resData)

	article = append(article, resData)
	return article, nil
}
func (u *usecase) UpdateArticle(ctx context.Context, id int, title string, slug string, htmlcontent string, categoryid int, metadata string) ([]interface{}, error) {
	var (
		article        []interface{}
		resArticleData m.ResArticle
		articleData    m.Article
	)

	resArticleData, _ = u.repository.GetArticleDetails(ctx, id)
	if title != "" {
		resArticleData.Title = title
	}

	if slug != "" {
		resArticleData.Slug = slug
	}

	if htmlcontent != "" {
		resArticleData.HtmlContent = htmlcontent
	}

	if categoryid != 0 {
		resArticleData.ResCategory.Id = categoryid
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
		HtmlContent: resArticleData.HtmlContent,
		CategoryID:  resArticleData.ResCategory.Id,
		MetaData:    resArticleData.MetaData,
		CreatedAt:   resArticleData.CreatedAt,
	}

	articleData.UpdatedAt = utils.FormattedTime(utils.TimeNow)

	resArticleData, err := u.repository.UpdateArticle(ctx, articleData)
	if err != nil {
		log.Println("[Usecase][UpdateArticle] can't update article, err:", err.Error())
		return article, err
	}

	resArticleData.CreatedAt = utils.FormattedTime(resArticleData.CreatedAt)
	resArticleData.UpdatedAt = utils.FormattedTime(utils.TimeNow)

	article = append(article, resArticleData)
	return article, nil
}
func (u *usecase) DeleteArticle(ctx context.Context, id int) error {
	err := u.repository.DeleteArticle(ctx, id)
	if err != nil {
		log.Println("[Usecase][DeleteArticle] can't delete article, err:", err.Error())
		return err
	}

	return nil
}
