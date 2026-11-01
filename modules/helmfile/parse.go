package helmfile

import (
	"errors"
	"os"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"gopkg.in/yaml.v3"
)

type helmfileRepository struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type helmfileDocument struct {
	Repositories []helmfileRepository `yaml:"repositories"`
}

// Parses Helmfile repositories from a YAML string (used in testing and actual function)
func parseHelmfileRepositoriesContent(data string) []analyzerapi.ProjectRepository {
	var result []analyzerapi.ProjectRepository

	// multiple yaml documents, allows parsing of helmfiles using go templates that split the repositories section off
	docs := strings.Split(data, "---")
	for _, doc := range docs {
		doc = strings.TrimSpace(doc)
		if doc == "" {
			continue
		}

		var helmfile helmfileDocument
		err := yaml.Unmarshal([]byte(doc), &helmfile)
		if err == nil {
			result = append(result, repositoriesFromHelmfileDocument(helmfile)...)
			continue
		}

		// fall back to a line scan, e.g. for gotmpl documents that are not valid YAML
		result = append(result, scanHelmfileRepositories(doc)...)
	}

	return analyzerapi.MergeRepositories(result)
}

// Reads a Helmfile from disk and extracts repositories
func parseHelmfileRepositories(filePath string) []analyzerapi.ProjectRepository {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	return parseHelmfileRepositoriesContent(string(data))
}

func repositoriesFromHelmfileDocument(helmfile helmfileDocument) []analyzerapi.ProjectRepository {
	var result []analyzerapi.ProjectRepository

	for _, item := range helmfile.Repositories {
		repository, ok := analyzerapi.ParseRepositoryReference(item.URL, analyzerapi.RepositoryTypeHelm)
		if !ok {
			continue
		}
		repository.Name = helmfileRepositoryName(item.Name)
		result = append(result, repository)
	}

	return result
}

// scanHelmfileRepositories extracts name/url pairs from a top-level repositories block without requiring valid YAML.
func scanHelmfileRepositories(data string) []analyzerapi.ProjectRepository {
	var result []analyzerapi.ProjectRepository
	inBlock := false
	name := ""

	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimRight(line, "\r")
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		indent := len(line) - len(strings.TrimLeft(line, " \t"))
		if !inBlock {
			if indent == 0 && isRepositoriesKey(trimmed) {
				inBlock = true
			}
			continue
		}

		// a new top-level key ends the repositories block
		if indent == 0 && !strings.HasPrefix(trimmed, "-") && isTopLevelKey(trimmed) {
			break
		}

		key, value, ok := parseHelmfileRepositoryLine(trimmed)
		if !ok {
			continue
		}

		switch key {
		case "name":
			name = value
		case "url":
			repository, valid := analyzerapi.ParseRepositoryReference(value, analyzerapi.RepositoryTypeHelm)
			if valid {
				repository.Name = helmfileRepositoryName(name)
				result = append(result, repository)
			}
			name = ""
		}
	}

	return result
}

func isRepositoriesKey(line string) bool {
	if index := strings.Index(line, "#"); index >= 0 {
		line = line[:index]
	}
	return strings.TrimSpace(line) == "repositories:"
}

func isTopLevelKey(line string) bool {
	index := strings.Index(line, ":")
	if index <= 0 {
		return false
	}

	for _, char := range line[:index] {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') && char != '_' && char != '-' && char != '.' {
			return false
		}
	}

	return true
}

func parseHelmfileRepositoryLine(line string) (string, string, bool) {
	line = strings.TrimSpace(strings.TrimPrefix(line, "-"))
	index := strings.Index(line, ":")
	if index <= 0 {
		return "", "", false
	}

	key := strings.TrimSpace(line[:index])
	if key != "name" && key != "url" {
		return "", "", false
	}

	value := trimYAMLValue(line[index+1:])
	if value == "" {
		return "", "", false
	}

	return key, value, true
}

func trimYAMLValue(value string) string {
	value = strings.TrimSpace(value)
	if len(value) > 0 && (value[0] == '"' || value[0] == '\'') {
		quote := value[0]
		if end := strings.IndexByte(value[1:], quote); end >= 0 {
			return value[1 : end+1]
		}
		return strings.Trim(value, string(quote))
	}

	if index := strings.Index(value, " #"); index >= 0 {
		value = value[:index]
	}

	return strings.TrimSpace(value)
}

func helmfileRepositoryName(name string) string {
	name = strings.TrimSpace(name)
	if strings.Contains(name, "{{") {
		return ""
	}
	return name
}

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
