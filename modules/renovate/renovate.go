package renovate

import (
	"path/filepath"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/gosimple/slug"
)

type Analyzer struct{}

func (a Analyzer) GetName() string {
	return "renovate"
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	renovateFiles := []string{
		"renovate.json",
		"renovate.json5",
		//".github/renovate.json",
		//".github/renovate.json5",
		//".gitlab/renovate.json",
		//".gitlab/renovate.json5",
		".renovaterc",
		".renovaterc.json",
		".renovaterc.json5",
	}
	for _, file := range ctx.Files {
		for _, renovateFile := range renovateFiles {
			if strings.EqualFold(filepath.Base(file), renovateFile) || strings.Contains(file, renovateFile) {
				name := filepath.Base(filepath.Dir(file))
				if name == ".github" || name == ".gitlab" {
					name = filepath.Base(ctx.ProjectDir)
				}
				module := analyzerapi.ProjectModule{
					RootDirectory:    ctx.ProjectDir,
					Directory:        filepath.Dir(file),
					Name:             name,
					Slug:             slug.Make(name),
					Discovery:        []analyzerapi.ProjectModuleDiscovery{{File: file}},
					Type:             analyzerapi.ModuleTypeConfig,
					ConfigType:       analyzerapi.ConfigTypeRenovate,
					Language:         nil,
					Dependencies:     nil,
					Submodules:       nil,
					Files:            ctx.Files,
					FilesByExtension: ctx.FilesByExtension,
				}
				analyzerapi.AddModuleToResult(&result, &module)
			}
		}
	}

	return result
}
