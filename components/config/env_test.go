package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setEnvVars(vars map[string]string) func() {
	originalEnvVar := make(map[string]string, len(vars))
	for envKey, envValue := range vars {
		originalEnvVar[envKey] = os.Getenv(envKey)
		os.Setenv(envKey, envValue)
	}
	return func() {
		for envKey, envValue := range originalEnvVar {
			os.Setenv(envKey, envValue)
		}
	}
}

func TestNewEnvironmentVariablesConfig(t *testing.T) {
	t.Run(
		"Get all fields from env vars successfully", func(t *testing.T) {
			restoreEnvVars := setEnvVars(
				map[string]string{
					redisUrlVar:            "redis url",
					wikimediaBaseApiUrlVar: "api url",
					cacheTimeTtlInHoursVar: "12",
				})
			defer restoreEnvVars()

			envCfg := NewEnvironmentVariablesConfig()

			assert.Equal(
				t, envCfg, EnvironmentVariablesConfig{
					RedisURL:            "redis url",
					WikimediaBaseURL:    "api url",
					CacheTimeTTLInHours: 12 * time.Hour,
				})
		})

	t.Run(
		"Get all fields from env vars with default value successfully", func(t *testing.T) {
			restoreEnvVars := setEnvVars(
				map[string]string{
					redisUrlVar:            "redis url",
					wikimediaBaseApiUrlVar: "api url",
				})
			defer restoreEnvVars()

			envCfg := NewEnvironmentVariablesConfig()

			assert.Equal(
				t, envCfg, EnvironmentVariablesConfig{
					RedisURL:            "redis url",
					WikimediaBaseURL:    "api url",
					CacheTimeTTLInHours: 1 * time.Hour,
				})
		})
}
