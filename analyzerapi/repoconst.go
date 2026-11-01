package analyzerapi

// Well-known external artifact repositories. These are exported so consumers
// (such as cid) can reference the name, type and url without re-declaring them.

var (
	// RepositoryMavenCentral is the default Maven repository.
	RepositoryMavenCentral = ProjectRepository{Type: RepositoryTypeMaven, Name: "Maven Central", URL: "https://repo.maven.apache.org/maven2"}

	// RepositoryPrimeLibMaven is the PrimeLib Maven package registry.
	RepositoryPrimeLibMaven = ProjectRepository{Type: RepositoryTypeMaven, Name: "PrimeLib Maven", URL: "https://gitlab.com/api/v4/groups/primelib/-/packages/maven"}

	// RepositoryGradlePluginPortal is the Gradle plugin portal repository.
	RepositoryGradlePluginPortal = ProjectRepository{Type: RepositoryTypeGradle, Name: "gradlePluginPortal", URL: "https://plugins.gradle.org/m2/"}

	// RepositoryGoogleMaven is the Google Maven repository used by Gradle's google() shortcut.
	RepositoryGoogleMaven = ProjectRepository{Type: RepositoryTypeGradle, Name: "google", URL: "https://maven.google.com"}

	// RepositoryGoProxy is the default Go module proxy.
	RepositoryGoProxy = ProjectRepository{Type: RepositoryTypeGo, Name: "proxy.golang.org", URL: "https://proxy.golang.org"}

	// RepositoryNPMRegistry is the default npm registry.
	RepositoryNPMRegistry = ProjectRepository{Type: RepositoryTypeNpm, Name: "registry.npmjs.org", URL: "https://registry.npmjs.org"}

	// RepositoryCratesIO is the default crates.io registry.
	RepositoryCratesIO = ProjectRepository{Type: RepositoryTypeCargo, Name: "crates.io", URL: "https://crates.io"}

	// RepositoryPackagist is the default composer/packagist repository.
	RepositoryPackagist = ProjectRepository{Type: RepositoryTypeComposer, Name: "packagist.org", URL: "https://packagist.org"}

	// RepositoryPyPI is the default Python package index.
	RepositoryPyPI = ProjectRepository{Type: RepositoryTypePython, Name: "pypi.org", URL: "https://pypi.org"}

	// RepositoryNuGet is the default NuGet package source.
	RepositoryNuGet = ProjectRepository{Type: RepositoryTypeNuget, Name: "nuget.org", URL: "https://api.nuget.org/v3/index.json"}

	// RepositoryAnsibleGalaxy is the default Ansible Galaxy server.
	RepositoryAnsibleGalaxy = ProjectRepository{Type: RepositoryTypeGalaxy, Name: "galaxy.ansible.com", URL: "https://galaxy.ansible.com"}
)
