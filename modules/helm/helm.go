package helm

import (
	"path/filepath"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "helm"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.FilesByExtension["yaml"] {
		filename := filepath.Base(file)

		if filename == "Chart.yaml" {
			// module
			module := analyzerapi.ProjectModule{
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeBuildSystem,
				BuildSystem:       analyzerapi.BuildSystemHelm,
				BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
				Language:          nil,
				Dependencies:      nil,
				Submodules:        nil,
				Files:             ctx.Files,
				FilesByExtension:  ctx.FilesByExtension,
			}
			analyzerapi.AddModuleToResult(&result, &module)
		}
	}

	return result
}
