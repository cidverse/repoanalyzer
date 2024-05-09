package gradle

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "gradle"
}

func (a Analyzer) Analyze(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.Files {
		filename := filepath.Base(file)

		// detect build system syntax
		var buildSystemSyntax analyzerapi.ProjectBuildSystemSyntax
		var buildGradle BuildGradle
		if filename == "build.gradle" {
			buildSystemSyntax = analyzerapi.BuildSystemSyntaxGradleGroovyDSL
			var buildGradleErr error
			buildGradle, buildGradleErr = ParseBuildGradleGroovy(file)
			if buildGradleErr != nil {
				if os.Getenv("DEBUG") == "true" || os.Getenv("REPOANAYLZER_DEBUG") == "true" {
					fmt.Printf("%v", buildGradleErr)
				}
			}
		} else if filename == "build.gradle.kts" {
			buildSystemSyntax = analyzerapi.BuildSystemSyntaxGradleKotlinDSL
			var buildGradleErr error
			buildGradle, buildGradleErr = ParseBuildGradleKotlin(file)
			if buildGradleErr != nil {
				if os.Getenv("DEBUG") == "true" || os.Getenv("REPOANAYLZER_DEBUG") == "true" {
					fmt.Printf("%v", buildGradleErr)
				}
			}
		} else {
			continue
		}

		// deps
		var dependencies []analyzerapi.ProjectDependency
		dependencies = append(dependencies, buildGradle.Dependencies...)

		// module
		module := analyzerapi.ProjectModule{
			RootDirectory:     ctx.ProjectDir,
			Directory:         filepath.Dir(file),
			Name:              filepath.Base(filepath.Dir(file)),
			Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
			Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
			BuildSystem:       analyzerapi.BuildSystemGradle,
			BuildSystemSyntax: buildSystemSyntax,
			Language:          buildGradle.Languages,
			Dependencies:      buildGradle.Dependencies,
			Submodules:        nil,
			Files:             ctx.Files,
			FilesByExtension:  ctx.FilesByExtension,
		}
		analyzerapi.AddModuleToResult(&result, &module)
	}

	return result
}
