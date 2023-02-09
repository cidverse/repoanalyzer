package composer

import (
	"encoding/json"
	"fmt"
	"os"
)

type ComposerJSON struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Homepage    string   `json:"homepage"`
	Keywords    []string `json:"keywords"`
	Type        string   `json:"type"`
	License     string   `json:"license"`
	Authors     []struct {
		Name  string `json:"name"`
		Role  string `json:"role"`
		Email string `json:"email,omitempty"`
	} `json:"authors"`
	Funding []struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"funding"`
	Require    map[string]string   `json:"require"`
	RequireDev map[string]string   `json:"require-dev"`
	Autoload   interface{}         `json:"autoload"`
	Scripts    map[string][]string `json:"scripts"`
}

// ParseComposerJSON parses the composer.json
func ParseComposerJSON(file string) (ComposerJSON, error) {
	// Read the contents of the build.gradle file
	content, err := os.ReadFile(file)
	if err != nil {
		return ComposerJSON{}, fmt.Errorf("failed to open composer.json: %v", err)
	}

	return ParseComposerJSONFromByteArray(content)
}

// ParseComposerJSONFromByteArray parses the composer.json
func ParseComposerJSONFromByteArray(content []byte) (ComposerJSON, error) {
	var data ComposerJSON

	err := json.Unmarshal(content, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
