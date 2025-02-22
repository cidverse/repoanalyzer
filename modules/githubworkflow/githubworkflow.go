package githubworkflow

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return string(analyzerapi.ConfigTypeGitHubWorkflow)
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.Files {
		filename := filepath.Base(file)
		if !strings.HasPrefix(file, filepath.Join(ctx.ProjectDir, ".github", "workflows")) {
			continue
		}

		if strings.HasSuffix(filename, ".yml") || strings.HasSuffix(filename, ".yaml") {
			filenameNoExt := strings.TrimSuffix(filename, filepath.Ext(filename))
			module := analyzerapi.ProjectModule{
				ID:               analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
				RootDirectory:    ctx.ProjectDir,
				Directory:        path.Join(ctx.ProjectDir, ".github", "workflows"),
				Name:             "github-workflow-" + filenameNoExt,
				Slug:             "github-workflow-" + filenameNoExt,
				Discovery:        []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:             analyzerapi.ModuleTypeConfig,
				ConfigType:       analyzerapi.ConfigTypeGitHubWorkflow,
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
