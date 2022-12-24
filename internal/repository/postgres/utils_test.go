package postgres

import (
	m "cms-admin/models"
	"testing"
)

func TestFormatTimeResArticle(t *testing.T) {
	resArticle := &m.ResArticle{
		Id:          1,
		Title:       "title 1",
		Slug:        "article-1",
		HtmlContent: "<p> this is article 1</p>",
		ResCategory: m.ResCategory{
			Id:    1,
			Title: "category 1",
			Slug:  "category-1",
		},
		MetaData: m.MetaData{
			Title:       "metatitle 1",
			Description: "metadescription 1",
			Author:      "muhammad sholeh",
			Keywords: []string{
				"description", "testing1",
			},
			Robots: []string{
				"following", "no-index",
			},
		},
		CreatedAt: "2022-12-01T20:29:00Z",
		UpdatedAt: "2022-12-01T20:29:00Z",
	}

	formattedArticle := FormatTimeResArticle(resArticle)

	if formattedArticle.CreatedAt != "2022-12-01 20:29:00" {
		t.Errorf("Expected CreatedAt to be '2022-12-01 20:29:00', got %s", formattedArticle.CreatedAt)
	}
	if formattedArticle.UpdatedAt != "2022-12-01 20:29:00" {
		t.Errorf("Expected UpdatedAt to be '2022-12-01 20:29:00', got %s", formattedArticle.UpdatedAt)
	}
}

func TestFormatTimeResCategory(t *testing.T) {
	resCategory := &m.Category{
		Id:        1,
		Title:     "category 1",
		Slug:      "category-1",
		CreatedAt: "2022-12-01T20:29:00Z",
		UpdatedAt: "2022-12-01T20:29:00Z",
	}

	formattedCategory := FormatTimeResCategory(resCategory)

	if formattedCategory.CreatedAt != "2022-12-01 20:29:00" {
		t.Errorf("Expected CreatedAt to be '2022-12-01 20:29:00', got %s", formattedCategory.CreatedAt)
	}
	if formattedCategory.UpdatedAt != "2022-12-01 20:29:00" {
		t.Errorf("Expected UpdatedAt to be '2022-12-01 20:29:00', got %s", formattedCategory.UpdatedAt)
	}
}
