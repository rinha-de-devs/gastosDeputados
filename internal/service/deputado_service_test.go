package service_test

import (
	"deputySpending/internal/adapter/client/memory_client"
	"deputySpending/internal/adapter/repository/memory_repository"
	"deputySpending/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchExpendDeputy(t *testing.T) {
	service := service.New(memory_repository.New(), memory_client.New())

	deputies := service.SearchExpendDeputy()

	assert.Len(t, deputies, 1)
	assert.Equal(t, "1", deputies[0].ID)
}
