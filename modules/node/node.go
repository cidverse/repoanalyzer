package node

import (
	"path/filepath"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "node"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	// iterate
	for _, file := range ctx.FilesByExtension["json"] {
		filename := filepath.Base(file)

		// detect build system syntax
		if filename == "package.json" {
			packageData, packageDataErr := ParsePackageJSON(file)
			if packageDataErr != nil {
				continue
			}

			// skip package.json files without a usable module name
			moduleName := strings.TrimSpace(packageData.Name)
			moduleSlug := slug.Make(moduleName)
			if moduleName == "" || moduleSlug == "" {
				continue
			}

			// language
			language := make(map[analyzerapi.ProjectLanguage]string)
			if _, ok := packageData.Dependencies["typescript"]; ok {
				typescriptVersion := packageData.Dependencies["typescript"]
				language[analyzerapi.LanguageTypescript] = typescriptVersion
			} else {
				language[analyzerapi.LanguageJavascript] = "0.0.0"
			}

			// deps
			var dependencies []analyzerapi.ProjectDependency
			for key, value := range packageData.Dependencies {
				dep := analyzerapi.ProjectDependency{
					Type:    string(analyzerapi.BuildSystemNpm),
					ID:      key,
					Version: value,
				}
				dependencies = append(dependencies, dep)
			}

			// module
			module := analyzerapi.ProjectModule{
				ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              moduleName,
				Slug:              moduleSlug,
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeBuildSystem,
				Category:          analyzerapi.ModuleCategoryNode,
				BuildSystem:       analyzerapi.BuildSystemNpm,
				BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
				Language:          language,
				Dependencies:      dependencies,
				Repositories:      npmRepositories(filepath.Dir(file), ctx.ProjectDir),
				Submodules:        nil,
				Files:             ctx.Files,
				FilesByExtension:  ctx.FilesByExtension,
			}

			analyzerapi.AddModuleToResult(&result, &module)
		} else {
			continue
		}
	}

	return result
}
