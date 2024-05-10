package composer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseComposerJSONSuccess(t *testing.T) {
	content := []byte(`{
			"name": "my-org/my-project",
			"description": "Awesome Project",
			"homepage": "https://github.com",
			"type": "library",
			"license": "MIT",
			"require": {
				"php": ">=7.0.0",
				"ext-mbstring": "*"
			}
		}`)
	data, err := ParseComposerJSONFromByteArray(content)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "my-org/my-project", data.Name)
	assert.Equal(t, "Awesome Project", data.Description)
	assert.Equal(t, "https://github.com", data.Homepage)
	assert.Equal(t, "library", data.Type)
	assert.Equal(t, "MIT", data.License)
	assert.Equal(t, map[string]string{
		"php":          ">=7.0.0",
		"ext-mbstring": "*",
	}, data.Require)
}
