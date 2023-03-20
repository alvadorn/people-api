package application

import (
	"context"

	"github.com/alvadorn/people-api/components/people/domain"
)

type useCases struct {
	repository domain.PeopleRepository
}

type UseCases interface {
	GetByName(ctx context.Context, name string) (*PersonDTO, error)
}

func New(repository domain.PeopleRepository) *useCases {
	return &useCases{repository: repository}
}
