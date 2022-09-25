package container

import (
	"github.com/cidverse/repoanalyzer/logger"
	"github.com/cidverse/repoanalyzer/util"
	"golang.org/x/exp/slices"
	"path/filepath"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "container"
}

func (a Analyzer) Analyze(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	// dockerfile
	for _, file := range ctx.Files {
		filename := filepath.Base(file)

		if filename == "Dockerfile" || filename == "Containerfile" || strings.HasSuffix(filename, ".Dockerfile") || strings.HasSuffix(filename, ".Containerfile") {
			// add new module or append file to existing module
			moduleIdx := slices.IndexFunc(result, func(m *analyzerapi.ProjectModule) bool {
				return m.Name == filepath.Base(filepath.Dir(file)) && m.BuildSystem == analyzerapi.BuildSystemContainer && m.BuildSystemSyntax == analyzerapi.ContainerFile
			})
			if moduleIdx == -1 {
				module := analyzerapi.ProjectModule{
					RootDirectory:     ctx.ProjectDir,
					Directory:         filepath.Dir(file),
					Name:              filepath.Base(filepath.Dir(file)),
					Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
					Discovery:         []string{"file~" + file},
					BuildSystem:       analyzerapi.BuildSystemContainer,
					BuildSystemSyntax: analyzerapi.ContainerFile,
					Language:          nil,
					Dependencies:      nil,
					Submodules:        nil,
					Files:             ctx.Files,
					FilesByExtension:  ctx.FilesByExtension,
				}
				analyzerapi.AddModuleToResult(&result, &module)
			} else {
				result[moduleIdx].Discovery = append(result[moduleIdx].Discovery, "file~"+file)
			}
		}
	}

	// buildah
	for _, file := range ctx.FilesByExtension["sh"] {
		filename := filepath.Base(file)

		if strings.HasSuffix(filename, ".sh") {
			content, contentErr := util.GetFileContent(file)
			if contentErr != nil {
				logger.Warn("failed to read file content", "file", file)
			} else if strings.Contains(content, "buildah from") {
				module := analyzerapi.ProjectModule{
					RootDirectory:     ctx.ProjectDir,
					Directory:         filepath.Dir(file),
					Name:              filepath.Base(filepath.Dir(file)),
					Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
					Discovery:         []string{"file~" + file},
					BuildSystem:       analyzerapi.BuildSystemContainer,
					BuildSystemSyntax: analyzerapi.ContainerBuildahScript,
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
