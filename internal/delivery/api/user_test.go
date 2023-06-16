package api

import (
	mock_usecase "cms-admin/mock/usecase"
	"cms-admin/models"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func Test_handler_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)

	type args struct {
		method string
		path   string
		body   string
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
				path:   "/v1/login",
				body:   `{"phone": "08123", "password": "rand"}`,
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := models.User{
					Phone:    "08123",
					Name:     "Sholeh",
					Role:     "admin",
					Password: "secret_pass",
				}
				req := models.GetUserReq{Phone: "08123", Password: "rand"}
				mockUsecase.EXPECT().Login(gomock.Any(), req.Phone, req.Password).
					Return(data, nil)
			},
		},
		{
			name: "error binding",
			args: args{
				method: http.MethodPost,
				path:   "/v1/login",
				body:   `{"phone": "08123", "password": 1234}`,
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {
			},
		},
		{
			name: "error empty field",
			args: args{
				method: http.MethodPost,
				path:   "/v1/login",
				body:   `{"phone": "08123", "password": ""}`},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {
			},
		},
		{
			name: "error no sql result",
			args: args{
				method: http.MethodPost,
				path:   "/v1/login",
				body:   `{"phone": "08123", "password": "rand"}`,
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := models.User{
					Phone:    "08123",
					Name:     "Sholeh",
					Role:     "admin",
					Password: "secret_pass",
				}
				req := models.GetUserReq{Phone: "08123", Password: "rand"}
				mockUsecase.EXPECT().Login(gomock.Any(), req.Phone, req.Password).
					Return(data, sql.ErrNoRows)
			},
		},
		{
			name: "error repository",
			args: args{
				method: http.MethodPost,
				path:   "/v1/login",
				body:   `{"phone": "08123", "password": "rand"}`,
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				data := models.User{
					Phone:    "08123",
					Name:     "Sholeh",
					Role:     "admin",
					Password: "secret_pass",
				}
				req := models.GetUserReq{Phone: "08123", Password: "rand"}
				mockUsecase.EXPECT().Login(gomock.Any(), req.Phone, req.Password).
					Return(data, errors.New("repository error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, strings.NewReader(tt.args.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				usecase: mockUsecase,
			}

			if err := h.LoginUser(c); err != nil {
				t.Errorf("handler.LoginUser() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
