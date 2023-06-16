package postgres

import (
	DB_QUERY "cms-admin/database/queries"

	m "cms-admin/models"
	"context"
	"database/sql"
	"log"
)

func (r *repository) GetUserRegistry(ctx context.Context, phone string, password string) (m.User, error) {
	var (
		user = m.User{}
		rows *sql.Rows
		err  error
	)

	rows, err = r.db.Query(DB_QUERY.GET_USER, phone)
	if err != nil {
		log.Printf("[Repository][MySQL][GetUserRegistry] failed to get user registry, err: %v\n", err)
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&user.Phone, &user.Name, &user.Role, &user.Password); err != nil {
			log.Printf("[Repository][MySQL][GetUserRegistry] failed while scanning user, err: %v\n", err)
			return user, err
		}
	}

	return user, nil
}
