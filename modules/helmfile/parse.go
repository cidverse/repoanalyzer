package helmfile

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parses Helmfile environments from a YAML string (used in testing and actual function)
func parseHelmfileEnvironmentsContent(data string) (map[string]interface{}, error) {
	// multiple yaml documents, allows parsing of helmfiles using go templates if the environments section is split off
	docs := strings.Split(data, "---")
	for _, doc := range docs {
		doc = strings.TrimSpace(doc)
		if doc == "" {
			continue
		}

		var helmfile map[string]interface{}
		err := yaml.Unmarshal([]byte(doc), &helmfile)
		if err != nil {
			continue
		}

		if envs, ok := helmfile["environments"].(map[string]interface{}); ok {
			return envs, nil
		}
	}

	return nil, errors.New("no environments section found in any YAML document")
}

// Reads a Helmfile from disk and extracts environments
func parseHelmfileEnvironments(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return parseHelmfileEnvironmentsContent(string(data))
}
