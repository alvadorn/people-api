package people

import (
	"testing"

	"github.com/alvadorn/people-api/components/config"
	"github.com/stretchr/testify/assert"
)

func TestNew_Component(t *testing.T) {
	t.Run(
		"instantiates successfully", func(t *testing.T) {
			cmp := New(
				&Options{
					RedisClient: nil,
					EnvVars:     config.EnvironmentVariablesConfig{},
				})

			assert.NotNil(t, cmp)
		})
}
