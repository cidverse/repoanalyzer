package apispec

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cidverse/cidverseutils/filesystem"
	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "apispec"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.Files {
		filename := filepath.Base(file)

		if filename == "openapi.json" || strings.HasSuffix(filename, ".openapi.json") {
			version := "0.0.0"
			if out, err := parseSpecFile(file); err == nil {
				version = out
			}

			module := analyzerapi.ProjectModule{
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeSpec,
				SpecificationType: analyzerapi.SpecificationTypeOpenAPI,
				Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguageOpenAPI, version),
				Dependencies:      nil,
				Submodules:        nil,
				Files:             ctx.Files,
				FilesByExtension:  ctx.FilesByExtension,
			}
			analyzerapi.AddModuleToResult(&result, &module)
		} else if filename == "openapi.yaml" || strings.HasSuffix(filename, ".openapi.yaml") {
			version := "0.0.0"
			if out, err := parseSpecFile(file); err == nil {
				version = out
			}

			module := analyzerapi.ProjectModule{
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeSpec,
				SpecificationType: analyzerapi.SpecificationTypeOpenAPI,
				Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguageOpenAPI, version),
				Dependencies:      nil,
				Submodules:        nil,
				Files:             ctx.Files,
				FilesByExtension:  ctx.FilesByExtension,
			}
			analyzerapi.AddModuleToResult(&result, &module)
		} else if filename == "asyncapi.yml" || strings.HasSuffix(filename, ".asyncapi.yml") || strings.HasSuffix(filename, ".asyncapi.yaml") {
			version := "0.0.0"
			if out, err := parseSpecFile(file); err == nil {
				version = out
			}

			module := analyzerapi.ProjectModule{
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeSpec,
				SpecificationType: analyzerapi.SpecificationTypeAsyncAPI,
				Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguageAsyncAPI, version),
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

func parseSpecFile(file string) (string, error) {
	content, err := filesystem.GetFileContent(file)
	if err != nil {
		return "", fmt.Errorf("failed to open spec file: %w", err)
	}

	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "openapi:") {
			return strings.TrimSpace(line[8:]), nil
		} else if strings.HasPrefix(line, "swagger:") {
			return strings.TrimSpace(line[8:]), nil
		} else if strings.HasPrefix(line, "asyncapi:") {
			return strings.TrimSpace(line[9:]), nil
		}
	}

	return "0.0.0", nil
}
