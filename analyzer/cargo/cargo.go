package cargo

import (
	"path/filepath"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "cargo"
}

func (a Analyzer) Analyze(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	for _, file := range ctx.FilesByExtension["toml"] {
		filename := filepath.Base(file)
		if filename == "Cargo.toml" {
			// parse Cargo.toml
			cargoFile, err := parseCargoFile(file)
			if err != nil {
				log.Debug().Err(err).Msg("failed to parse Cargo.toml")
			}

			// module
			module := analyzerapi.ProjectModule{
				RootDirectory:     ctx.ProjectDir,
				Directory:         filepath.Dir(file),
				Name:              filepath.Base(filepath.Dir(file)),
				Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
				Discovery:         []analyzerapi.ProjectModuleDiscovery{{File: file}},
				BuildSystem:       analyzerapi.BuildSystemCargo,
				BuildSystemSyntax: analyzerapi.BuildSystemSyntaxDefault,
				Language:          analyzerapi.GetSingleLanguageMap(analyzerapi.LanguageRust, getRustVersionFromCargoFile(cargoFile)),
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

func getRustVersionFromCargoFile(cargoFile Config) string {
	if cargoFile.Package.RustVersion != "" {
		return cargoFile.Package.RustVersion + ".0"
	}
	return "0.0.0"
}
