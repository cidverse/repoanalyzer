package helmfile

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func parseHelmfileEnvironments(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var helmfile map[string]interface{}
	err = yaml.Unmarshal(data, &helmfile)
	if err != nil {
		return nil, err
	}

	envs, ok := helmfile["environments"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("no environments section found in %s", filePath)
	}

	return envs, nil
}
