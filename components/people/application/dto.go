package application

import (
	"strings"

	"github.com/alvadorn/people-api/components/people/domain"
)

type PersonDTO struct {
	Name             string  `json:"name"`
	ShortDescription *string `json:"short_description"`
}

func parseDomainToDTO(person *domain.Person) *PersonDTO {
	var description *string
	if person.ShortDescription != "" {
		desc := strings.Clone(person.ShortDescription)
		description = &desc
	}

	return &PersonDTO{
		Name:             person.Name,
		ShortDescription: description,
	}
}
