package analyzer

import (
	"sort"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// buildSystemPriority defines the specificity of a build system, a higher value is more specific
var buildSystemPriority = map[analyzerapi.ProjectBuildSystem]int{
	analyzerapi.BuildSystemRspress:         100,
	analyzerapi.BuildSystemQuartz:          90,
	analyzerapi.BuildSystemNpm:             10,
	analyzerapi.BuildSystemGradle:          100,
	analyzerapi.BuildSystemMaven:           100,
	analyzerapi.BuildSystemGoMod:           100,
	analyzerapi.BuildSystemPyprojectPoetry: 100,
	analyzerapi.BuildSystemPyprojectUV:     100,
	analyzerapi.BuildSystemSetupPy:         80,
	analyzerapi.BuildSystemPipfile:         70,
	analyzerapi.BuildSystemRequirementsTXT: 60,
	analyzerapi.BuildSystemCargo:           100,
	analyzerapi.BuildSystemComposer:        100,
	analyzerapi.BuildSystemDotNet:          100,
}

// DeduplicateModules removes duplicate modules, only one module per directory and category may exist (the most specific one)
func DeduplicateModules(modules []*analyzerapi.ProjectModule) []*analyzerapi.ProjectModule {
	// group build system modules by directory and category
	grouped := make(map[string][]*analyzerapi.ProjectModule)
	for _, module := range modules {
		if module.Type != analyzerapi.ModuleTypeBuildSystem || module.Category == "" {
			continue
		}
		key := module.Directory + "\x00" + string(module.Category)
		grouped[key] = append(grouped[key], module)
	}

	// keep the most specific module per group, drop the rest
	drop := make(map[*analyzerapi.ProjectModule]bool)
	for _, candidates := range grouped {
		if len(candidates) <= 1 {
			continue
		}

		sort.SliceStable(candidates, func(i, j int) bool {
			return buildSystemPriority[candidates[i].BuildSystem] > buildSystemPriority[candidates[j].BuildSystem]
		})
		for _, m := range candidates[1:] {
			drop[m] = true
		}
	}

	// rebuild result preserving the original order
	result := make([]*analyzerapi.ProjectModule, 0, len(modules))
	for _, module := range modules {
		if drop[module] {
			continue
		}
		result = append(result, module)
	}

	return result
}
