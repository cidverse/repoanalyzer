package repoanalyzer

import (
	"strings"
	"time"

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
	"github.com/cidverse/repoanalyzer/analyzer/node"
	"github.com/cidverse/repoanalyzer/analyzer/python"
	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/rs/zerolog/log"
)

var analyzerCache = make(map[string][]*analyzerapi.ProjectModule)

// AnalyzeProject will analyze a project and return all modules in path
func AnalyzeProject(projectDir string, path string) []*analyzerapi.ProjectModule {
	if _, ok := analyzerCache[path]; ok {
		return analyzerCache[path]
	}

	if len(analyzerapi.Analyzers) == 0 {
		initAnalyzers()
	}

	start := time.Now()
	log.Debug().Str("path", path).Int("scanners", len(analyzerapi.Analyzers)).Msg("repo analyzer start")

	// prepare context
	ctx := analyzerapi.GetAnalyzerContext(projectDir)

	// run
	var allModules []*analyzerapi.ProjectModule
	var allModuleNames []string
	for _, a := range analyzerapi.Analyzers {
		log.Trace().Str("name", a.GetName()).Msg("repo analyzer run")
		modules := a.Analyze(ctx)
		for _, module := range modules {
			currentModule := module
			if strings.HasPrefix(currentModule.Directory, path) && !strings.Contains(currentModule.Directory, "testdata") {
				allModules = append(allModules, currentModule)
				allModuleNames = append(allModuleNames, currentModule.Slug+" ("+string(currentModule.BuildSystem)+")")
			}
		}
	}

	log.Debug().Str("path", path).Int("module_count", len(allModules)).Strs("modules", allModuleNames).Str("duration", time.Since(start).String()).Int("file_count", len(ctx.Files)).Msg("repo analyzer complete")

	analyzerCache[projectDir] = allModules
	return allModules
}

func initAnalyzers() {
	analyzerapi.Analyzers = append(analyzerapi.Analyzers,
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
		node.Analyzer{},
		python.Analyzer{},
	)
}
