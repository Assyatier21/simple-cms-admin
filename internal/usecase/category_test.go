package usecase

import (
	mock_postgres "cms-admin/mock/repository/postgres"
	m "cms-admin/models"
	"cms-admin/utils"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_usecase_GetCategoryTree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockRepository := mock_postgres.NewMockRepositoryHandler(ctrl)

	type args struct {
		ctx context.Context
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
			},
			want: func() []interface{} {
				data := []m.Category{
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
				}

				var categories []interface{}
				for _, v := range data {
					categories = append(categories, v)
				}
				return categories
			},
			wantErr: false,
			mock: func() {
				data := []m.Category{
					{
						Id:        1,
						Title:     "category 1",
						Slug:      "category-1",
						CreatedAt: "2022-12-01T20:29:00Z",
						UpdatedAt: "2022-12-01T20:29:00Z",
					},
					{
						Id:        2,
						Title:     "category 2",
						Slug:      "category-2",
						CreatedAt: "2022-12-01T20:29:00Z",
						UpdatedAt: "2022-12-01T20:29:00Z",
					},
				}
				mockRepository.EXPECT().GetCategoryTree(gomock.Any()).Return(data, nil)
			},
		},
		{
			name: "error repository",
			args: args{
				ctx: ctx,
			},
			want: func() []interface{} {
				return nil
			},
			wantErr: true,
			mock: func() {
				mockRepository.EXPECT().GetCategoryTree(gomock.Any()).Return(nil, errors.New("error repository"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			got, err := u.GetCategoryTree(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetCategoryTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("usecase.GetCategoryTree() = %v, want %v", got, tt.want())
			}
		})
	}
}
func Test_usecase_GetCategoryDetails(t *testing.T) {
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
				data := []m.Category{
					{
						Id:        1,
						Title:     "category 1",
						Slug:      "category-1",
						CreatedAt: "2022-12-01 20:29:00",
						UpdatedAt: "2022-12-01 20:29:00",
					},
				}

				var category []interface{}
				for _, v := range data {
					category = append(category, v)
				}
				return category
			},
			wantErr: false,
			mock: func() {
				data := m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				}
				mockRepository.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(data, nil)
			},
		},
		{
			name: "error repository",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: func() []interface{} {
				return nil
			},
			wantErr: true,
			mock: func() {
				data := m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				}
				mockRepository.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(data, errors.New("error repository"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			got, err := u.GetCategoryDetails(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetCategoryDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("usecase.GetCategoryDetails() = %v, want %v", got, tt.want())
			}
		})
	}
}
func Test_usecase_InsertCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockRepository := mock_postgres.NewMockRepositoryHandler(ctrl)

	type args struct {
		ctx   context.Context
		title string
		slug  string
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
				ctx:   ctx,
				title: "category 1",
				slug:  "category-1",
			},
			want: func() []interface{} {
				data := m.Category{
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: utils.TimeNow,
					UpdatedAt: utils.TimeNow,
				}

				utils.FormatTimeResCategory(&data)

				var category []interface{}
				category = append(category, data)

				return category
			},
			wantErr: false,
			mock: func() {
				data := m.Category{
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: utils.TimeNow,
					UpdatedAt: utils.TimeNow,
				}
				mockRepository.EXPECT().InsertCategory(gomock.Any(), data).Return(data, nil)

			},
		},
		{
			name: "error repository",
			args: args{
				ctx:   ctx,
				title: "category 1",
				slug:  "category-1",
			},
			want: func() []interface{} {
				return nil
			},
			wantErr: true,
			mock: func() {
				data := m.Category{
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: utils.TimeNow,
					UpdatedAt: utils.TimeNow,
				}
				mockRepository.EXPECT().InsertCategory(gomock.Any(), data).Return(data, errors.New("repository error"))

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			got, err := u.InsertCategory(tt.args.ctx, tt.args.title, tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.InsertCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("usecase.InsertCategory() = %v, want %v", got, tt.want())
			}
		})
	}
}
func Test_usecase_UpdateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockRepository := mock_postgres.NewMockRepositoryHandler(ctrl)

	type args struct {
		ctx   context.Context
		id    int
		title string
		slug  string
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
				ctx:   ctx,
				id:    1,
				title: "category 1",
				slug:  "category-1",
			},
			want: func() []interface{} {
				data := m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: utils.TimeNow,
				}

				utils.FormatTimeResCategory(&data)

				var category []interface{}
				category = append(category, data)

				return category
			},
			wantErr: false,
			mock: func() {
				resCategory := m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: utils.TimeNow,
				}

				mockRepository.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(resCategory, nil)

				data := m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: utils.FormattedTime(utils.TimeNow),
				}

				mockRepository.EXPECT().UpdateCategory(gomock.Any(), data).Return(data, nil)

			},
		},
		{
			name: "error repository",
			args: args{
				ctx:   ctx,
				id:    1,
				title: "category 1",
				slug:  "category-1",
			},
			want: func() []interface{} {
				var category []interface{}
				return category
			},
			wantErr: true,
			mock: func() {
				resCategory := m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: utils.TimeNow,
				}

				mockRepository.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(resCategory, nil)

				data := m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: utils.FormattedTime(utils.TimeNow),
				}

				mockRepository.EXPECT().UpdateCategory(gomock.Any(), data).Return(data, errors.New("repository error"))

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			got, err := u.UpdateCategory(tt.args.ctx, tt.args.id, tt.args.title, tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("usecase.UpdateCategory() = %v, want %v", got, tt.want())
			}
		})
	}
}
func Test_usecase_DeleteCategory(t *testing.T) {
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
				mockRepository.EXPECT().DeleteCategory(gomock.Any(), 1).Return(nil)
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
				mockRepository.EXPECT().DeleteCategory(gomock.Any(), 1).Return(errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			u := &usecase{
				repository: mockRepository,
			}

			err := u.DeleteCategory(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
