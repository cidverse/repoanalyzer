package maven

import (
	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// pomXMLRepositories extracts the external repositories declared by a pom.xml.
func pomXMLRepositories(pom PomXML) []analyzerapi.ProjectRepository {
	var repositories []analyzerapi.ProjectRepository

	seen := false
	for _, repository := range pom.Repositories {
		seen = true
		parsed, ok := analyzerapi.ParseRepositoryReference(repository.URL, analyzerapi.RepositoryTypeMaven)
		if !ok {
			continue
		}
		parsed.Name = repository.ID
		repositories = append(repositories, parsed)
	}

	// a pom without any declared repository defaults to Maven Central
	if !seen {
		repositories = append(repositories, analyzerapi.RepositoryMavenCentral)
	}

	return analyzerapi.MergeRepositories(repositories)
}
