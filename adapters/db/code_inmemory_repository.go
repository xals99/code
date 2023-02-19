package db

import (
	"github.com/xals99/code/domain/model"
	"fmt"
)

type InMemoryCodeRepository struct {
	codes map[string]model.Code
}

func NewInMemoryCodeRepository() *InMemoryCodeRepository {
	return &InMemoryCodeRepository{codes: map[string]model.Code{}}
}

func (icr InMemoryCodeRepository) Get(code string) (*model.Code, error) {
	codeDb, exists := icr.codes[code]

	if !exists {
		return nil, fmt.Errorf("code not found: %s", code)
	}

	return &codeDb, nil
}

func (icr InMemoryCodeRepository) Set(code model.Code) {
	icr.codes[code.Code] = code
}
