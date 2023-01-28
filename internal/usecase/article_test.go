package usecase

import (
	mock_postgres "cms-admin/mock/repository/postgres"
	m "cms-admin/models"
	"cms-admin/utils"
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_usecase_GetArticles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockRepository := mock_postgres.NewMockRepositoryHandler(ctrl)

	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    func() []interface{}
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx:    ctx,
				limit:  10,
				offset: 0,
			},
			want: func() []interface{} {
				data := []m.ResArticle{
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
				}
				var articles []interface{}
				for _, v := range data {
					articles = append(articles, v)
				}
				return articles
			},
			wantErr: false,
			mock: func() {
				data := []m.ResArticle{
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
						CreatedAt: "2022-12-01T20:29:00Z",
						UpdatedAt: "2022-12-01T20:29:00Z",
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
						CreatedAt: "2022-12-01T20:29:00Z",
						UpdatedAt: "2022-12-01T20:29:00Z",
					},
				}
				mockRepository.EXPECT().GetArticles(gomock.Any(), 10, 0).Return(data, nil)
			},
		},
		{
			name: "error repository",
			args: args{
				ctx:    ctx,
				limit:  10,
				offset: 0,
			},
			want: func() []interface{} {
				return nil
			},
			wantErr: true,
			mock: func() {
				data := []m.ResArticle{
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
						CreatedAt: "2022-12-01T20:29:00Z",
						UpdatedAt: "2022-12-01T20:29:00Z",
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
						CreatedAt: "2022-12-01T20:29:00Z",
						UpdatedAt: "2022-12-01T20:29:00Z",
					},
				}
				mockRepository.EXPECT().GetArticles(gomock.Any(), 10, 0).Return(data, errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			got, err := u.GetArticles(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("usecase.GetArticles() = %v, want %v", got, tt.want())
			}
		})
	}
}
func Test_usecase_GetArticleDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockRepository := mock_postgres.NewMockRepositoryHandler(ctrl)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    func() []interface{}
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: func() []interface{} {
				data := []m.ResArticle{
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
				}
				var article []interface{}
				for _, v := range data {
					article = append(article, v)
				}
				return article
			},
			wantErr: false,
			mock: func() {
				data := m.ResArticle{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				}
				mockRepository.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(data, nil)
			},
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: func() []interface{} {
				return nil
			},
			wantErr: true,
			mock: func() {
				data := m.ResArticle{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					ResCategory: m.ResCategory{
						Id:    1,
						Title: "category 1",
						Slug:  "category-1",
					},
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				}
				mockRepository.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(data, errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			got, err := u.GetArticleDetails(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetArticleDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("usecase.GetArticleDetails() = %v, want %v", got, tt.want())
			}
		})
	}
}
func Test_usecase_InsertArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockRepository := mock_postgres.NewMockRepositoryHandler(ctrl)
	metadataString := `{
		"meta_title":"metatitle 1",
		"meta_description":"metadescription 1",
		"meta_author":"muhammad sholeh",
		"meta_keywords":
		[
			"description",
			"testing1"
		],
		"meta_robots":
		[	
			"following",
			"no-index"
		]
	}`

	type args struct {
		ctx         context.Context
		title       string
		slug        string
		htmlcontent string
		categoryid  int
		metadata    string
	}
	tests := []struct {
		name    string
		args    args
		want    func() []interface{}
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx:         ctx,
				title:       "title 1",
				slug:        "article-1",
				htmlcontent: "<p> this is article 1</p>",
				categoryid:  1,
				metadata:    metadataString,
			},
			want: func() []interface{} {
				data := m.ResArticle{
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
					CreatedAt: utils.TimeNow,
					UpdatedAt: utils.TimeNow,
				}

				utils.FormatTimeResArticle(&data)

				var article []interface{}
				article = append(article, data)

				return article
			},
			wantErr: false,
			mock: func() {
				data := m.Article{
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					CategoryID:  1,
					CreatedAt:   utils.TimeNow,
					UpdatedAt:   utils.TimeNow,
				}
				_ = json.Unmarshal([]byte(metadataString), &data.MetaData)

				resArticle := m.ResArticle{
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
					CreatedAt: utils.TimeNow,
					UpdatedAt: utils.TimeNow,
				}
				mockRepository.EXPECT().InsertArticle(gomock.Any(), data).Return(resArticle, nil)
			},
		},
		{
			name: "repository error",
			args: args{
				ctx:         ctx,
				title:       "title 1",
				slug:        "article-1",
				htmlcontent: "<p> this is article 1</p>",
				categoryid:  1,
				metadata:    metadataString,
			},
			want: func() []interface{} {
				return nil
			},
			wantErr: true,
			mock: func() {
				data := m.Article{
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					CategoryID:  1,
					CreatedAt:   utils.TimeNow,
					UpdatedAt:   utils.TimeNow,
				}
				_ = json.Unmarshal([]byte(metadataString), &data.MetaData)

				resArticle := m.ResArticle{
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
					CreatedAt: utils.TimeNow,
					UpdatedAt: utils.TimeNow,
				}
				mockRepository.EXPECT().InsertArticle(gomock.Any(), data).Return(resArticle, errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			got, err := u.InsertArticle(tt.args.ctx, tt.args.title, tt.args.slug, tt.args.htmlcontent, tt.args.categoryid, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.InsertArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("usecase.InsertArticle() = %v, want %v", got, tt.want())
			}
		})
	}
}
func Test_usecase_UpdateArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockRepository := mock_postgres.NewMockRepositoryHandler(ctrl)
	metadataString := `{
		"meta_title":"metatitle 1",
		"meta_description":"metadescription 1",
		"meta_author":"muhammad sholeh",
		"meta_keywords":
		[
			"description",
			"testing1"
		],
		"meta_robots":
		[	
			"following",
			"no-index"
		]
	}`

	type args struct {
		ctx         context.Context
		id          int
		title       string
		slug        string
		htmlcontent string
		categoryid  int
		metadata    string
	}
	tests := []struct {
		name    string
		args    args
		want    func() []interface{}
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx:         ctx,
				id:          1,
				title:       "title 1",
				slug:        "article-1",
				htmlcontent: "<p> this is article 1</p>",
				categoryid:  1,
				metadata:    metadataString,
			},
			want: func() []interface{} {
				data := m.ResArticle{
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
					UpdatedAt: utils.TimeNow,
				}

				utils.FormatTimeResArticle(&data)

				var article []interface{}
				article = append(article, data)

				return article
			},
			wantErr: false,
			mock: func() {
				resArticle := m.ResArticle{
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
					UpdatedAt: utils.FormattedTime(utils.TimeNow),
				}

				mockRepository.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(resArticle, nil)

				data := m.Article{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					CategoryID:  1,
					CreatedAt:   "2022-12-01T20:29:00Z",
					UpdatedAt:   utils.FormattedTime(utils.TimeNow),
				}
				_ = json.Unmarshal([]byte(metadataString), &data.MetaData)

				mockRepository.EXPECT().UpdateArticle(gomock.Any(), data).Return(resArticle, nil)
			},
		},
		{
			name: "repository error",
			args: args{
				ctx:         ctx,
				id:          1,
				title:       "title 1",
				slug:        "article-1",
				htmlcontent: "<p> this is article 1</p>",
				categoryid:  1,
				metadata:    metadataString,
			},
			want: func() []interface{} {
				var articles []interface{}
				return articles
			},
			wantErr: true,
			mock: func() {
				resArticle := m.ResArticle{
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
					UpdatedAt: utils.FormattedTime(utils.TimeNow),
				}

				mockRepository.EXPECT().GetArticleDetails(gomock.Any(), 1).Return(resArticle, nil)

				data := m.Article{
					Id:          1,
					Title:       "title 1",
					Slug:        "article-1",
					HtmlContent: "<p> this is article 1</p>",
					CategoryID:  1,
					CreatedAt:   "2022-12-01T20:29:00Z",
					UpdatedAt:   utils.FormattedTime(utils.TimeNow),
				}
				_ = json.Unmarshal([]byte(metadataString), &data.MetaData)

				mockRepository.EXPECT().UpdateArticle(gomock.Any(), data).Return(resArticle, errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			got, err := u.UpdateArticle(tt.args.ctx, tt.args.id, tt.args.title, tt.args.slug, tt.args.htmlcontent, tt.args.categoryid, tt.args.metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.UpdateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("usecase.UpdateArticle() = %v, want %v", got, tt.want())
			}
		})
	}
}
func Test_usecase_DeleteArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockRepository := mock_postgres.NewMockRepositoryHandler(ctrl)

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
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: false,
			mock: func() {
				mockRepository.EXPECT().DeleteArticle(gomock.Any(), 1).Return(nil)
			},
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				mockRepository.EXPECT().DeleteArticle(gomock.Any(), 1).Return(errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			err := u.DeleteArticle(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
