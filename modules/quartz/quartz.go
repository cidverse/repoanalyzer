package quartz

import (
	"path/filepath"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

// Analyzer for the Quartz Static Site Generator - https://quartz.jzhao.xyz/
type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "quartz"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.Files {
		filename := filepath.Base(file)
		if filename == "quartz.config.ts" {
			module := analyzerapi.ProjectModule{
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeBuildSystem,
				BuildSystem:       analyzerapi.BuildSystemQuartz,
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
