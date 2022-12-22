package postgres

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
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
			got := New(tt.args.db)
			_, ok := got.(Repository)
			if !ok {
				t.Errorf("Not Repository interface")
			}
		})
	}
}
