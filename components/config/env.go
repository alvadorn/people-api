package config

import (
	"os"
	"strconv"
	"time"
)

const (
	redisUrlVar            = "REDIS_URL"
	wikimediaBaseApiUrlVar = "WIKIMEDIA_BASE_API_URL"
	cacheTimeTtlInHoursVar = "CACHE_TIME_TTL_IN_HOURS"
)

type EnvironmentVariablesConfig struct {
	RedisURL            string
	WikimediaBaseURL    string
	CacheTimeTTLInHours time.Duration
}

func NewEnvironmentVariablesConfig() EnvironmentVariablesConfig {
	return EnvironmentVariablesConfig{
		RedisURL:            os.Getenv(redisUrlVar),
		WikimediaBaseURL:    os.Getenv(wikimediaBaseApiUrlVar),
		CacheTimeTTLInHours: getTimeEnvVar(cacheTimeTtlInHoursVar, 1) * time.Hour,
	}
}

func getTimeEnvVar(variable string, defaultValue time.Duration) time.Duration {
	data := os.Getenv(variable)
	timeInt, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return defaultValue
	}

	return time.Duration(timeInt)
}
