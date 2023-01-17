package main

import (
	db "cms-admin/database"
	"cms-admin/internal/delivery/api"
	"cms-admin/internal/repository/postgres"
	"cms-admin/internal/usecase"
	"cms-admin/routes"
	"fmt"
)

func main() {
	db := db.Init()

	repository := postgres.New(db)
	usecase := usecase.New(repository)
	delivery := api.New(usecase)
	echo := routes.GetRoutes(delivery)

	host := fmt.Sprintf("%s:%s", "127.0.0.1", "8800")
	_ = echo.Start(host)
}
