package service

import (
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

func (s *service) SearchExpendDeputy() {
	deputies, err := s.deputadoClient.SearchDeputySlice()
	if err != nil {
		fmt.Printf("ERROR TO SEARCH DEPUTIES SLICE")
		panic(err.Error())
	}

	deps, err := s.deputadoClient.ScrappingDeputies(deputies)
	if err != nil {
		fmt.Printf("ERROR TO SCRAPPING DEPUTIES")
		panic(err.Error())
	}

	for _, dep := range deps {
		_, err := s.deputadoRepository.SaveDeputy(dep)
		if err != nil {
			fmt.Printf("ERROR TO SAVE DEPUTY %s - %s", dep.ID, dep.Nome)
			panic(err.Error())
		}
	}

}
