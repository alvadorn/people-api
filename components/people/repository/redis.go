package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/alvadorn/people-api/components/people/domain"
	"github.com/redis/go-redis/v9"
)

type redisRepository struct {
	redis        *redis.Client
	cacheTimeTtl time.Duration
}

type redisPeopleOperator interface {
	personGetter
	personSaver
}

func NewRedisRepository(redis *redis.Client, cacheTimeTtl time.Duration) *redisRepository {
	return &redisRepository{
		redis:        redis,
		cacheTimeTtl: cacheTimeTtl,
	}
}

func (rr *redisRepository) getByName(ctx context.Context, name string) (*domain.Person, error) {
	value, err := rr.redis.Get(ctx, serializeNameToKey(name)).
		Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, NotFoundErr
		}
		return nil, fmt.Errorf("%v: %s", UnexpectedErr, err.Error())
	}

	var person domain.Person

	err = json.Unmarshal([]byte(value), &person)

	if err != nil {
		return nil, fmt.Errorf("%v: %s", JsonDecodingErr, err.Error())
	}

	return &person, nil
}

func (rr *redisRepository) savePerson(ctx context.Context, person *domain.Person) error {
	data, err := json.Marshal(person)

	if err != nil {
		return fmt.Errorf("%v: %s", JsonEncodingErr, err.Error())
	}

	_, err = rr.redis.Set(ctx, serializeNameToKey(person.Name), data, rr.cacheTimeTtl).
		Result()

	if err != nil {
		return fmt.Errorf("%v: %s", UnexpectedErr, err.Error())
	}

	return nil
}

func serializeNameToKey(name string) string {
	lowerCaseName := strings.ToLower(name)
	return strings.ReplaceAll(lowerCaseName, " ", "_")
}
