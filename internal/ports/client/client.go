package client

import "deputySpending/internal/domain"

type Client interface {
	//SearchDeputySlice
	SearchDeputySlice() ([]domain.Deputy, error)
	// ScrappingDeputies
	ScrappingDeputies([]domain.Deputy) ([]domain.Deputy, error)
}
