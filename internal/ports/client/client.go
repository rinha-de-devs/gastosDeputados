package client

import "deputySpending/internal/domain"

type Client interface {
	SearchDeputySlice() ([]domain.Deputy, error)
	ScrappingDeputies([]domain.Deputy) ([]domain.Deputy, error)
}
