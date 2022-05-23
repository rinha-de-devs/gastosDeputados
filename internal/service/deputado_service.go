package service

import (
	"deputySpending/internal/domain"
	"deputySpending/internal/ports/client"
	"deputySpending/internal/ports/repository"
	"fmt"
)

type service struct {
	deputadoRepository repository.Repository
	deputadoClient     client.Client
}

func New(deputadoRepository repository.Repository, deputadoClient client.Client) *service {
	return &service{
		deputadoRepository: deputadoRepository,
		deputadoClient:     deputadoClient,
	}
}

func (s *service) SearchExpendDeputy(year string) []domain.Deputy {
	deputies, err := s.deputadoClient.SearchDeputySlice()
	if err != nil {
		fmt.Printf("ERROR TO SEARCH DEPUTIES SLICE")
		panic(err.Error())
	}

	deps, err := s.deputadoClient.ScrappingDeputies(deputies, year)
	if err != nil {
		fmt.Printf("ERROR TO SCRAPPING DEPUTIES")
		panic(err.Error())
	}

	deputySlice := make([]domain.Deputy, len(deps))
	for index, dep := range deps {
		deputySaved, err := s.deputadoRepository.SaveDeputy(dep)
		if err != nil {
			fmt.Printf("ERROR TO SAVE DEPUTY %s - %s", dep.ID, dep.Nome)
			panic(err.Error())
		}
		deputySlice[index] = deputySaved
	}

	return deputySlice

}
