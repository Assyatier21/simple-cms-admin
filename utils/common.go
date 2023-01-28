package utils

import (
	m "cms-admin/models"
	"log"
	"time"
)

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
