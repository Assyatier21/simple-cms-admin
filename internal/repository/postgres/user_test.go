package postgres

import (
	m "cms-admin/models"
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repository_GetUserRegistry(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx      context.Context
		phone    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    m.User
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx:      ctx,
				phone:    "08123",
				password: "secretpass",
			},
			want:    m.User{Phone: "08123", Name: "Muhammad Sholeh", Role: "Admin", Password: "randString"},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"phone", "name", "role", "password"}).
					AddRow("08123", "Muhammad Sholeh", "Admin", "randString")

				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT phone, name, role, password 
                        FROM user WHERE phone = ?`)).WillReturnRows(rows)
			},
		},
		{
			name: "error repository",
			args: args{
				ctx:      ctx,
				phone:    "08123",
				password: "secretpass",
			},
			want:    m.User{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT phone, name, role, password 
				FROM efishery_users WHERE phone = ?`)).WillReturnError(errors.New("repository error"))
			},
		},
		{
			name: "error scan",
			args: args{
				ctx:      ctx,
				phone:    "08123",
				password: "secretpass",
			},
			want:    m.User{},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"phone", "name", "role", "password"}).
					AddRow(nil, "Muhammad Sholeh", nil, 1).RowError(1, errors.New("scanErr"))
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT phone, name, role, password 
						FROM efishery_users WHERE phone = ?`)).WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}

			got, err := r.GetUserRegistry(tt.args.ctx, tt.args.phone, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUserRegistry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetUserRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}
