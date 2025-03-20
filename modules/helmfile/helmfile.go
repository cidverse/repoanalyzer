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

		if filename == "helmfile.yaml" {
			envs, err := parseHelmfileEnvironments(file)
			if err == nil {
				for k, _ := range envs {
					if k == "default" {
						continue
					}
					analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), "deployment-helmfile-"+k, analyzerapi.DeploymentSpecHelmfile, "helmfile", k))
				}
			} else {
				analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectBuildSystemModule(ctx, file, a.GetName(), analyzerapi.BuildSystemHelmfile))
			}
		} else if strings.HasPrefix(filename, "helmfile-") && strings.HasSuffix(filename, ".yaml") {
			envName := strings.TrimSuffix(strings.TrimPrefix(filename, "helmfile-"), ".yaml")
			analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), "deployment-helmfile-"+envName, analyzerapi.DeploymentSpecHelmfile, "helmfile", envName))
		}
	}

	for _, file := range ctx.FilesByExtension["gotmpl"] {
		filename := filepath.Base(file)

		if filename == "helmfile.yaml.gotmpl" {
			analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectBuildSystemModule(ctx, file, a.GetName(), analyzerapi.BuildSystemHelmfile))
		} else if strings.HasPrefix(filename, "helmfile-") && strings.HasSuffix(filename, ".yaml.gotmpl") {
			envName := strings.TrimSuffix(strings.TrimPrefix(filename, "helmfile-"), ".yaml.gotmpl")
			analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), "deployment-helmfile-"+envName, analyzerapi.DeploymentSpecHelmfile, "helmfile", envName))
		}
	}

	return result
}
