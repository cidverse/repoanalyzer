package maven

import (
	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
	"path/filepath"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "maven"
}

func (a Analyzer) Analyze(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.Files {
		filename := filepath.Base(file)
		if filename != "pom.xml" {
			continue
		}

		// language
		language := make(map[analyzerapi.ProjectLanguage]*string)
		language[analyzerapi.LanguageJava] = nil

		// deps
		var dependencies []analyzerapi.ProjectDependency

		// module
		module := analyzerapi.ProjectModule{
			RootDirectory:     ctx.ProjectDir,
			Directory:         filepath.Dir(file),
			Name:              filepath.Base(filepath.Dir(file)),
			Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
			Discovery:         []string{"file~" + file},
			BuildSystem:       analyzerapi.BuildSystemMaven,
			BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
			Language:          language,
			Dependencies:      dependencies,
			Submodules:        nil,
			Files:             ctx.Files,
			FilesByExtension:  ctx.FilesByExtension,
		}
		analyzerapi.AddModuleToResult(&result, &module)
	}

	return result
}
