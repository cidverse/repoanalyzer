package analyzer

import (
	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/cidverse/repoanalyzer/modules/apispec"
	"github.com/cidverse/repoanalyzer/modules/cargo"
	"github.com/cidverse/repoanalyzer/modules/composer"
	"github.com/cidverse/repoanalyzer/modules/container"
	"github.com/cidverse/repoanalyzer/modules/dotnet"
	"github.com/cidverse/repoanalyzer/modules/gomod"
	"github.com/cidverse/repoanalyzer/modules/gradle"
	"github.com/cidverse/repoanalyzer/modules/helm"
	"github.com/cidverse/repoanalyzer/modules/helmfile"
	"github.com/cidverse/repoanalyzer/modules/hugo"
	"github.com/cidverse/repoanalyzer/modules/maven"
	"github.com/cidverse/repoanalyzer/modules/mkdocs"
	"github.com/cidverse/repoanalyzer/modules/nix"
	"github.com/cidverse/repoanalyzer/modules/node"
	"github.com/cidverse/repoanalyzer/modules/python"
)

// AllScanners contains all available scanners
var AllScanners = []analyzerapi.Scanner{
	apispec.Analyzer{},
	cargo.Analyzer{},
	composer.Analyzer{},
	container.Analyzer{},
	dotnet.Analyzer{},
	gomod.Analyzer{},
	gradle.Analyzer{},
	helm.Analyzer{},
	helmfile.Analyzer{},
	hugo.Analyzer{},
	maven.Analyzer{},
	mkdocs.Analyzer{},
	nix.Analyzer{},
	node.Analyzer{},
	python.Analyzer{},
}
