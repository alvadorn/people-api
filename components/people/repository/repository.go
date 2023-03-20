package repository

import (
	"context"

	"github.com/alvadorn/people-api/components/people/domain"
)

type Config struct {
	RedisRepository     redisPeopleOperator
	APIClientRepository apiClientPeopleOperator
}

type repository struct {
	redisRepository redisPeopleOperator
	apiClient       apiClientPeopleOperator
}

type personGetter interface {
	getByName(ctx context.Context, name string) (*domain.Person, error)
}

type personSaver interface {
	savePerson(ctx context.Context, person *domain.Person) error
}

func New(cfg Config) *repository {
	return &repository{
		redisRepository: cfg.RedisRepository,
		apiClient:       cfg.APIClientRepository,
	}
}

func (r *repository) GetByName(ctx context.Context, name string) (*domain.Person, error) {
	person, err := r.redisRepository.getByName(ctx, name)

	if err == nil {
		return person, err
	}

	person, err = r.apiClient.getByName(ctx, name)

	if err != nil {
		return nil, err
	}

	r.redisRepository.savePerson(ctx, person)
	return person, nil
}
