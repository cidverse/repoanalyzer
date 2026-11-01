package nix

import (
	"regexp"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

var nixURLPattern = regexp.MustCompile(`(?m)url\s*=\s*['"]([^'"]+)['"]`)

// normalizeNixReference converts shorthand Nix flake references (github:/gitlab:) to absolute URLs.
func normalizeNixReference(reference string) (string, bool) {
	reference = strings.TrimSpace(reference)
	lower := strings.ToLower(reference)

	if strings.HasPrefix(lower, "github:") {
		return "https://github.com/" + reference[len("github:"):], true
	}
	if strings.HasPrefix(lower, "gitlab:") {
		return "https://gitlab.com/" + reference[len("gitlab:"):], true
	}
	return reference, false
}

func nixRepositories(content string) []analyzerapi.ProjectRepository {
	var repositories []analyzerapi.ProjectRepository

	for _, match := range nixURLPattern.FindAllStringSubmatch(content, -1) {
		if len(match) < 2 {
			continue
		}

		reference, _ := normalizeNixReference(match[1])
		repository, ok := analyzerapi.ParseRepositoryReference(reference, analyzerapi.RepositoryTypeNix)
		if !ok {
			continue
		}
		repositories = append(repositories, repository)
	}

	return analyzerapi.MergeRepositories(repositories)
}
