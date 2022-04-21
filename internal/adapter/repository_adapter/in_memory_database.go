package repository_adapter

import (
	"deputySpending/internal/domain"
	"encoding/json"
	"fmt"
)

type inMemoryDB struct {
	keyValue map[string][]byte
}

func New() *inMemoryDB {
	return &inMemoryDB{keyValue: map[string][]byte{}}
}

func (database *inMemoryDB) SaveDeputy(deputy domain.Deputy) (domain.Deputy, error) {
	bytes, err := json.Marshal(deputy)
	if err != nil {
		fmt.Printf("ERROR TO SAVE DEPUTY")
		panic(err.Error())
	}

	database.keyValue[deputy.ID] = bytes

	return deputy, nil
}
