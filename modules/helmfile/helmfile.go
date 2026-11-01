package helmfile

import (
	"path/filepath"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "helmfile"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.FilesByExtension["yaml"] {
		filename := filepath.Base(file)
		repositories := parseHelmfileRepositories(file)

		if filename == "helmfile.yaml" {
			envs, err := parseHelmfileEnvironments(file)
			if err == nil {
				for k, _ := range envs {
					if k == "default" {
						continue
					}
					module := analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), "deployment-helmfile-"+k, analyzerapi.DeploymentSpecHelmfile, "helmfile", k)
					module.Repositories = repositories
					analyzerapi.AddModuleToResult(&result, module)
				}
			} else {
				module := analyzerapi.CreateProjectBuildSystemModule(ctx, file, a.GetName(), analyzerapi.BuildSystemHelmfile)
				module.Repositories = repositories
				analyzerapi.AddModuleToResult(&result, module)
			}
		} else if strings.HasPrefix(filename, "helmfile-") && strings.HasSuffix(filename, ".yaml") {
			envName := strings.TrimSuffix(strings.TrimPrefix(filename, "helmfile-"), ".yaml")
			module := analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), "deployment-helmfile-"+envName, analyzerapi.DeploymentSpecHelmfile, "helmfile", envName)
			module.Repositories = repositories
			analyzerapi.AddModuleToResult(&result, module)
		}
	}

	for _, file := range ctx.FilesByExtension["gotmpl"] {
		filename := filepath.Base(file)
		repositories := parseHelmfileRepositories(file)

		if filename == "helmfile.yaml.gotmpl" {
			envs, err := parseHelmfileEnvironments(file)
			if err == nil {
				for k, _ := range envs {
					if k == "default" {
						continue
					}
					module := analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), "deployment-helmfile-"+k, analyzerapi.DeploymentSpecHelmfile, "helmfile", k)
					module.Repositories = repositories
					analyzerapi.AddModuleToResult(&result, module)
				}
			} else {
				module := analyzerapi.CreateProjectBuildSystemModule(ctx, file, a.GetName(), analyzerapi.BuildSystemHelmfile)
				module.Repositories = repositories
				analyzerapi.AddModuleToResult(&result, module)
			}
		} else if strings.HasPrefix(filename, "helmfile-") && strings.HasSuffix(filename, ".yaml.gotmpl") {
			envName := strings.TrimSuffix(strings.TrimPrefix(filename, "helmfile-"), ".yaml.gotmpl")
			module := analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), "deployment-helmfile-"+envName, analyzerapi.DeploymentSpecHelmfile, "helmfile", envName)
			module.Repositories = repositories
			analyzerapi.AddModuleToResult(&result, module)
		}
	}

	return result
}
