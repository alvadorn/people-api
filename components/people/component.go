package people

import (
	"github.com/alvadorn/people-api/components/config"
	"github.com/alvadorn/people-api/components/people/application"
	"github.com/alvadorn/people-api/components/people/interfaces/http"
	"github.com/alvadorn/people-api/components/people/repository"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type component struct {
	http.HTTP
}

type Component interface {
	RegisterRoutes(r *gin.RouterGroup)
}

type Options struct {
	RedisClient *redis.Client
	EnvVars     config.EnvironmentVariablesConfig
}

func New(opts *Options) *component {
	redisRepository := repository.NewRedisRepository(opts.RedisClient, opts.EnvVars.CacheTimeTTLInHours)
	apiClientRepository := repository.NewApiClient(opts.EnvVars.WikimediaBaseURL)

	repo := repository.New(
		repository.Config{
			RedisRepository:     redisRepository,
			APIClientRepository: apiClientRepository,
		})

	useCases := application.New(repo)
	httpInterface := http.New(useCases)
	return &component{
		HTTP: httpInterface,
	}
}
