package gradle

import (
	"regexp"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

var gradleURLPattern = regexp.MustCompile(`(?m)url\s*\(?\s*['"]([^'"]+)['"]\s*\)?`)

// gradleShortcutRepositories maps common Gradle repository shortcuts to the host actually contacted.
// The shortcut is what appears in the build script (e.g. mavenCentral(), google(), gradlePluginPortal()).
var gradleShortcutRepositories = map[string]analyzerapi.ProjectRepository{
	"mavenCentral":       analyzerapi.RepositoryMavenCentral,
	"google":             analyzerapi.RepositoryGoogleMaven,
	"gradlePluginPortal": analyzerapi.RepositoryGradlePluginPortal,
}

// gradleRepositories extracts external repositories from a build.gradle / build.gradle.kts file, best-effort.
func gradleRepositories(content string) []analyzerapi.ProjectRepository {
	var repositories []analyzerapi.ProjectRepository

	for _, match := range gradleURLPattern.FindAllStringSubmatch(content, -1) {
		if len(match) < 2 {
			continue
		}
		repository, ok := analyzerapi.ParseRepositoryReference(match[1], analyzerapi.RepositoryTypeGradle)
		if !ok {
			continue
		}
		repositories = append(repositories, repository)
	}

	for shortcut, repository := range gradleShortcutRepositories {
		if containsWord(content, shortcut+"()") {
			repositories = append(repositories, repository)
		}
	}

	return analyzerapi.MergeRepositories(repositories)
}

func containsWord(content string, word string) bool {
	if !strings.Contains(content, word) {
		return false
	}
	index := strings.Index(content, word)
	// ensure the word is a whole token, e.g. not part of a longer identifier
	if index > 0 && isIdentifierChar(content[index-1]) {
		return false
	}
	end := index + len(word)
	if end < len(content) && isIdentifierChar(content[end]) {
		return false
	}
	return true
}

func isIdentifierChar(char byte) bool {
	return char == '_' || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}
