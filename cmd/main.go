package main

import (
	"deputySpending/internal/adapter/client/scrapping_adapter"
	"deputySpending/internal/adapter/repository/postgres_repository"
	"deputySpending/internal/service"
)

func main() {

	service.New(postgres_repository.New(), scrapping_adapter.New()).SearchExpendDeputy()

}
