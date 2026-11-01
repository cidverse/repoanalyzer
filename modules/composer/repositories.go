package composer

import (
	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// composerRepositories extracts the external repositories declared by a composer.json.
// repositories may be a list [{type,url}] or a map {name:{type,url}}.
func composerRepositories(data ComposerJSON) []analyzerapi.ProjectRepository {
	var repositories []analyzerapi.ProjectRepository

	switch value := data.Repositories.(type) {
	case []interface{}:
		for _, item := range value {
			entry, ok := item.(map[string]interface{})
			if !ok {
				continue
			}
			if repository := repositoryFromComposerEntry(entry, ""); repository.URL != "" {
				repositories = append(repositories, repository)
			}
		}
	case map[string]interface{}:
		for name, item := range value {
			entry, ok := item.(map[string]interface{})
			if !ok {
				continue
			}
			if repository := repositoryFromComposerEntry(entry, name); repository.URL != "" {
				repositories = append(repositories, repository)
			}
		}
	}

	// a composer.json with no declared repository defaults to packagist.org
	if len(repositories) == 0 {
		repositories = append(repositories, analyzerapi.RepositoryPackagist)
	}

	return analyzerapi.MergeRepositories(repositories)
}

func repositoryFromComposerEntry(entry map[string]interface{}, name string) analyzerapi.ProjectRepository {
	rawURL, ok := entry["url"]
	if !ok {
		return analyzerapi.ProjectRepository{}
	}
	urlString, ok := rawURL.(string)
	if !ok {
		return analyzerapi.ProjectRepository{}
	}

	repository, ok := analyzerapi.ParseRepositoryReference(urlString, analyzerapi.RepositoryTypeComposer)
	if !ok {
		return analyzerapi.ProjectRepository{}
	}
	repository.Name = name
	return repository
}
