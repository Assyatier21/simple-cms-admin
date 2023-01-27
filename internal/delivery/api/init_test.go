package api

import (
	"cms-admin/internal/usecase"
	mock_usecase "cms-admin/mock/usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockUsecaseHandler(ctrl)

	type args struct {
		usecase usecase.UsecaseHandler
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				usecase: mockUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHandler(tt.args.usecase)
			_, ok := got.(DeliveryHandler)
			if !ok {
				t.Errorf("Not Delivery Handler interface")
			}
		})
	}
}
