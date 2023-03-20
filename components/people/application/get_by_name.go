package application

import (
	"context"
)

func (uc *useCases) GetByName(ctx context.Context, name string) (*PersonDTO, error) {
	person, err := uc.repository.GetByName(ctx, name)

	if err != nil {
		return nil, uc.handleError(err)
	}

	return parseDomainToDTO(person), nil
}
