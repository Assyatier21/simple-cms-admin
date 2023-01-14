package api

import (
	mock_repo "cms-admin/mock/repository/postgres"
	m "cms-admin/models"
	"cms-admin/utils"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func Test_handler_GetArticles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repo.NewMockRepository(ctrl)

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
				mockRepository.EXPECT().GetArticles(gomock.Any(), 100, 0).Return(
					[]m.ResArticle{
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
								Title: "category 2",
								Slug:  "category-2",
							},
							CreatedAt: "2022-12-01 20:29:00",
							UpdatedAt: "2022-12-01 20:29:00",
						},
					}, nil)
			},
		},
		{
			name: "success with defined limit",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=5",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetArticles(gomock.Any(), 5, 0).Return(
					[]m.ResArticle{
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
								Title: "category 2",
								Slug:  "category-2",
							},
							CreatedAt: "2022-12-01 20:29:00",
							UpdatedAt: "2022-12-01 20:29:00",
						},
					}, nil)
			},
		},
		{
			name: "success with defined offset",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=5&offset=0",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetArticles(gomock.Any(), 5, 0).Return(
					[]m.ResArticle{
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
								Title: "category 2",
								Slug:  "category-2",
							},
							CreatedAt: "2022-12-01 20:29:00",
							UpdatedAt: "2022-12-01 20:29:00",
						},
					}, nil)
			},
		},
		{
			name: "error limit not an integer",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=not_integer&offset=0",
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
			name: "error no data found",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=5&offset=0",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetArticles(gomock.Any(), 5, 0).Return([]m.ResArticle{}, utils.ErrNotFound)
			},
		},
		{
			name: "error repository",
			args: args{
				method: http.MethodGet,
				path:   "/articles?limit=5&offset=0",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().GetArticles(gomock.Any(), 5, 0).Return(nil, errors.New("repository error"))
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
				repository: mockRepository,
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

	mockRepository := mock_repo.NewMockRepository(ctrl)

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
				mockRepository.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(m.ResArticle{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
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
				mockRepository.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(m.ResArticle{}, utils.ErrNotFound)
			},
		},
		{
			name: "repository error",
			args: args{
				method: http.MethodGet,
				path:   "/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(m.ResArticle{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, errors.New("repository error"))
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
				repository: mockRepository,
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

	mockRepository := mock_repo.NewMockRepository(ctrl)

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
					values.Add("title", "title1")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().InsertArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{
					Id:          1,
					Title:       "title1",
					Slug:        "article-1",
					HtmlContent: "<p>article1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "error empty title",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
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
					values.Add("title", "new title")
					values.Add("slug", "")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
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
			name: "error empty html_content",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "new title")
					values.Add("slug", "new-slug")
					values.Add("html_content", "")
					values.Add("category_id", "1")
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
			name: "error empty category_id",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "new title")
					values.Add("slug", "new-slug")
					values.Add("html_content", "<p>This is article content</p>")
					values.Add("category_id", "")
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
					values.Add("title", "new title")
					values.Add("slug", "new-slug")
					values.Add("html_content", "<p>This is article content</p>")
					values.Add("category_id", "1")
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
			name: "error unmarshal metadata",
			args: args{
				method: http.MethodPost,
				path: func() string {
					metaString := `{
						meta_ title:"Test Article",
						meta_ description:"This is a test article",
						meta_ author:"Test Author",
						meta_ keywords:
						[
							"test",
							"article"
						],
						meta_ robots:
						[	
							"index",
							"follow"
						]
					}`
					values := url.Values{}
					values.Add("title", "new title")
					values.Add("slug", "new-slug")
					values.Add("html_content", "<p>This is article content</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metaString)

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
			name: "repository error",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("title", "title1")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().InsertArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{
					Id:          1,
					Title:       "title1",
					Slug:        "article-1",
					HtmlContent: "<p>article1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, errors.New("repository error"))
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
				repository: mockRepository,
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

	mockRepository := mock_repo.NewMockRepository(ctrl)

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
					values.Add("title", "title1")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{
					Id:          1,
					Title:       "title1",
					Slug:        "article-1",
					HtmlContent: "<p>article1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "Success with empty title",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{
					Id:          1,
					Title:       "",
					Slug:        "article-1",
					HtmlContent: "<p>article1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "Success with empty slug",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title1")
					values.Add("slug", "")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{
					Id:          1,
					Title:       "title1",
					Slug:        "",
					HtmlContent: "<p>article1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "Success with empty html_content",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title1")
					values.Add("slug", "title-1")
					values.Add("html_content", "")
					values.Add("category_id", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{
					Id:          1,
					Title:       "title1",
					Slug:        "article-1",
					HtmlContent: "",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "Success with empty category_id",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title1")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{
					Id:          1,
					Title:       "title1",
					Slug:        "article-1",
					HtmlContent: "<p>article1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "Success with empty metadata",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title1")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
					values.Add("metadata", "")

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{
					Id:          1,
					Title:       "title1",
					Slug:        "article-1",
					HtmlContent: "<p>article1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "error empty id",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "")
					values.Add("title", "title1")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
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
			name: "error wrong slug format",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title1")
					values.Add("slug", "Title@#$%^&(!+-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "1")
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
			name: "error category_id not integer",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "title1")
					values.Add("slug", "title-1")
					values.Add("html_content", "<p>article1</p>")
					values.Add("category_id", "not_number")
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
			name: "error unmarshal metadata",
			args: args{
				method: http.MethodPatch,
				path: func() string {
					metaString := `{
						meta_ title:"Test Article",
						meta_ description:"This is a test article",
						meta_ author:"Test Author",
						meta_ keywords:
						[
							"test",
							"article"
						],
						meta_ robots:
						[	
							"index",
							"follow"
						]
					}`
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "new title")
					values.Add("slug", "new-slug")
					values.Add("html_content", "<p>This is article content</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metaString)

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
			name: "error data not found",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("id", "20")
					values.Add("title", "new title")
					values.Add("slug", "new-slug")
					values.Add("html_content", "<p>This is article content</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusNotFound,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{}, utils.ErrNotFound)
			},
		},
		{
			name: "error repository",
			args: args{
				method: http.MethodPost,
				path: func() string {
					values := url.Values{}
					values.Add("id", "1")
					values.Add("title", "new title")
					values.Add("slug", "new-slug")
					values.Add("html_content", "<p>This is article content</p>")
					values.Add("category_id", "1")
					values.Add("metadata", metadataString)

					urlPath := fmt.Sprintf("/admin/v1/article?%s", values.Encode())
					return urlPath
				},
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(m.ResArticle{}, errors.New("repository error"))
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
				repository: mockRepository,
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

	mockRepository := mock_repo.NewMockRepository(ctrl)

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
				path:   "/admin/v1/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().DeleteArticle(gomock.Any(), 1).Return(nil)
			},
		},
		{
			name: "error id not number",
			args: args{
				method: http.MethodDelete,
				path:   "/admin/v1/article?id=not_number",
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
				path:   "/admin/v1/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().DeleteArticle(gomock.Any(), 1).Return(utils.NoRowsAffected)
			},
		},
		{
			name: "internal server error",
			args: args{
				method: http.MethodDelete,
				path:   "/admin/v1/article?id=1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().DeleteArticle(gomock.Any(), 1).Return(errors.New("repository error"))
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
				repository: mockRepository,
			}

			if err := h.DeleteArticle(c); err != nil {
				t.Errorf("handler.DeleteArticle() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
