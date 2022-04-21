package repository

import "deputySpending/internal/domain"

type Repository interface {
	SaveDeputy(domain.Deputy) (domain.Deputy, error)
}
