package main

import (
	repo "cms-admin/database"
	"cms-admin/internal/delivery/api"
	elastic "cms-admin/internal/repository/elasticsearch"
	"cms-admin/internal/repository/postgres"
	"cms-admin/internal/usecase"
	"cms-admin/routes"
	"fmt"
)

func main() {
	db := repo.Init()
	es := repo.InitElasticClient()

	repository := postgres.NewRepository(db)
	esRepository := elastic.NewElasticRepository(es)

	usecase := usecase.NewUsecase(repository, esRepository)
	delivery := api.NewHandler(usecase)

	echo := routes.GetRoutes(delivery)
	host := fmt.Sprintf("%s:%s", "127.0.0.1", "8888")
	_ = echo.Start(host)
}
