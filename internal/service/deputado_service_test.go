package service_test

import (
	"deputySpending/internal/adapter/client_adapter"
	"deputySpending/internal/adapter/repository_adapter"
	"deputySpending/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchExpendDeputy(t *testing.T) {
	service := service.New(repository_adapter.New(), client_adapter.New())

	deputies := service.SearchExpendDeputy()

	assert.Len(t, deputies, 1)
	assert.Equal(t, "1", deputies[0].ID)
}
