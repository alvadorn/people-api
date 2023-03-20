package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alvadorn/people-api/components/people/domain"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const testRedisUrl = "localhost:6379"

type redisRepositorySuite struct {
	suite.Suite
	repository *redisRepository
	redis      *redis.Client
	ctx        context.Context
}

func TestNewRedisRepository(t *testing.T) {
	t.Run(
		"instantiate redis repository successfully", func(t *testing.T) {
			repo := NewRedisRepository(nil, 1*time.Hour)
			assert.NotNil(t, repo)
		})
}

func (s *redisRepositorySuite) SetupSuite() {
	s.ctx = context.Background()
	s.redis = redis.NewClient(
		&redis.Options{
			Addr: os.Getenv("REDIS_URL"),
			DB:   5,
		})
}

func (s *redisRepositorySuite) SetupTest() {
	s.repository = NewRedisRepository(s.redis, 2*time.Hour)
}

func (s *redisRepositorySuite) SetupSubTest() {
	s.SetupTest()
}

func (s *redisRepositorySuite) TearDownSuite() {
	s.redis.FlushDB(s.ctx)
	s.redis.Close()
	s.redis = nil
	s.repository = nil
}

func (s *redisRepositorySuite) TestGetByName() {
	s.Run(
		"retrieve by name successfully", func() {
			s.redis.Set(s.ctx, "yoshua_bengio", `{ "Name": "Yoshua Bengio", "ShortDescription": "xpto" }`, 2*time.Hour)
			output, err := s.repository.getByName(s.ctx, "yoshua_bengio")
			s.Nil(err)
			s.Equal(
				&domain.Person{
					Name:             "Yoshua Bengio",
					ShortDescription: "xpto",
				}, output)
		})

	s.Run(
		"fails on retrieval by name because of redis connection", func() {
			s.repository.redis = redis.NewClient(
				&redis.Options{
					Addr: "localhost:10000",
				})

			data, err := s.repository.getByName(s.ctx, "name")
			s.Nil(data)
			s.Error(err)
			s.ErrorContains(err, "unexpected_error")
		})

	s.Run(
		"fails on retrieval by name because value not found", func() {
			data, err := s.repository.getByName(s.ctx, "jose_valim")
			s.Nil(data)
			s.Error(err)
			s.ErrorContains(err, "not_found_error")
		})
}

func (s *redisRepositorySuite) TestSavePerson() {
	s.Run(
		"save person", func() {
			err := s.repository.savePerson(
				s.ctx, &domain.Person{
					Name:             "Dennis Ritchie",
					ShortDescription: "great computer scientist",
				})
			s.Nil(err)
			data, _ := s.redis.Get(s.ctx, "dennis_ritchie").Result()
			s.Equal(`{"Name":"Dennis Ritchie","ShortDescription":"great computer scientist"}`, data)
		})

	s.Run(
		"fails on save because there is no redis connection", func() {
			s.repository.redis = redis.NewClient(
				&redis.Options{
					Addr: "localhost:10000",
				})
			err := s.repository.savePerson(
				s.ctx, &domain.Person{
					Name:             "Dennis Ritchie",
					ShortDescription: "great computer scientist",
				})
			s.Error(err)
			s.ErrorContains(err, "unexpected_error")
		})
}

func TestRedisRepositorySuite(t *testing.T) {
	suite.Run(t, new(redisRepositorySuite))
}
