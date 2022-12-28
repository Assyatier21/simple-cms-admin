package utils

import (
	m "cms-admin/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidAlphabet(t *testing.T) {
	result := IsValidAlphabet("alphabet")
	assert.True(t, result)

	result = IsValidAlphabet("123")
	assert.False(t, result)
}

func TestIsValidNumeric(t *testing.T) {
	result := IsValidNumeric("123")
	assert.True(t, result)

	result = IsValidNumeric("invalid format string")
	assert.False(t, result)
}

func TestIsValidAlphaNumeric(t *testing.T) {
	result := IsValidAlphaNumeric("Alpha123")
	assert.True(t, result)

	result = IsValidAlphaNumeric("!@#")
	assert.False(t, result)
}

func TestIsValidSlug(t *testing.T) {
	result := IsValidSlug("valid-number-2-with-hyphen")
	assert.True(t, result)

	result = IsValidSlug("")
	assert.False(t, result)
}

func TestFormattedTime(t *testing.T) {
	result := FormattedTime("2022-12-20T12:34:56Z")
	assert.Equal(t, "2022-12-20 12:34:56", result)

	result = FormattedTime("invalid time string")
	assert.Equal(t, "", result)
}

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
