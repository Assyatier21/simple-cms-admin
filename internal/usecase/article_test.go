package usecase

import (
	mock_postgres "cms-admin/mock/repository/postgres"
	m "cms-admin/models"
	"context"
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
