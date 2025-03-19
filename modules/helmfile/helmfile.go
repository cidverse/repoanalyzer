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
		isChart := strings.Contains(filepath.Dir(file), "/charts/")

		if filename == "helmfile.yaml" {
			if isChart {
				analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectBuildSystemModule(ctx, file, a.GetName(), analyzerapi.BuildSystemHelmfile))
			} else {
				analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), analyzerapi.DeploymentSpecHelmfile, "helmfile"))
			}
		}
	}

	for _, file := range ctx.FilesByExtension["gotmpl"] {
		filename := filepath.Base(file)
		isChart := strings.Contains(filepath.Dir(file), "/charts/")

		if filename == "helmfile.yaml.gotmpl" {
			if isChart {
				analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectBuildSystemModule(ctx, file, a.GetName(), analyzerapi.BuildSystemHelmfile))
			} else {
				analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), analyzerapi.DeploymentSpecHelmfile, "helmfile"))
			}
		}
	}

	return result
}
