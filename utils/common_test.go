package utils

import (
	m "cms-admin/models"
	"testing"
	"time"
)

func TestSetArticleCreatedUpdatedTimeNow(t *testing.T) {
	article := &m.Article{Title: "Test Article"}
	now := time.Now().Format("2006-01-02T15:04:05Z")

	result := SetArticleCreatedUpdatedTimeNow(article)

	if result.CreatedAt != now {
		t.Errorf("Expected created at time to be %v, but got %v", now, result.CreatedAt)
	}
	if result.UpdatedAt != now {
		t.Errorf("Expected updated at time to be %v, but got %v", now, result.UpdatedAt)
	}
}
func TestSetArticleUpdatedTimeNow(t *testing.T) {
	article := &m.Article{Title: "Test Article", CreatedAt: time.Now().Format("2006-01-02T15:04:05Z")}
	now := time.Now().Format("2006-01-02T15:04:05Z")

	result := SetArticleUpdatedTimeNow(article)

	if result.UpdatedAt != now {
		t.Errorf("Expected updated at time to be %v, but got %v", now, result.UpdatedAt)
	}
}
func TestFormatTimeResArticle(t *testing.T) {
	// setup
	createdAt := time.Now().Format("2006-01-02T15:04:05Z")
	updatedAt := time.Now().Format("2006-01-02T15:04:05Z")
	article := &m.ResArticle{Title: "Test Article", CreatedAt: createdAt, UpdatedAt: updatedAt}

	// run test
	result := FormatTimeResArticle(article)

	// asserts
	expectedCreatedAt := FormattedTime(createdAt)
	expectedUpdatedAt := FormattedTime(updatedAt)

	if result.CreatedAt != expectedCreatedAt {
		t.Errorf("Expected created at time to be %s, but got %s", expectedCreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != expectedUpdatedAt {
		t.Errorf("Expected updated at time to be %s, but got %s", expectedUpdatedAt, result.UpdatedAt)
	}
}

func TestSetCategoryCreatedUpdatedTimeNow(t *testing.T) {
	category := &m.Category{Title: "Test Category"}
	now := time.Now().Format("2006-01-02T15:04:05Z")

	result := SetCategoryCreatedUpdatedTimeNow(category)

	if result.CreatedAt != now {
		t.Errorf("Expected created at time to be %v, but got %v", now, result.CreatedAt)
	}
	if result.UpdatedAt != now {
		t.Errorf("Expected updated at time to be %v, but got %v", now, result.UpdatedAt)
	}
}
func TestSetCategoryUpdatedTimeNow(t *testing.T) {
	category := &m.Category{Title: "Test Category", CreatedAt: time.Now().Format("2006-01-02T15:04:05Z")}
	now := time.Now().Format("2006-01-02T15:04:05Z")

	result := SetCategoryUpdatedTimeNow(category)

	if result.UpdatedAt != now {
		t.Errorf("Expected updated at time to be %v, but got %v", now, result.UpdatedAt)
	}
}
func TestFormatTimeResCategory(t *testing.T) {
	// setup
	createdAt := time.Now().Format("2006-01-02T15:04:05Z")
	updatedAt := time.Now().Format("2006-01-02T15:04:05Z")
	category := &m.Category{Title: "Test Category", CreatedAt: createdAt, UpdatedAt: updatedAt}

	// run test
	result := FormatTimeResCategory(category)

	// asserts
	expectedCreatedAt := FormattedTime(createdAt)
	expectedUpdatedAt := FormattedTime(updatedAt)

	if result.CreatedAt != expectedCreatedAt {
		t.Errorf("Expected created at time to be %s, but got %s", expectedCreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != expectedUpdatedAt {
		t.Errorf("Expected updated at time to be %s, but got %s", expectedUpdatedAt, result.UpdatedAt)
	}
}

func TestFormattedTime(t *testing.T) {
	ts := time.Now().Format("2006-01-02T15:04:05Z")

	result := FormattedTime(ts)
	expected := time.Now().Format("2006-01-02 15:04:05")
	if result != expected {
		t.Errorf("Expected formatted time to be %s, but got %s", expected, result)
	}

	invalidTs := "invalid time format"
	result = FormattedTime(invalidTs)
	if result != "" {
		t.Errorf("Expected empty string, but got %s", result)
	}
}
