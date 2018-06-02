package responsetransformer

import (
	"testing"

	"github.com/hellofresh/janus/pkg/api"
	"github.com/hellofresh/janus/pkg/plugin"
	"github.com/stretchr/testify/assert"
)

func TestResponseTransformerConfig(t *testing.T) {
	var config Config
	rawConfig := map[string]interface{}{
		"add": map[string]interface{}{
			"headers": map[string]string{
				"NAME": "TEST",
			},
			"querystring": map[string]string{
				"name": "test",
			},
		},
	}

	err := plugin.Decode(rawConfig, &config)
	assert.NoError(t, err)

	assert.IsType(t, map[string]string{}, config.Add.Headers)
	assert.Contains(t, config.Add.Headers, "NAME")
}

func TestResponseTransformerPlugin(t *testing.T) {
	rawConfig := map[string]interface{}{
		"add": map[string]interface{}{
			"headers": map[string]string{
				"NAME": "TEST",
			},
		},
	}

	def := api.NewDefinition()
	err := setupResponseTransformer(def, rawConfig)
	assert.NoError(t, err)

	assert.Len(t, def.Proxy.Middleware(), 1)
}
