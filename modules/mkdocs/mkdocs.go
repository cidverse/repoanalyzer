package mkdocs

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "mkdocs"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range append(ctx.FilesByExtension["yml"], ctx.FilesByExtension["yaml"]...) {
		filename := filepath.Base(file)
		if filename == "mkdocs.yml" || filename == "mkdocs.yaml" {
			mkdocsDir := filepath.Dir(file)

			if _, err := os.Stat(filepath.Join(mkdocsDir, "catalog-info.yaml")); err == nil {
				// module
				module := analyzerapi.ProjectModule{
					ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
					RootDirectory:     ctx.ProjectDir,
					Directory:         filepath.Dir(file),
					Name:              filepath.Base(filepath.Dir(file)),
					Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
					Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
					Type:              analyzerapi.ModuleTypeBuildSystem,
					BuildSystem:       analyzerapi.BuildSystemMkdocs,
					BuildSystemSyntax: analyzerapi.BuildSystemSyntaxMkdocsTechdocs,
					Language:          nil,
					Dependencies:      nil,
					Submodules:        nil,
					Files:             ctx.Files,
					FilesByExtension:  ctx.FilesByExtension,
				}
				analyzerapi.AddModuleToResult(&result, &module)
			} else if errors.Is(err, os.ErrNotExist) {
				// module
				module := analyzerapi.ProjectModule{
					ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
					RootDirectory:     ctx.ProjectDir,
					Directory:         filepath.Dir(file),
					Name:              filepath.Base(filepath.Dir(file)),
					Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
					Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
					Type:              analyzerapi.ModuleTypeBuildSystem,
					BuildSystem:       analyzerapi.BuildSystemMkdocs,
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
	}

	return result
}
