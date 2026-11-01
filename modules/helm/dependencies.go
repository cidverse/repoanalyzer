package helm

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"gopkg.in/yaml.v3"
)

type chartDependency struct {
	Name       string `yaml:"name"`
	Repository string `yaml:"repository"`
}

type chartDependencies struct {
	Dependencies []chartDependency `yaml:"dependencies"`
}

// parseChartRepositories collects the external repositories declared by a Chart.yaml and its sibling Chart.lock.
// Parsing is best-effort: unreadable or invalid files are ignored.
func parseChartRepositories(chartFile string) []analyzerapi.ProjectRepository {
	var chart chartDependencies
	if err := readYAMLFile(chartFile, &chart); err != nil {
		return nil
	}

	lock := chartDependencies{}
	lockFile := filepath.Join(filepath.Dir(chartFile), "Chart.lock")
	if err := readYAMLFile(lockFile, &lock); err != nil {
		lock.Dependencies = nil
	}

	resolved := make(map[string]string, len(lock.Dependencies))
	for _, dependency := range lock.Dependencies {
		if dependency.Name == "" || dependency.Repository == "" {
			continue
		}
		resolved[dependency.Name] = dependency.Repository
	}

	var repositories []analyzerapi.ProjectRepository

	// dependencies declared in Chart.yaml, resolving aliases through Chart.lock
	for _, dependency := range chart.Dependencies {
		reference := strings.TrimSpace(dependency.Repository)
		if reference == "" || strings.HasPrefix(reference, "@") {
			reference = resolved[dependency.Name]
		}
		repositories = append(repositories, chartRepository(reference, dependency.Name)...)
	}

	// dependencies only present in Chart.lock (e.g. a lock without a matching Chart.yaml entry)
	for _, dependency := range lock.Dependencies {
		repositories = append(repositories, chartRepository(dependency.Repository, dependency.Name)...)
	}

	return analyzerapi.MergeRepositories(repositories)
}

func chartRepository(reference string, name string) []analyzerapi.ProjectRepository {
	repository, ok := analyzerapi.ParseRepositoryReference(reference, analyzerapi.RepositoryTypeHelm)
	if !ok {
		return nil
	}

	repository.Name = name
	return []analyzerapi.ProjectRepository{repository}
}

func readYAMLFile(file string, target interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, target)
}
