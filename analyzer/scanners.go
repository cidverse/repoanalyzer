package analyzer

import (
	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/cidverse/repoanalyzer/modules/ansible"
	"github.com/cidverse/repoanalyzer/modules/apispec"
	"github.com/cidverse/repoanalyzer/modules/cargo"
	"github.com/cidverse/repoanalyzer/modules/composer"
	"github.com/cidverse/repoanalyzer/modules/container"
	"github.com/cidverse/repoanalyzer/modules/deploymentdotenv"
	"github.com/cidverse/repoanalyzer/modules/dotnet"
	"github.com/cidverse/repoanalyzer/modules/githubworkflow"
	"github.com/cidverse/repoanalyzer/modules/gitlabci"
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
	"github.com/cidverse/repoanalyzer/modules/quartz"
	"github.com/cidverse/repoanalyzer/modules/renovate"
)

// AllScanners contains all available scanners
var AllScanners = []analyzerapi.Scanner{
	ansible.Analyzer{},
	apispec.Analyzer{},
	cargo.Analyzer{},
	composer.Analyzer{},
	container.Analyzer{},
	deploymentdotenv.Analyzer{
		AllowedEnvironmentNames: []string{"dev", "development", "test", "stage", "staging", "production"},
	},
	dotnet.Analyzer{},
	githubworkflow.Analyzer{},
	gitlabci.Analyzer{},
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
	quartz.Analyzer{},
	renovate.Analyzer{},
}
