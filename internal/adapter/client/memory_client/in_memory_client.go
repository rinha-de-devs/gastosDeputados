package memory_client

import "deputySpending/internal/domain"

type inMemoryClient struct {
	deputies []domain.Deputy
}

func New() *inMemoryClient {
	return &inMemoryClient{
		deputies: []domain.Deputy{
			{
				ID:                        "1",
				Nome:                      "POLITICO",
				Partido:                   "PARTIDO",
				Estado:                    "",
				Cota:                      "",
				VerbaDeGabineteDisponivel: "",
				PorcentagemDisponivel:     "",
				VerbaDeGabineteGasto:      "",
				PorcentagemGasto:          "",
			},
		},
	}
}

func (repo *inMemoryClient) SearchDeputySlice() ([]domain.Deputy, error) {
	return repo.deputies, nil
}

func (repo *inMemoryClient) ScrappingDeputies(deputies []domain.Deputy) ([]domain.Deputy, error) {
	return deputies, nil
}
