package analyzerapi

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// defaultRepositoryPort maps a URL scheme to the port the host is contacted on,
// when the URL does not declare an explicit port.
func defaultRepositoryPort(scheme string) int {
	switch strings.ToLower(scheme) {
	case "https", "oci", "git+https":
		return 443
	case "http", "git+http":
		return 80
	case "ssh", "git+ssh":
		return 22
	case "git":
		return 9418
	default:
		return 443
	}
}

var (
	// scp-like git sources, e.g. git@github.com:org/repo.git
	scpLikeRepositoryPattern = regexp.MustCompile(`^([A-Za-z0-9._-]+@[A-Za-z0-9._-]+):(.+)$`)
	windowsPathPattern       = regexp.MustCompile(`^[A-Za-z]:[\\/]`)
)

// ParseRepositoryReference converts a URL / scp-like git source / OCI reference into a ProjectRepository.
// Returns false for local, templated, aliased or non-URL references.
func ParseRepositoryReference(reference string, defaultType ProjectRepositoryType) (ProjectRepository, bool) {
	reference = strings.TrimSpace(reference)
	if reference == "" || strings.Contains(reference, "{{") || strings.Contains(reference, "}}") {
		return ProjectRepository{}, false
	}
	if strings.HasPrefix(reference, "file://") || isLocalRepositoryReference(reference) {
		return ProjectRepository{}, false
	}

	// scp-like git source (git@host:org/repo.git) -> ssh://git@host/org/repo.git
	if match := scpLikeRepositoryPattern.FindStringSubmatch(reference); match != nil {
		return ProjectRepository{
			Type: RepositoryTypeGit,
			URL:  "ssh://" + match[1] + "/" + strings.TrimPrefix(match[2], "/"),
		}, true
	}

	parsed, err := url.Parse(reference)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return ProjectRepository{}, false
	}

	var repositoryType ProjectRepositoryType
	switch strings.ToLower(parsed.Scheme) {
	case "oci":
		repositoryType = RepositoryTypeOCI
	case "http", "https":
		repositoryType = defaultType
	case "git", "ssh", "git+ssh", "git+http", "git+https":
		repositoryType = RepositoryTypeGit
	default:
		return ProjectRepository{}, false
	}

	if repositoryType == "" {
		return ProjectRepository{}, false
	}

	return ProjectRepository{
		Type: repositoryType,
		URL:  reference,
	}, true
}

// isLocalRepositoryReference reports whether the reference points to a local file or directory.
func isLocalRepositoryReference(reference string) bool {
	if reference == "." || reference == ".." ||
		strings.HasPrefix(reference, "/") ||
		strings.HasPrefix(reference, "\\") ||
		strings.HasPrefix(reference, "./") ||
		strings.HasPrefix(reference, "../") ||
		strings.HasPrefix(reference, ".\\") ||
		strings.HasPrefix(reference, "..\\") ||
		strings.HasPrefix(reference, "~/") {
		return true
	}

	return windowsPathPattern.MatchString(reference)
}

// MergeRepositories concatenates slices, deduplicating by (type, url) and preserving order.
func MergeRepositories(groups ...[]ProjectRepository) []ProjectRepository {
	var result []ProjectRepository
	seen := make(map[string]struct{})

	for _, group := range groups {
		for _, repository := range group {
			if repository.URL == "" {
				continue
			}

			key := string(repository.Type) + "\x00" + repository.URL
			if _, exists := seen[key]; exists {
				continue
			}
			seen[key] = struct{}{}
			result = append(result, repository)
		}
	}

	return result
}

// Host returns the hostname of the repository endpoint, or an empty string if the URL cannot be resolved.
func (r ProjectRepository) Host() string {
	parsed, err := url.Parse(r.URL)
	if err != nil {
		return ""
	}
	return parsed.Hostname()
}

// Port returns the port the repository endpoint is served on. An explicit port in the URL wins,
// otherwise the scheme default is used.
func (r ProjectRepository) Port() int {
	parsed, err := url.Parse(r.URL)
	if err != nil {
		return 0
	}
	if parsed.Port() != "" {
		if port, err := strconv.Atoi(parsed.Port()); err == nil {
			return port
		}
	}
	return defaultRepositoryPort(parsed.Scheme)
}

// PortString returns the port as a string, suitable for building a "host:port" endpoint.
func (r ProjectRepository) PortString() string {
	return strconv.Itoa(r.Port())
}

// HostAndPort returns the "host:port" egress allow-list endpoint, or an empty string if unresolvable.
func (r ProjectRepository) HostAndPort() string {
	host := r.Host()
	if host == "" {
		return ""
	}
	port := r.Port()
	if port == 0 {
		return host
	}
	return host + ":" + strconv.Itoa(port)
}
