package api

import (
	mock_usecase "cms-admin/mock/usecase"
	m "cms-admin/models"
	msg "cms-admin/models/lib"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func Test_handler_GetArticles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)

	type args struct {
		method string
		path   string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "success",
			args: args{
				method: http.MethodGet,
				path:   "/articles",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := []m.ResArticle{
					{
						Id:          1,
						Title:       "title 1",
						Slug:        "article-1",
						HtmlContent: "<p> this is article 1</p>",
						ResCategory: m.ResCategory{
							Id:    1,
							Title: "catgegory 1",
							Slug:  "category-1",
						},
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
					{
						Id:          2,
						Title:       "title 2",
						Slug:        "article-2",
						HtmlContent: "<p> this is article 2</p>",
						ResCategory: m.ResCategory{
							Id:    2,
							Title: "catgegory 2",
							Slug:  "category-2",
						},
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}
				var articles []interface{}
				for _, v := range data {
					articles = append(articles, v)
				}
				mockUsecase.EXPECT().GetArticles(gomock.Any(), 100, 0).Return(articles, nil)
			},
		},
		{
			name: "success defined limit",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=5",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := []m.ResArticle{
					{
						Id:          1,
						Title:       "title 1",
						Slug:        "article-1",
						HtmlContent: "<p> this is article 1</p>",
						ResCategory: m.ResCategory{
							Id:    1,
							Title: "catgegory 1",
							Slug:  "category-1",
						},
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
					{
						Id:          2,
						Title:       "title 2",
						Slug:        "article-2",
						HtmlContent: "<p> this is article 2</p>",
						ResCategory: m.ResCategory{
							Id:    2,
							Title: "catgegory 2",
							Slug:  "category-2",
						},
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}
				var articles []interface{}
				for _, v := range data {
					articles = append(articles, v)
				}
				mockUsecase.EXPECT().GetArticles(gomock.Any(), 5, 0).Return(articles, nil)
			},
		},
		{
			name: "success defined offset",
			args: args{
				method: http.MethodGet,
				path:   "/articles?offset=0",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := []m.ResArticle{
					{
						Id:          1,
						Title:       "title 1",
						Slug:        "article-1",
						HtmlContent: "<p> this is article 1</p>",
						ResCategory: m.ResCategory{
							Id:    1,
							Title: "catgegory 1",
							Slug:  "category-1",
						},
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
					{
						Id:          2,
						Title:       "title 2",
						Slug:        "article-2",
						HtmlContent: "<p> this is article 2</p>",
						ResCategory: m.ResCategory{
							Id:    2,
							Title: "catgegory 2",
							Slug:  "category-2",
						},
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}
				var articles []interface{}
				for _, v := range data {
					articles = append(articles, v)
				}
				mockUsecase.EXPECT().GetArticles(gomock.Any(), 100, 0).Return(articles, nil)
			},
		},
		{
			name: "error limit not an integer",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=not_integer",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "error offset not an integer",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=5&offset=not_integer",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "error usecase",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=5&offset=0",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockUsecase.EXPECT().GetArticles(gomock.Any(), 5, 0).Return(nil, errors.New("usecase error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				usecase: mockUsecase,
			}

			if err := h.GetArticles(c); err != nil {
				t.Errorf("handler.GetArticles() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_GetArticleDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)

	type args struct {
		method string
		path   string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "success",
			args: args{
				method: http.MethodGet,
				path:   "/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := []m.ResArticle{
					{
						Id:          1,
						Title:       "title 1",
						Slug:        "article-1",
						HtmlContent: "<p> this is article 1</p>",
						ResCategory: m.ResCategory{
							Id:    1,
							Title: "catgegory 1",
							Slug:  "category-1",
						},
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}
				var article []interface{}
				for _, v := range data {
					article = append(article, v)
				}
				mockUsecase.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(article, nil)
			},
		},
		{
			name: "error id not an integer",
			args: args{
				method: http.MethodGet,
				path:   "/article?id=not_integer",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "sql no rows error",
			args: args{
				method: http.MethodGet,
				path:   "/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				var article []interface{}
				mockUsecase.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(article, sql.ErrNoRows)
			},
		},
		{
			name: "usecase error",
			args: args{
				method: http.MethodGet,
				path:   "/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				data := []m.ResArticle{
					{
						Id:          1,
						Title:       "title 1",
						Slug:        "article-1",
						HtmlContent: "<p> this is article 1</p>",
						ResCategory: m.ResCategory{
							Id:    1,
							Title: "catgegory 1",
							Slug:  "category-1",
						},
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}
				var article []interface{}
				for _, v := range data {
					article = append(article, v)
				}
				mockUsecase.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(article, errors.New("usecase error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				usecase: mockUsecase,
			}

			if err := h.GetArticleDetails(c); err != nil {
				t.Errorf("handler.GetArticleDetails() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_InsertArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)
	metadataString := `{
		"meta_title":"Test Article",
		"meta_description":"This is a test article",
		"meta_author":"Test Author",
		"meta_keywords":
		[
			"test",
			"article"
		],
		"meta_robots":
		[	
			"index",
			"follow"
		]
	}`

	type args struct {
		method string
		path   func() string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "Success",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "title 1")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := m.ResArticle{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "catgegory 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}
				var article []interface{}
				article = append(article, data)

				mockUsecase.EXPECT().InsertArticle(gomock.Any(), "title 1", "article-1", "<p> this is article 1</p>", 1, metadataString).Return(article, nil)
			},
		},
		{
			name: "error empty title",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "error empty slug",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "title")
					values.Add("slug", "")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "error empty htmlcontent",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "title")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "error empty categoryid",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "title")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "error empty metadata",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "title")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", "")

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "error usecase",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "title 1")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockUsecase.EXPECT().InsertArticle(gomock.Any(), "title 1", "article-1", "<p> this is article 1</p>", 1, metadataString).Return(nil, errors.New("error usecase"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path(), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				usecase: mockUsecase,
			}

			if err := h.InsertArticle(c); err != nil {
				t.Errorf("handler.InsertArticle() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_UpdateArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)
	metadataString := `{
		"meta_title":"Test Article",
		"meta_description":"This is a test article",
		"meta_author":"Test Author",
		"meta_keywords":
		[
			"test",
			"article"
		],
		"meta_robots":
		[	
			"index",
			"follow"
		]
	}`

	type args struct {
		method string
		path   func() string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "Success",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title 1")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := m.ResArticle{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "catgegory 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}
				var article []interface{}
				article = append(article, data)

				mockUsecase.EXPECT().UpdateArticle(gomock.Any(), 1, "title 1", "article-1", "<p> this is article 1</p>", 1, metadataString).Return(article, nil)
			},
		},
		{
			name: "Error id not integer",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "not_number")
					values.Add("title", "title 1")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "Error invalid slug",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title 1")
					values.Add("slug", "article - 1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "Error categoryid not integer",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title 1")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "not_number")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "Error sql no rows affected",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title 1")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := m.ResArticle{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "catgegory 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}
				var article []interface{}
				article = append(article, data)

				mockUsecase.EXPECT().UpdateArticle(gomock.Any(), 1, "title 1", "article-1", "<p> this is article 1</p>", 1, metadataString).Return(article, sql.ErrNoRows)
			},
		},
		{
			name: "Error usecase",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title 1")
					values.Add("slug", "article-1")
					values.Add("htmlcontent", "<p> this is article 1</p>")
					values.Add("categoryid", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				data := m.ResArticle{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "catgegory 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}
				var article []interface{}
				article = append(article, data)

				mockUsecase.EXPECT().UpdateArticle(gomock.Any(), 1, "title 1", "article-1", "<p> this is article 1</p>", 1, metadataString).Return(article, errors.New("usecase error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path(), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				usecase: mockUsecase,
			}

			if err := h.UpdateArticle(c); err != nil {
				t.Errorf("handler.UpdateArticle() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_DeleteArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)

	type args struct {
		method string
		path   string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "success",
			args: args{
				method: http.MethodDelete,
				path:   "/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockUsecase.EXPECT().DeleteArticle(gomock.Any(), 1).Return(nil)
			},
		},
		{
			name: "error id not number",
			args: args{
				method: http.MethodDelete,
				path:   "/article?id=not_number",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "no rows affected",
			args: args{
				method: http.MethodDelete,
				path:   "/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockUsecase.EXPECT().DeleteArticle(gomock.Any(), 1).Return(msg.ERROR_NO_ROWS_AFFECTED)
			},
		},
		{
			name: "usecase error",
			args: args{
				method: http.MethodDelete,
				path:   "/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockUsecase.EXPECT().DeleteArticle(gomock.Any(), 1).Return(errors.New("usecase error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				usecase: mockUsecase,
			}

			if err := h.DeleteArticle(c); err != nil {
				t.Errorf("handler.DeleteArticle() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
