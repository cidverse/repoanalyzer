package gitlabci

import (
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return string(analyzerapi.ConfigTypeGitLabCI)
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.Files {
		if strings.HasSuffix(file, "/.gitlab-ci.yml") {
			module := analyzerapi.ProjectModule{
				RootDirectory:    ctx.ProjectDir,
				Directory:        ctx.ProjectDir,
				Name:             "gitlab-ci",
				Slug:             "gitlab-ci",
				Discovery:        []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:             analyzerapi.ModuleTypeConfig,
				ConfigType:       analyzerapi.ConfigTypeGitLabCI,
				Language:         nil,
				Dependencies:     nil,
				Submodules:       nil,
				Files:            ctx.Files,
				FilesByExtension: ctx.FilesByExtension,
			}
			analyzerapi.AddModuleToResult(&result, &module)
		}
	}

	return result
}
