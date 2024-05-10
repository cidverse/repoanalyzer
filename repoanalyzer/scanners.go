package repoanalyzer

import (
	"github.com/cidverse/repoanalyzer/analyzer/apispec"
	"github.com/cidverse/repoanalyzer/analyzer/cargo"
	"github.com/cidverse/repoanalyzer/analyzer/composer"
	"github.com/cidverse/repoanalyzer/analyzer/container"
	"github.com/cidverse/repoanalyzer/analyzer/dotnet"
	"github.com/cidverse/repoanalyzer/analyzer/gomod"
	"github.com/cidverse/repoanalyzer/analyzer/gradle"
	"github.com/cidverse/repoanalyzer/analyzer/helm"
	"github.com/cidverse/repoanalyzer/analyzer/hugo"
	"github.com/cidverse/repoanalyzer/analyzer/maven"
	"github.com/cidverse/repoanalyzer/analyzer/mkdocs"
	"github.com/cidverse/repoanalyzer/analyzer/nix"
	"github.com/cidverse/repoanalyzer/analyzer/node"
	"github.com/cidverse/repoanalyzer/analyzer/python"
	"github.com/cidverse/repoanalyzer/analyzerapi"
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
	hugo.Analyzer{},
	maven.Analyzer{},
	mkdocs.Analyzer{},
	nix.Analyzer{},
	node.Analyzer{},
	python.Analyzer{},
}
