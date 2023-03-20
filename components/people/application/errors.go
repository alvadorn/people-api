package application

import (
	"errors"
	"log"

	"github.com/alvadorn/people-api/components/people/repository"
)

var (
	NotFoundErr     = errors.New("not_found_error")
	UnexpectedError = errors.New("unexpected_error")
)

func (uc *useCases) handleError(err error) error {
	log.Println("Error happened:", err.Error())

	if errors.Is(err, repository.NotFoundErr) {
		return NotFoundErr
	}
	return UnexpectedError
}
