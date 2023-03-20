package repository

import (
	"context"
	"testing"

	"github.com/alvadorn/people-api/components/people/domain"
	"github.com/stretchr/testify/suite"
)

type repositorySuite struct {
	suite.Suite
	ctx             context.Context
	apiRepository   *apiClientPeopleOperatorMock
	redisRepository *redisPeopleOperatorMock
	repository      *repository
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(repositorySuite))
}

func (s *repositorySuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *repositorySuite) SetupTest() {
	s.apiRepository = newApiClientPeopleOperator(s.T())
	s.redisRepository = newRedisPeopleOperator(s.T())
	s.repository = New(
		Config{
			RedisRepository:     s.redisRepository,
			APIClientRepository: s.apiRepository,
		})
}

func (s *repositorySuite) TestGetByName() {
	s.Run(
		"get successfully on cache", func() {
			s.redisRepository.On("getByName", s.ctx, "yoshua bengio").
				Once().
				Return(
					&domain.Person{
						Name: "Yoshua Bengio",
					}, nil)
			value, err := s.repository.GetByName(s.ctx, "yoshua bengio")
			s.Nil(err)
			s.Equal(value, &domain.Person{Name: "Yoshua Bengio"})

		})

	s.Run(
		"get on api request", func() {
			s.redisRepository.On("getByName", s.ctx, "yoshua bengio").
				Once().
				Return(nil, NotFoundErr)
			s.apiRepository.On("getByName", s.ctx, "yoshua bengio").
				Once().
				Return(
					&domain.Person{
						Name: "Yoshua Bengio",
					}, nil)
			s.redisRepository.On("savePerson", s.ctx, &domain.Person{Name: "Yoshua Bengio"}).
				Once().
				Return(nil)
			value, err := s.repository.GetByName(s.ctx, "yoshua bengio")
			s.Nil(err)
			s.Equal(value, &domain.Person{Name: "Yoshua Bengio"})

		})

	s.Run(
		"does not find person", func() {
			s.redisRepository.On("getByName", s.ctx, "yoshua bengio").
				Once().
				Return(nil, NotFoundErr)
			s.apiRepository.On("getByName", s.ctx, "yoshua bengio").
				Once().
				Return(
					nil, NotFoundErr)

			value, err := s.repository.GetByName(s.ctx, "yoshua bengio")
			s.Nil(value)
			s.Error(err)
			s.ErrorIs(err, NotFoundErr)
		})
}
