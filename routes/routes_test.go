package routes

import (
	"cms-admin/internal/delivery/api"
	"cms-admin/internal/repository/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
)

func TestGetRoutes(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := GetRoutes(api.New(postgres.New(db)))

	assert.Equal(t, "/v1/articles", e.Routes()[0].Path)
	assert.Equal(t, "/v1/article", e.Routes()[1].Path)
}
