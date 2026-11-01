package cargo

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

var cargoURLPattern = regexp.MustCompile(`(?m)^\s*(?:index|git|registry)\s*=\s*['"]([^'"]+)['"]`)

// cargoRepositories extracts external sources from a Cargo.toml and any sibling .cargo/config.toml.
func cargoRepositories(moduleDir string, projectDir string) []analyzerapi.ProjectRepository {
	var sources []string

	for _, directory := range []string{moduleDir, projectDir} {
		if directory == "" {
			continue
		}
		cargoToml := filepath.Join(directory, "Cargo.toml")
		if content, err := os.ReadFile(cargoToml); err == nil {
			sources = append(sources, string(content))
		}
		config := filepath.Join(directory, ".cargo", "config.toml")
		if content, err := os.ReadFile(config); err == nil {
			sources = append(sources, string(content))
		}
	}

	var repositories []analyzerapi.ProjectRepository
	hasExplicit := false

	for _, source := range sources {
		for _, match := range cargoURLPattern.FindAllStringSubmatch(source, -1) {
			if len(match) < 2 {
				continue
			}
			hasExplicit = true
			line := strings.TrimSpace(match[0])
			if strings.HasPrefix(line, "index =") || strings.HasPrefix(line, "registry =") {
				repository, ok := analyzerapi.ParseRepositoryReference(match[1], analyzerapi.RepositoryTypeCargo)
				if !ok {
					continue
				}
				repositories = append(repositories, repository)
				continue
			}
			// git = "..." is a direct git dependency source
			repository, ok := analyzerapi.ParseRepositoryReference(match[1], analyzerapi.RepositoryTypeGit)
			if !ok {
				continue
			}
			repositories = append(repositories, repository)
		}
	}

	if !hasExplicit || len(repositories) == 0 {
		repositories = append(repositories, analyzerapi.RepositoryCratesIO)
	}

	return analyzerapi.MergeRepositories(repositories)
}
