package utils

import (
	m "cms-admin/models"
	"errors"
	"log"
	"regexp"
	"time"
)

var (
	PathArticles   = "/articles"
	PathArticle    = "/article"
	PathCategories = "/categories"
	PathCategory   = "/category"
	ErrNotFound    = errors.New("data not found")
	NoRowsAffected = errors.New("no rows affected")
	jakartaLoc, _  = time.LoadLocation("Asia/Jakarta")
	TimeNow        = time.Now().In(jakartaLoc).Format("2006-01-02T15:04:05Z")
)

func IsValidAlphabet(s string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z]*$`)
	return regex.MatchString(s)
}

func IsValidNumeric(s string) bool {
	regex, _ := regexp.Compile(`([0-9])`)
	return regex.MatchString(s)
}

func IsValidAlphaNumeric(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z0-9]*$)`)
	return regex.MatchString(s)
}

func IsValidSlug(s string) bool {
	regex, _ := regexp.Compile(`^[a-z0-9-]+$`)
	return regex.MatchString(s)
}

func SetArticleCreatedUpdatedTimeNow(article *m.Article) m.Article {
	article.CreatedAt = TimeNow
	article.UpdatedAt = TimeNow
	return *article
}
func SetArticleUpdatedTimeNow(article *m.Article) m.Article {
	article.UpdatedAt = TimeNow
	return *article
}
func FormatTimeResArticle(article *m.ResArticle) m.ResArticle {
	article.CreatedAt = FormattedTime(article.CreatedAt)
	article.UpdatedAt = FormattedTime(article.UpdatedAt)
	return *article
}

func SetCategoryCreatedUpdatedTimeNow(category *m.Category) m.Category {
	category.CreatedAt = TimeNow
	category.UpdatedAt = TimeNow
	return *category
}
func SetCategoryUpdatedTimeNow(category *m.Category) m.Category {
	category.UpdatedAt = TimeNow
	return *category
}
func FormatTimeResCategory(category *m.Category) m.Category {
	category.CreatedAt = FormattedTime(category.CreatedAt)
	category.UpdatedAt = FormattedTime(category.UpdatedAt)
	return *category
}

func FormattedTime(ts string) string {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Println(err)
		return ""
	}

	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}
