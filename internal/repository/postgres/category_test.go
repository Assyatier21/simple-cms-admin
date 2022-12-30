package postgres

import (
	m "cms-admin/models"
	"context"
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repository_GetCategoryTree(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []m.Category
		wantErr bool
		mock    func()
	}{
		{
			name: "scan error",
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow("WRONG TYPE ID", "category 1", "category-1", "2022-12-01 20:29:00", "2022-12-01 20:29:00").
					AddRow("WRONG TYPE ID", "category 2", "category-2", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category ORDER BY id`)).WillReturnRows(rows)
			},
		},
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: []m.Category{
				{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				},
				{
					Id:        2,
					Title:     "category 2",
					Slug:      "category-2",
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				},
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow(1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z").
					AddRow(2, "category 2", "category-2", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category ORDER BY id`)).WillReturnRows(rows)
			},
		},
		{
			name: "success with empty categories",
			args: args{
				ctx: ctx,
			},
			want:    []m.Category{},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"})
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category ORDER BY id`)).WillReturnRows(rows)
			},
		},
		{
			name: "sql no rows error",
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnError(sql.ErrNoRows)
			},
		},
		{
			name: "query error",
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnError(errors.New("query error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.GetCategoryTree(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetCategoryTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetCategoryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetCategoryDetails(t *testing.T) {
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
		want    m.Category
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: m.Category{
				Id:        1,
				Title:     "category 1",
				Slug:      "category-1",
				CreatedAt: "2022-12-01 20:29:00",
				UpdatedAt: "2022-12-01 20:29:00",
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow(1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category WHERE id = $1`)).WillReturnRows(rows)
			},
		},
		{
			name: "sql no rows error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    m.Category{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category WHERE id = $1`)).WillReturnError(sql.ErrNoRows)
			},
		},
		{
			name: "scan error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    m.Category{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM cms_category WHERE id = $1`)).WillReturnError(errors.New("error while scanning"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &repository{
				db: db,
			}
			got, err := r.GetCategoryDetails(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetCategoryDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetCategoryDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_InsertCategory(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx      context.Context
		category m.Category
	}
	tests := []struct {
		name    string
		args    args
		want    m.Category
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				category: m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				},
			},
			want: m.Category{
				Id:        1,
				Title:     "category 1",
				Slug:      "category-1",
				CreatedAt: "2022-12-01 20:29:00",
				UpdatedAt: "2022-12-01 20:29:00",
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow(1, "category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO cms_category`)).WillReturnRows(rows)
			},
		},
		{
			name: "query error",
			args: args{
				ctx: ctx,
				category: m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				},
			},
			want:    m.Category{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(`UPDATE INTO cms_category`).WillReturnError(errors.New("query error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &repository{
				db: db,
			}
			got, err := r.InsertCategory(tt.args.ctx, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.InsertCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.InsertCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_UpdateCategory(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx      context.Context
		category m.Category
	}
	tests := []struct {
		name    string
		args    args
		want    m.Category
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				category: m.Category{
					Id:        1,
					Title:     "new category 1",
					Slug:      "new-category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				},
			},
			want: m.Category{
				Id:        1,
				Title:     "new category 1",
				Slug:      "new-category-1",
				CreatedAt: "2022-12-01 20:29:00",
				UpdatedAt: "2022-12-01 20:29:00",
			},
			wantErr: false,
			mock: func() {
				sqlMock.ExpectExec("UPDATE cms_category").WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "no rows affected",
			args: args{
				ctx: ctx,
				category: m.Category{
					Id:        1,
					Title:     "new category 1",
					Slug:      "new-category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				},
			},
			want:    m.Category{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec("UPDATE cms_category").WillReturnResult(sqlmock.NewResult(1, 0))
			},
		},
		{
			name: "query error",
			args: args{
				ctx: ctx,
				category: m.Category{
					Id:        1,
					Title:     "new category 1",
					Slug:      "new-category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				},
			},
			want:    m.Category{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec("UPDATE cms_category").WillReturnError(errors.New("query erro"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.UpdateCategory(tt.args.ctx, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.UpdateCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_repository_DeleteCategory(t *testing.T) {
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
		wantErr bool
		mock    func()
	}{
		{
			name: "no rows affected",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(`DELETE FROM cms_category`).
					WillDelayFor(time.Second).
					WillReturnResult(sqlmock.NewResult(int64(1), int64(0)))
			},
		},
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: false,
			mock: func() {
				sqlMock.ExpectExec(`DELETE FROM cms_category`).
					WillDelayFor(time.Second).
					WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))
			},
		},
		{
			name: "query error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(`INSERT FROM cms_category`).
					WillDelayFor(time.Second).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("query error")))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			if err := r.DeleteCategory(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
