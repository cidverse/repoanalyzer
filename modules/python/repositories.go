package python

import (
	"os"
	"regexp"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// python index-url patterns across requirements.txt, Pipfile, pyproject.toml and setup.py
var pythonIndexURLPattern = regexp.MustCompile(`(?i)(?:--?index-url|--?extra-index-url|\bindex_url\b|\bindex-url\b)\s*[=: ]\s*['"]?([a-z][a-z0-9+.-]*://[^'"\s]+)`)
var pythonSourceURLPattern = regexp.MustCompile(`(?m)^\s*(?:(?:\[\[|\b)source(?:\]\]|\b))?\s*url\s*=\s*['"]([a-z][a-z0-9+.-]*://[^'"\s]+)['"]`)

// pythonRepositories extracts the package index(es) a file declares, best-effort.
func pythonRepositories(file string) []analyzerapi.ProjectRepository {
	content, err := os.ReadFile(file)
	if err != nil {
		return []analyzerapi.ProjectRepository{analyzerapi.RepositoryPyPI}
	}
	data := string(content)

	var repositories []analyzerapi.ProjectRepository

	for _, match := range pythonIndexURLPattern.FindAllStringSubmatch(data, -1) {
		if len(match) < 2 {
			continue
		}
		if repository, ok := analyzerapi.ParseRepositoryReference(match[1], analyzerapi.RepositoryTypePython); ok {
			repositories = append(repositories, repository)
		}
	}

	for _, match := range pythonSourceURLPattern.FindAllStringSubmatch(data, -1) {
		if len(match) < 2 {
			continue
		}
		if repository, ok := analyzerapi.ParseRepositoryReference(match[1], analyzerapi.RepositoryTypePython); ok {
			repositories = append(repositories, repository)
		}
	}

	if len(repositories) == 0 {
		repositories = append(repositories, analyzerapi.RepositoryPyPI)
	}

	return analyzerapi.MergeRepositories(repositories)
}
