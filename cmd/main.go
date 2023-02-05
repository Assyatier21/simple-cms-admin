package main

import (
	repo "cms-admin/database"
	"cms-admin/internal/delivery/api"
	"cms-admin/internal/repository/postgres"
	"cms-admin/internal/usecase"
	"cms-admin/routes"
	"fmt"
)

func main() {
	db := repo.Init()
	es := repo.InitElasticClient()

	repository := postgres.NewRepository(db)
	usecase := usecase.NewUsecase(repository, es)
	delivery := api.NewHandler(usecase)
	echo := routes.GetRoutes(delivery)

	host := fmt.Sprintf("%s:%s", "127.0.0.1", "8800")
	_ = echo.Start(host)
}
