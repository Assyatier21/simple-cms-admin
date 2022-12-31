package api

import (
	mock_repo "cms-admin/mock/repository/postgres"
	m "cms-admin/models"
	"cms-admin/utils"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func Test_handler_GetCategoryTree(t *testing.T) {
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
				path:   "/categories",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryTree(gomock.Any()).Return(
					[]m.Category{
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
					}, nil)
			},
		},
		{
			name: "sql no rows error",
			args: args{
				method: http.MethodGet,
				path:   "/categories",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryTree(gomock.Any()).Return(
					[]m.Category{}, utils.ErrNotFound)
			},
		},
		{
			name: "repository error",
			args: args{
				method: http.MethodGet,
				path:   "/categories",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryTree(gomock.Any()).Return(
					[]m.Category{
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

			if err := h.GetCategoryTree(c); err != nil {
				t.Errorf("handler.GetArticles() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_GetCategoryDetails(t *testing.T) {
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
				path:   "/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01 20:29:00",
					UpdatedAt: "2022-12-01 20:29:00",
				}, nil)
			},
		},
		{
			name: "error id not an integer",
			args: args{
				method: http.MethodGet,
				path:   "/category?id=not_integer",
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
				path:   "/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(m.Category{}, utils.ErrNotFound)
			},
		},
		{
			name: "repository error",
			args: args{
				method: http.MethodGet,
				path:   "/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().GetCategoryDetails(gomock.Any(), 1).Return(m.Category{
					Id:        1,
					Title:     "category 1",
					Slug:      "category-1",
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

			if err := h.GetCategoryDetails(c); err != nil {
				t.Errorf("handler.GetCategoryDetails() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_InsertCategory(t *testing.T) {
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
				method: http.MethodPost,
				path:   "/admin/v1/category?title=NewCategory&slug=new-category",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().InsertCategory(gomock.Any(), gomock.Any()).
					Return(m.Category{Id: 1, Title: "NewCategory", Slug: "new-category"}, nil)
			},
		},
		{
			name: "error empty title",
			args: args{
				method: http.MethodPost,
				path:   "/admin/v1/category?title=&slug=new-category",
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
				path:   "/admin/v1/category?title=NewCategory&slug=",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "internal server error",
			args: args{
				method: http.MethodPost,
				path:   "/admin/v1/category?title=NewCategory&slug=new-category",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().InsertCategory(gomock.Any(), gomock.Any()).
					Return(m.Category{Id: 1, Title: "NewCategory", Slug: "new-category"}, errors.New("internal server error"))
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

			if err := h.InsertCategory(c); err != nil {
				t.Errorf("handler.InsertCategory() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_UpdateCategory(t *testing.T) {
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
				method: http.MethodPatch,
				path:   "/admin/v1/category?id=1&title=NewCategory&slug=new-category",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).
					Return(m.Category{Id: 1, Title: "NewCategory", Slug: "new-category"}, nil)
			},
		},
		{
			name: "error id not an integer",
			args: args{
				method: http.MethodPatch,
				path:   "/admin/v1/category?id=not_number&title=NewCategory&slug=new-category",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "success with empty title",
			args: args{
				method: http.MethodPatch,
				path:   "/admin/v1/category?id=1&title=&slug=new-category",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).
					Return(m.Category{Id: 1, Title: "NewCategory", Slug: "new-category"}, nil)
			},
		},
		{
			name: "success with empty slug",
			args: args{
				method: http.MethodPatch,
				path:   "/admin/v1/category?id=1&title=NewCategory&slug=",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).
					Return(m.Category{Id: 1, Title: "NewCategory", Slug: "new-category"}, nil)
			},
		},
		{
			name: "error wrong slug format",
			args: args{
				method: http.MethodPatch,
				path:   "/admin/v1/category?id=1&title=NewCategory&slug=###",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "no rows affected error",
			args: args{
				method: http.MethodPatch,
				path:   "/admin/v1/category?id=1&title=NewCategory&slug=new-category",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).
					Return(m.Category{}, utils.NoRowsAffected)
			},
		},
		{
			name: "repository error",
			args: args{
				method: http.MethodPatch,
				path:   "/admin/v1/category?id=1&title=NewCategory&slug=new-category",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).
					Return(m.Category{Id: 1, Title: "NewCategory", Slug: "new-category"}, errors.New("repository error"))
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

			if err := h.UpdateCategory(c); err != nil {
				t.Errorf("handler.UpdateCategory() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_DeleteCategory(t *testing.T) {
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
				path:   "/admin/v1/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().DeleteCategory(gomock.Any(), 1).Return(nil)
			},
		},
		{
			name: "error id not number",
			args: args{
				method: http.MethodDelete,
				path:   "/admin/v1/category?id=not_number",
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
				path:   "/admin/v1/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().DeleteCategory(gomock.Any(), 1).Return(utils.NoRowsAffected)
			},
		},
		{
			name: "internal server error",
			args: args{
				method: http.MethodDelete,
				path:   "/admin/v1/category?id=1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().DeleteCategory(gomock.Any(), 1).Return(errors.New("repository error"))
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

			if err := h.DeleteCategory(c); err != nil {
				t.Errorf("handler.DeleteCategory() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
