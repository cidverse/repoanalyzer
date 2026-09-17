package rspress

import (
	"path/filepath"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// Analyzer for the Rspress Static Site Generator - https://rspress.dev/
type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "rspress"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.Files {
		filename := filepath.Base(file)
		switch filename {
		case "rspress.config.ts", "rspress.config.mjs", "rspress.config.js", "rspress.config.mts", "rspress.config.cts":
			module := analyzerapi.CreateProjectBuildSystemModule(ctx, file, a.GetName(), analyzerapi.BuildSystemRspress)
			module.Category = analyzerapi.ModuleCategoryNode
			analyzerapi.AddModuleToResult(&result, module)
		}
	}

	return result
}
