package repoanalyzer

import (
	"github.com/cidverse/repoanalyzer/analyzer/container"
	"github.com/cidverse/repoanalyzer/analyzer/gomod"
	"github.com/cidverse/repoanalyzer/analyzer/gradle"
	"github.com/cidverse/repoanalyzer/analyzer/helm"
	"github.com/cidverse/repoanalyzer/analyzer/hugo"
	"github.com/cidverse/repoanalyzer/analyzer/node"
	"github.com/cidverse/repoanalyzer/analyzer/python"
	"github.com/cidverse/repoanalyzer/logger"
	"github.com/thoas/go-funk"
	"strings"
	"time"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

var analyzerCache = make(map[string][]*analyzerapi.ProjectModule)

// AnalyzeProject will analyze a project and return all modules in path
func AnalyzeProject(projectDir string, path string) []*analyzerapi.ProjectModule {
	if funk.Contains(analyzerCache, path) {
		return analyzerCache[path]
	}

	if len(analyzerapi.Analyzers) == 0 {
		initAnalyzers()
	}

	start := time.Now()
	logger.Logger.Info("repo analyzer start", "path", path, "scanners", len(analyzerapi.Analyzers))

	// prepare context
	ctx := analyzerapi.GetAnalyzerContext(projectDir)

	// run
	var allModules []*analyzerapi.ProjectModule
	var allModuleNames []string
	for _, a := range analyzerapi.Analyzers {
		logger.Debug("repo analyzer run", "name", a.GetName())

		modules := a.Analyze(ctx)
		for _, module := range modules {
			currentModule := module
			if strings.HasPrefix(currentModule.Directory, path) && !strings.Contains(currentModule.Directory, "testdata") {
				allModules = append(allModules, currentModule)
				allModuleNames = append(allModuleNames, currentModule.Slug)
			}
		}
	}

	logger.Info("repo analyzer complete", "module_count", len(allModules), "modules", allModuleNames, "duration", time.Since(start).String(), "file_count", len(ctx.Files))

	analyzerCache[projectDir] = allModules
	return allModules
}

func initAnalyzers() {
	analyzerapi.Analyzers = append(analyzerapi.Analyzers,
		container.Analyzer{},
		gomod.Analyzer{},
		gradle.Analyzer{},
		helm.Analyzer{},
		hugo.Analyzer{},
		node.Analyzer{},
		python.Analyzer{},
	)
}
