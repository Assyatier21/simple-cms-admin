package postgres

import (
	m "cms-admin/models"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repository_GetArticles(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    []m.ResArticle
		wantErr bool
		mock    func()
	}{
		{
			name: "scan error",
			args: args{
				ctx:    ctx,
				limit:  5,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				// usecase: id type changed from int to string
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "category_id", "category_title", "category_slug", "created_at", "updated_at"}).
					AddRow("id not number", "title 1", "article-1", "<p> this is article 1</p>", 1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z").
					AddRow("id not number", "title 2", "article-2", "<p> this is article 2</p>", 2, "category 2", "category-2", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id ORDER BY a.id LIMIT $1 OFFSET $2`)).WillReturnRows(rows)

			},
		},
		{
			name: "scan metadata error",
			args: args{
				ctx:    ctx,
				limit:  5,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "category_id", "category_title", "category_slug", "created_at", "updated_at"}).
					AddRow(1, "title 1", "article-1", "<p> this is article 1</p>", 1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id ORDER BY a.id LIMIT $1 OFFSET $2`)).WillReturnRows(rows)
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WithArgs(1).WillReturnError(errors.New("failed to scan"))
			},
		},
		{
			name: "success",
			args: args{
				ctx:    ctx,
				limit:  5,
				offset: 0,
			},
			want: []m.ResArticle{
				{
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
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				},
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "category_id", "category_title", "category_slug", "created_at", "updated_at"}).
					AddRow(1, "title 1", "article-1", "<p> this is article 1</p>", 1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id ORDER BY a.id LIMIT $1 OFFSET $2`)).WillReturnRows(rows)

				temp := m.MetaData{
					Title:       "metatitle 1",
					Description: "metadescription 1",
					Author:      "muhammad sholeh",
					Keywords: []string{
						"description", "testing1",
					},
					Robots: []string{
						"following", "no-index",
					},
				}

				tempMetaData, _ := json.Marshal(temp)
				metadata := sqlMock.NewRows([]string{"metadata"}).
					AddRow(tempMetaData)
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WithArgs(1).WillReturnRows(metadata)

			},
		},
		{
			name: "success with empty article",
			args: args{
				ctx:    ctx,
				limit:  5,
				offset: 0,
			},
			want:    []m.ResArticle{},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "category_id", "category_title", "category_slug", "created_at", "updated_at"})
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id ORDER BY a.id LIMIT $1 OFFSET $2`)).WillReturnRows(rows)
			},
		},
		{
			name: "sql error no rows",
			args: args{
				ctx:    ctx,
				limit:  5,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id ORDER BY a.id LIMIT $1 OFFSET $2`)).WillReturnError(sql.ErrNoRows)
			},
		},
		{
			name: "query error",
			args: args{
				ctx:    ctx,
				limit:  5,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT d.id, d.title, d.slug, d.html_content, c.id, c.title , c.slug, d.created_at, d.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id ORDER BY a.id LIMIT $1 OFFSET $2`)).WillReturnError(errors.New("query error"))

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.GetArticles(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_repository_GetArticleDetails(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    m.ResArticle
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: m.ResArticle{
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
				CreatedAt: "2022-12-01 20:29:00",
				UpdatedAt: "2022-12-01 20:29:00",
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "category_id", "category_title", "category_slug", "created_at", "updated_at"}).
					AddRow(1, "title 1", "article-1", "<p> this is article 1</p>", 1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id WHERE a.id = $1`)).WillReturnRows(rows)

				temp := m.MetaData{
					Title:       "metatitle 1",
					Description: "metadescription 1",
					Author:      "muhammad sholeh",
					Keywords: []string{
						"description", "testing1",
					},
					Robots: []string{
						"following", "no-index",
					},
				}

				tempMetaData, _ := json.Marshal(temp)

				metadata := sqlMock.NewRows([]string{"metadata"}).
					AddRow(tempMetaData)
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT metadata FROM cms_article WHERE id = $1`)).WillReturnRows(metadata)
			},
		},
		{
			name: "failed get metadata",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    m.ResArticle{},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "category_id", "category_title", "category_slug", "created_at", "updated_at"}).
					AddRow(1, "title 1", "article-1", "<p> this is article 1</p>", 1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id WHERE a.id = $1`)).WillReturnRows(rows)
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT metadata FROM cms_article WHERE id = $1`)).WillReturnError(errors.New("failed to get metadata"))
			},
		},
		{
			name: "scan error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    m.ResArticle{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id WHERE a.id = $1`)).WillReturnError(errors.New("error while scanning"))
			},
		},
		{
			name: "sql no rows error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    m.ResArticle{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id WHERE a.id = $1`)).WillReturnError(sql.ErrNoRows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.GetArticleDetails(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetArticleDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetArticleDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
