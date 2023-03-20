package main

import (
	"fmt"

	"github.com/alvadorn/people-api/components/config"
	"github.com/alvadorn/people-api/components/people"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const defaultPort = "18080"

func main() {
	// Get environment variables
	cfg := config.NewEnvironmentVariablesConfig()

	// Create router
	router := gin.Default()
	v1Router := router.Group("/api/v1")

	// Create Redis Client
	redisClient := redis.NewClient(
		&redis.Options{
			Addr: cfg.RedisURL,
		})

	peopleComponent := people.New(
		&people.Options{
			RedisClient: redisClient,
			EnvVars:     cfg,
		})
	peopleComponent.RegisterRoutes(v1Router)

	router.Run(fmt.Sprint("0.0.0.0:", defaultPort))
}
