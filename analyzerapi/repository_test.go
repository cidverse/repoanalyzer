package analyzerapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRepositoryReference(t *testing.T) {
	tests := []struct {
		name        string
		reference   string
		defaultType ProjectRepositoryType
		expected    ProjectRepository
		valid       bool
	}{
		{
			name:        "helm repository",
			reference:   "https://charts.example.com/stable",
			defaultType: RepositoryTypeHelm,
			expected:    ProjectRepository{Type: RepositoryTypeHelm, URL: "https://charts.example.com/stable"},
			valid:       true,
		},
		{
			name:        "oci registry",
			reference:   "oci://registry.example.com/charts",
			defaultType: RepositoryTypeHelm,
			expected:    ProjectRepository{Type: RepositoryTypeOCI, URL: "oci://registry.example.com/charts"},
			valid:       true,
		},
		{
			name:        "galaxy server",
			reference:   "https://galaxy.example.com",
			defaultType: RepositoryTypeGalaxy,
			expected:    ProjectRepository{Type: RepositoryTypeGalaxy, URL: "https://galaxy.example.com"},
			valid:       true,
		},
		{
			name:        "scp-like git source",
			reference:   "git@github.com:org/repo.git",
			defaultType: RepositoryTypeGit,
			expected:    ProjectRepository{Type: RepositoryTypeGit, URL: "ssh://git@github.com/org/repo.git"},
			valid:       true,
		},
		{
			name:        "git+https source",
			reference:   "git+https://github.com/org/repo.git",
			defaultType: RepositoryTypeHelm,
			expected:    ProjectRepository{Type: RepositoryTypeGit, URL: "git+https://github.com/org/repo.git"},
			valid:       true,
		},
		{
			name:        "ssh source",
			reference:   "ssh://git@github.com/org/repo.git",
			defaultType: RepositoryTypeHelm,
			expected:    ProjectRepository{Type: RepositoryTypeGit, URL: "ssh://git@github.com/org/repo.git"},
			valid:       true,
		},
		{
			name:        "templated reference",
			reference:   "{{ .Values.repository }}",
			defaultType: RepositoryTypeHelm,
			valid:       false,
		},
		{
			name:        "helm alias",
			reference:   "@myrepo",
			defaultType: RepositoryTypeHelm,
			valid:       false,
		},
		{
			name:        "file reference",
			reference:   "file://../charts/local",
			defaultType: RepositoryTypeHelm,
			valid:       false,
		},
		{
			name:        "relative path",
			reference:   "../charts/local",
			defaultType: RepositoryTypeHelm,
			valid:       false,
		},
		{
			name:        "absolute path",
			reference:   "/charts/local",
			defaultType: RepositoryTypeHelm,
			valid:       false,
		},
		{
			name:        "windows path",
			reference:   `C:\charts\local`,
			defaultType: RepositoryTypeHelm,
			valid:       false,
		},
		{
			name:        "reference without scheme",
			reference:   "charts.example.com/stable",
			defaultType: RepositoryTypeHelm,
			valid:       false,
		},
		{
			name:        "empty reference",
			reference:   "",
			defaultType: RepositoryTypeHelm,
			valid:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repository, valid := ParseRepositoryReference(test.reference, test.defaultType)
			assert.Equal(t, test.valid, valid)
			if test.valid {
				assert.Equal(t, test.expected, repository)
			}
		})
	}
}

func TestMergeRepositories(t *testing.T) {
	first := []ProjectRepository{
		{Type: RepositoryTypeHelm, Name: "first", URL: "https://charts.example.com"},
		{Type: RepositoryTypeOCI, URL: "oci://registry.example.com/charts"},
	}
	second := []ProjectRepository{
		{Type: RepositoryTypeHelm, Name: "duplicate", URL: "https://charts.example.com"},
		{Type: RepositoryTypeGit, URL: "ssh://git@github.com/org/repo.git"},
	}

	result := MergeRepositories(first, second)

	assert.Equal(t, []ProjectRepository{
		{Type: RepositoryTypeHelm, Name: "first", URL: "https://charts.example.com"},
		{Type: RepositoryTypeOCI, URL: "oci://registry.example.com/charts"},
		{Type: RepositoryTypeGit, URL: "ssh://git@github.com/org/repo.git"},
	}, result)
}

func TestProjectRepository_Endpoint(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		host    string
		port    int
		dnsName string
	}{
		{name: "helm https", url: "https://charts.example.com/stable", host: "charts.example.com", port: 443, dnsName: "charts.example.com:443"},
		{name: "oci registry", url: "oci://registry.example.com/charts", host: "registry.example.com", port: 443, dnsName: "registry.example.com:443"},
		{name: "http", url: "http://example.com/repo", host: "example.com", port: 80, dnsName: "example.com:80"},
		{name: "ssh git", url: "ssh://git@github.com/org/repo.git", host: "github.com", port: 22, dnsName: "github.com:22"},
		{name: "explicit port", url: "https://registry.example.com:8443/charts", host: "registry.example.com", port: 8443, dnsName: "registry.example.com:8443"},
		{name: "go proxy", url: "https://proxy.golang.org", host: "proxy.golang.org", port: 443, dnsName: "proxy.golang.org:443"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repository := ProjectRepository{URL: test.url}
			assert.Equal(t, test.host, repository.Host())
			assert.Equal(t, test.port, repository.Port())
			assert.Equal(t, test.dnsName, repository.HostAndPort())
		})
	}
}

func TestWellKnownRepositories(t *testing.T) {
	tests := []struct {
		name       string
		repository ProjectRepository
		host       string
		port       int
	}{
		{name: "Maven Central", repository: RepositoryMavenCentral, host: "repo.maven.apache.org", port: 443},
		{name: "PrimeLib Maven", repository: RepositoryPrimeLibMaven, host: "gitlab.com", port: 443},
		{name: "NuGet", repository: RepositoryNuGet, host: "api.nuget.org", port: 443},
		{name: "Ansible Galaxy", repository: RepositoryAnsibleGalaxy, host: "galaxy.ansible.com", port: 443},
		{name: "Go proxy", repository: RepositoryGoProxy, host: "proxy.golang.org", port: 443},
		{name: "crates.io", repository: RepositoryCratesIO, host: "crates.io", port: 443},
		{name: "npm", repository: RepositoryNPMRegistry, host: "registry.npmjs.org", port: 443},
		{name: "packagist", repository: RepositoryPackagist, host: "packagist.org", port: 443},
		{name: "pypi", repository: RepositoryPyPI, host: "pypi.org", port: 443},
		{name: "gradle plugin portal", repository: RepositoryGradlePluginPortal, host: "plugins.gradle.org", port: 443},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.host, test.repository.Host())
			assert.Equal(t, test.port, test.repository.Port())
		})
	}
}
