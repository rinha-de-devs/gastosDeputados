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
				Nome:                      "DEPUTADO",
				Partido:                   "PARTIDO",
				Estado:                    "MG",
				Cota:                      "R$ 341.450,44",
				VerbaDeGabineteDisponivel: "19.121,73",
				PorcentagemDisponivel:     "1,43%",
				VerbaDeGabineteGasto:      "1.320.985,35",
				PorcentagemGasto:          "98,57%",
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
