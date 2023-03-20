package application

import (
	"testing"

	"github.com/alvadorn/people-api/components/people/domain"
	"github.com/stretchr/testify/suite"
)

type useCasesSuite struct {
	suite.Suite

	repository *domain.PeopleRepositoryMock
	uc         *useCases
}

func TestUseCasesSuite(t *testing.T) {
	suite.Run(t, new(useCasesSuite))
}

func (s *useCasesSuite) SetupTest() {
	s.repository = domain.NewPeopleRepository(s.T())
	s.uc = New(s.repository)
}
