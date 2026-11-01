package dotnet

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

type nugetConfigDocument struct {
	XMLName        xml.Name `xml:"configuration"`
	PackageSources struct {
		Add []struct {
			Key   string `xml:"key,attr"`
			Value string `xml:"value,attr"`
		} `xml:"add"`
	} `xml:"packageSources"`
}

// nugetRepositories resolves the NuGet package sources for a module, looking for a NuGet.config
// in the module directory and the project directory. Falls back to the default nuget.org source.
func nugetRepositories(moduleDir string, projectDir string) []analyzerapi.ProjectRepository {
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

		if config := findNuGetConfig(directory); config != "" {
			document, err := readNuGetConfig(config)
			if err != nil {
				continue
			}
			if len(document.PackageSources.Add) > 0 {
				for _, source := range document.PackageSources.Add {
					repository, ok := analyzerapi.ParseRepositoryReference(source.Value, analyzerapi.RepositoryTypeNuget)
					if !ok {
						continue
					}
					repository.Name = source.Key
					repositories = append(repositories, repository)
				}
			}
		}
	}

	if len(repositories) == 0 {
		repositories = append(repositories, analyzerapi.RepositoryNuGet)
	}

	return analyzerapi.MergeRepositories(repositories)
}

func findNuGetConfig(directory string) string {
	for _, name := range []string{"NuGet.config", "nuget.config", "NuGet.Config", "nuget.Config"} {
		path := filepath.Join(directory, name)
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path
		}
	}
	return ""
}

func readNuGetConfig(path string) (nugetConfigDocument, error) {
	var document nugetConfigDocument
	content, err := os.ReadFile(path)
	if err != nil {
		return document, err
	}
	if err := xml.Unmarshal(content, &document); err != nil {
		return document, err
	}
	return document, nil
}
