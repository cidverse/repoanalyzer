package composer

import (
	"path/filepath"
	"regexp"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "composer"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.FilesByExtension["json"] {
		filename := filepath.Base(file)
		if filename == "composer.json" {
			// language
			language := make(map[analyzerapi.ProjectLanguage]string)
			language[analyzerapi.LanguagePHP] = "0.0.0"

			// parse composer.json
			composerJSON, err := ParseComposerJSON(file)
			if err == nil {
				if val, ok := composerJSON.Require["php"]; ok {
					language[analyzerapi.LanguagePHP] = regexp.MustCompile("[^0-9.]").ReplaceAllString(val, "")
				}
			}

			module := analyzerapi.ProjectModule{
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				Type:              analyzerapi.ModuleTypeBuildSystem,
				BuildSystem:       analyzerapi.BuildSystemComposer,
				BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
				Language:          language,
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
