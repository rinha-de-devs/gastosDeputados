package client

import "deputySpending/internal/domain"

type Client interface {
	SearchDeputySlice() ([]domain.Deputy, error)
	ScrappingDeputies(deputies []domain.Deputy, year string) ([]domain.Deputy, error)
}
