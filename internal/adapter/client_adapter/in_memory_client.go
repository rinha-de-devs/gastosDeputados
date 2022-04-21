package client_adapter

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

func (c *inMemoryClient) SearchDeputySlice() ([]domain.Deputy, error) {
	return c.deputies, nil
}

func (c *inMemoryClient) ScrappingDeputies(deputies []domain.Deputy) ([]domain.Deputy, error) {
	return deputies, nil
}
