package python

import (
	"path/filepath"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "python"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	// iterate
	for _, file := range ctx.Files {
		filename := filepath.Base(file)

		// detect build system syntax
		if filename == "requirements.txt" {
			module := analyzerapi.ProjectModule{
				ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeBuildSystem,
				BuildSystem:       analyzerapi.BuildSystemRequirementsTXT,
				BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
				Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguagePython, "0.0.0"),
				Dependencies:      nil,
				Submodules:        nil,
				Files:             ctx.Files,
				FilesByExtension:  ctx.FilesByExtension,
			}
			analyzerapi.AddModuleToResult(&result, &module)
		} else if filename == "Pipfile" {
			module := analyzerapi.ProjectModule{
				ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeBuildSystem,
				BuildSystem:       analyzerapi.BuildSystemPipfile,
				BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
				Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguagePython, "0.0.0"),
				Dependencies:      nil,
				Submodules:        nil,
				Files:             ctx.Files,
				FilesByExtension:  ctx.FilesByExtension,
			}
			analyzerapi.AddModuleToResult(&result, &module)
		} else if filename == "setup.py" {
			module := analyzerapi.ProjectModule{
				ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeBuildSystem,
				BuildSystem:       analyzerapi.BuildSystemSetupPy,
				BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
				Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguagePython, "0.0.0"),
				Dependencies:      nil,
				Submodules:        nil,
				Files:             ctx.Files,
				FilesByExtension:  ctx.FilesByExtension,
			}
			analyzerapi.AddModuleToResult(&result, &module)
		} else if filename == "pyproject.toml" {
			module := analyzerapi.ProjectModule{
				ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeBuildSystem,
				BuildSystem:       analyzerapi.BuildSystemPoetry,
				BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
				Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguagePython, "0.0.0"),
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
