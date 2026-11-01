package ansible

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"gopkg.in/yaml.v3"
)

var (
	// requirementPaths contains the supported requirements files, relative to a module or project directory
	requirementPaths = []string{
		"roles/requirements.yml",
		"roles/requirements.yaml",
		"collections/requirements.yml",
		"collections/requirements.yaml",
		"requirements.yml",
		"requirements.yaml",
	}

	bareGalaxyNamePattern = regexp.MustCompile(`^[A-Za-z0-9_][A-Za-z0-9_.-]*$`)
)

type requirementsEntry struct {
	Name   string `yaml:"name"`
	Src    string `yaml:"src"`
	Source string `yaml:"source"`
	Type   string `yaml:"type"`
}

type requirementsFile struct {
	Roles       []requirementsEntry `yaml:"roles"`
	Collections []requirementsEntry `yaml:"collections"`
}

// ansibleRepositories collects external repositories from requirements files in the module and project directory.
// Parsing is best-effort: missing or invalid files are ignored.
func ansibleRepositories(moduleDir string, projectDir string) []analyzerapi.ProjectRepository {
	var repositories []analyzerapi.ProjectRepository
	visited := make(map[string]struct{})

	for _, directory := range []string{moduleDir, projectDir} {
		if directory == "" {
			continue
		}

		directory = filepath.Clean(directory)
		if _, exists := visited[directory]; exists {
			continue
		}
		visited[directory] = struct{}{}

		for _, relativePath := range requirementPaths {
			file := filepath.Join(directory, relativePath)
			info, err := os.Stat(file)
			if err != nil || info.IsDir() {
				continue
			}

			repositories = append(repositories, parseRequirementsFile(file, directory)...)
		}
	}

	return analyzerapi.MergeRepositories(repositories)
}

func parseRequirementsFile(file string, moduleDir string) []analyzerapi.ProjectRepository {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	var requirements requirementsFile
	if err := yaml.Unmarshal(data, &requirements); err != nil {
		return nil
	}

	var repositories []analyzerapi.ProjectRepository

	for _, role := range requirements.Roles {
		reference := strings.TrimSpace(role.Src)
		if reference == "" {
			reference = strings.TrimSpace(role.Name)
		}
		if reference == "" {
			continue
		}

		if repository, ok := analyzerapi.ParseRepositoryReference(reference, analyzerapi.RepositoryTypeGit); ok {
			repository.Name = requirementName(role.Name)
			repositories = append(repositories, repository)
			continue
		}

		if isLocalRequirement(reference, moduleDir) || !isBareGalaxyName(reference) {
			continue
		}

		repositories = append(repositories, defaultGalaxyRepository())
	}

	for _, collection := range requirements.Collections {
		source := strings.TrimSpace(collection.Source)
		if source != "" {
			if repository, ok := analyzerapi.ParseRepositoryReference(source, analyzerapi.RepositoryTypeGalaxy); ok {
				repository.Name = source
				repositories = append(repositories, repository)
				continue
			}
		}

		reference := strings.TrimSpace(collection.Name)
		if reference == "" {
			reference = source
		}
		if reference == "" {
			continue
		}

		if repository, ok := analyzerapi.ParseRepositoryReference(reference, analyzerapi.RepositoryTypeGit); ok {
			repository.Name = requirementName(collection.Name)
			repositories = append(repositories, repository)
			continue
		}

		if isLocalRequirement(reference, moduleDir) || !isBareGalaxyName(reference) {
			continue
		}

		repositories = append(repositories, defaultGalaxyRepository())
	}

	return repositories
}

func defaultGalaxyRepository() analyzerapi.ProjectRepository {
	return analyzerapi.RepositoryAnsibleGalaxy
}

func requirementName(name string) string {
	name = strings.TrimSpace(name)
	if strings.Contains(name, "{{") {
		return ""
	}
	return name
}

func isBareGalaxyName(reference string) bool {
	if strings.Contains(reference, "{{") || strings.Contains(reference, "}}") {
		return false
	}
	return bareGalaxyNamePattern.MatchString(reference)
}

// isLocalRequirement reports whether a requirements entry points to a local role/collection directory.
func isLocalRequirement(reference string, moduleDir string) bool {
	if strings.HasPrefix(reference, ".") ||
		strings.HasPrefix(reference, "/") ||
		strings.HasPrefix(reference, "\\") ||
		strings.HasPrefix(reference, "~") ||
		strings.HasPrefix(reference, "file://") ||
		strings.ContainsAny(reference, `/\`) {
		return true
	}

	for _, base := range []string{"roles", "collections"} {
		info, err := os.Stat(filepath.Join(moduleDir, base, reference))
		if err == nil && info.IsDir() {
			return true
		}
	}

	return false
}
