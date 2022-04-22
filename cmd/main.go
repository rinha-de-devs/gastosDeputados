package main

import (
	"deputySpending/internal/adapter/client/scrapping_adapter"
	"deputySpending/internal/adapter/repository/memory_repository"
	"deputySpending/internal/service"
)

func main() {

	service.New(memory_repository.New(), scrapping_adapter.New()).SearchExpendDeputy()

}
