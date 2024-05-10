package analyzer

import (
	"strings"
	"time"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"golang.org/x/exp/slog"
)

type RepoAnalyzer interface {
	Add(scanner analyzerapi.Scanner)
	List() []analyzerapi.Scanner
	Scan(path string) []*analyzerapi.ProjectModule
	EnableCache(enable bool)
}

type Analyzer struct {
	scanners    []analyzerapi.Scanner
	enableCache bool
	cache       map[string][]*analyzerapi.ProjectModule
}

func (a Analyzer) Add(analyzer analyzerapi.Scanner) {
	a.scanners = append(a.scanners, analyzer)
}

func (a Analyzer) List() []analyzerapi.Scanner {
	return a.scanners
}

func (a Analyzer) Scan(dir string) []*analyzerapi.ProjectModule {
	if a.enableCache {
		if _, ok := a.cache[dir]; ok {
			return a.cache[dir]
		}
	}

	start := time.Now()
	slog.Debug("repo analyzer start", slog.String("path", dir), slog.Int("scanners", len(a.scanners)))

	// prepare context
	ctx := analyzerapi.GetAnalyzerContext(dir)

	// run
	var allModules []*analyzerapi.ProjectModule
	var allModuleNames []string
	for _, a := range a.scanners {
		slog.Debug("repo analyzer run", slog.String("name", a.GetName()))
		modules := a.Scan(ctx)
		for _, module := range modules {
			currentModule := module
			if !strings.Contains(currentModule.Directory, "testdata") {
				allModules = append(allModules, currentModule)
				allModuleNames = append(allModuleNames, currentModule.Slug+" ("+string(currentModule.BuildSystem)+")")
			}
		}
	}

	slog.Debug("repo analyzer complete", slog.String("path", dir), slog.Int("module_count", len(allModules)), slog.String("modules", strings.Join(allModuleNames, ",")), slog.String("duration", time.Since(start).String()), slog.Int("file_count", len(ctx.Files)))

	if a.enableCache {
		a.cache[dir] = allModules
	}
	return allModules
}

func (a Analyzer) EnableCache(enable bool) {
	a.enableCache = enable
}

// NewAnalyzer returns a new analyzer with all available modules
func NewAnalyzer() Analyzer {
	return Analyzer{
		scanners:    AllScanners,
		enableCache: false,
		cache:       make(map[string][]*analyzerapi.ProjectModule),
	}
}

// NewCustomAnalyzer returns an analyzer without any pre-configured modules
func NewCustomAnalyzer() Analyzer {
	return Analyzer{
		scanners:    []analyzerapi.Scanner{},
		enableCache: false,
		cache:       make(map[string][]*analyzerapi.ProjectModule),
	}
}

var globalAnalyzer *Analyzer

// ScanDirectory scans a directory using the global analyzer instance
func ScanDirectory(path string) []*analyzerapi.ProjectModule {
	if globalAnalyzer == nil {
		a := NewAnalyzer()
		globalAnalyzer = &a
	}

	return globalAnalyzer.Scan(path)
}
