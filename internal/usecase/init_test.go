package usecase

import (
	"cms-admin/internal/repository/postgres"
	mock_repository "cms-admin/mock/repository/postgres"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRepositoryHandler(ctrl)

	type args struct {
		repository postgres.RepositoryHandler
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				repository: mockRepository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUsecase(tt.args.repository)
			_, ok := got.(UsecaseHandler)
			if !ok {
				t.Errorf("Not Delivery Handler interface")
			}
		})
	}
}
