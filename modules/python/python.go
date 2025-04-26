package python

import (
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"
	"strings"

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
			pyProject, err := readPyProjectFile(file)
			if err != nil {
				slog.With("err", err).Warn("Error reading Python project file")
				continue
			}

			var dependencies []analyzerapi.ProjectDependency
			if _, err = os.Stat(filepath.Join(filepath.Dir(file), "uv.lock")); err == nil {
				if pyProject.DependencyGroups != nil {
					for group, deps := range pyProject.DependencyGroups {
						for _, dep := range deps {
							name, version := splitDependency(dep)
							if strings.HasPrefix(version, "==") {
								version = strings.TrimPrefix(version, "==")
							}

							dependencies = append(dependencies, analyzerapi.ProjectDependency{
								Type:    "pypi",
								ID:      name,
								Version: version,
								Scope:   group,
							})
						}
					}
				}

				module := analyzerapi.ProjectModule{
					ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
					RootDirectory:     ctx.ProjectDir,
					Directory:         filepath.Dir(file),
					Name:              pyProject.Project.Name,
					Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
					Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
					Type:              analyzerapi.ModuleTypeBuildSystem,
					BuildSystem:       analyzerapi.BuildSystemPyprojectUV,
					BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
					Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguagePython, "0.0.0"),
					Dependencies:      dependencies,
					Submodules:        nil,
					Files:             ctx.Files,
					FilesByExtension:  ctx.FilesByExtension,
				}
				analyzerapi.AddModuleToResult(&result, &module)
			} else if _, err = os.Stat(filepath.Join(filepath.Dir(file), "poetry.lock")); err == nil {
				if pyProject.Tool.Poetry.Dependencies != nil {
					for dep, version := range pyProject.Tool.Poetry.Dependencies {
						dependencies = append(dependencies, analyzerapi.ProjectDependency{
							Type:    "pypi",
							ID:      dep,
							Version: version,
						})
					}
				}
				if pyProject.Tool.Poetry.Groups != nil {
					for group, groupDeps := range pyProject.Tool.Poetry.Groups {
						for dep, version := range groupDeps.Dependencies {
							dependencies = append(dependencies, analyzerapi.ProjectDependency{
								Type:    "pypi",
								ID:      dep,
								Version: version,
								Scope:   group,
							})
						}
					}
				}

				module := analyzerapi.ProjectModule{
					ID:                analyzerapi.GetSlugFromPath(ctx.ProjectDir, file, a.GetName()),
					RootDirectory:     ctx.ProjectDir,
					Directory:         filepath.Dir(file),
					Name:              pyProject.Tool.Poetry.Name,
					Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
					Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
					Type:              analyzerapi.ModuleTypeBuildSystem,
					BuildSystem:       analyzerapi.BuildSystemPyprojectPoetry,
					BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
					Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguagePython, "0.0.0"),
					Dependencies:      dependencies,
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
