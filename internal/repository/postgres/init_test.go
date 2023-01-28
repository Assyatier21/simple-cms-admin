package postgres

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

func TestNewRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, _, _ := sqlmock.New()

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepository(tt.args.db)
			_, ok := got.(RepositoryHandler)
			if !ok {
				t.Errorf("Not Repository interface")
			}
		})
	}
}
