package node

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// npmRepositories resolves the npm registry used by a package, looking for a sibling .npmrc
// in the module and project directory.
func npmRepositories(moduleDir string, projectDir string) []analyzerapi.ProjectRepository {
	var repositories []analyzerapi.ProjectRepository

	seen := make(map[string]struct{})
	for _, directory := range []string{moduleDir, projectDir} {
		if directory == "" {
			continue
		}
		directory = filepath.Clean(directory)
		if _, exists := seen[directory]; exists {
			continue
		}
		seen[directory] = struct{}{}

		npmrc := filepath.Join(directory, ".npmrc")
		file, err := os.Open(npmrc)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			if !strings.Contains(line, "=") {
				continue
			}
			key, value, _ := strings.Cut(line, "=")
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)
			if strings.EqualFold(key, "registry") {
				if repository, ok := analyzerapi.ParseRepositoryReference(value, analyzerapi.RepositoryTypeNpm); ok {
					repositories = append(repositories, repository)
				}
				continue
			}
			// scoped registries: @scope:registry=https://...
			if strings.HasSuffix(key, ":registry") && strings.HasPrefix(key, "@") {
				if repository, ok := analyzerapi.ParseRepositoryReference(value, analyzerapi.RepositoryTypeNpm); ok {
					name := strings.TrimSuffix(strings.TrimPrefix(key, "@"), ":registry")
					repository.Name = name
					repositories = append(repositories, repository)
				}
			}
		}
		file.Close()
	}

	if len(repositories) == 0 {
		repositories = append(repositories, analyzerapi.RepositoryNPMRegistry)
	}

	return analyzerapi.MergeRepositories(repositories)
}
