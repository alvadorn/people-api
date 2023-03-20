package application

import (
	"context"

	"github.com/alvadorn/people-api/components/people/domain"
	"github.com/alvadorn/people-api/components/people/repository"
)

func (s *useCasesSuite) TestGetByName() {
	s.Run(
		"successfully gets by name", func() {
			s.repository.On(
				"GetByName", context.Background(), "yoshua_bengio").
				Once().
				Return(&domain.Person{Name: "Yoshua Bengio"}, nil)

			value, err := s.uc.GetByName(context.Background(), "yoshua_bengio")
			s.Nil(err)
			s.Equal(value, &PersonDTO{Name: "Yoshua Bengio"})
		})

	s.Run(
		"not found person", func() {
			s.repository.On(
				"GetByName", context.Background(), "yoshua_bengio").
				Once().
				Return(nil, repository.NotFoundErr)

			value, err := s.uc.GetByName(context.Background(), "yoshua_bengio")
			s.Nil(value)
			s.Equal(err, NotFoundErr)
		})

	s.Run(
		"unexpected error", func() {
			s.repository.On(
				"GetByName", context.Background(), "yoshua_bengio").
				Once().
				Return(nil, repository.UnexpectedErr)

			value, err := s.uc.GetByName(context.Background(), "yoshua_bengio")
			s.Nil(value)
			s.Equal(err, UnexpectedError)
		})
}
