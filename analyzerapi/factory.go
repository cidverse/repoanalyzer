package analyzerapi

import (
	"path/filepath"

	"github.com/gosimple/slug"
)

func CreateProjectBuildSystemModule(ctx AnalyzerContext, file string, analyzerName string, buildSystem ProjectBuildSystem) *ProjectModule {
	return &ProjectModule{
		ID:                GetSlugFromPath(ctx.ProjectDir, file, analyzerName),
		RootDirectory:     ctx.ProjectDir,
		Directory:         filepath.Dir(file),
		Name:              filepath.Base(filepath.Dir(file)),
		Slug:              slug.Make(filepath.Base(filepath.Dir(file))),
		Discovery:         []ProjectModuleDiscovery{{File: file}},
		Type:              ModuleTypeBuildSystem,
		BuildSystem:       buildSystem,
		BuildSystemSyntax: BuildSystemSyntaxDefault,
		Language:          nil,
		Dependencies:      nil,
		Submodules:        nil,
		Files:             ctx.Files,
		FilesByExtension:  ctx.FilesByExtension,
	}
}

func CreateProjectDeploymentModule(ctx AnalyzerContext, file string, analyzerName string, name string, deploymentSpec DeploymentSpec, deploymentType string, envName string) *ProjectModule {
	return &ProjectModule{
		ID:                    GetSlugFromPath(ctx.ProjectDir, file, analyzerName),
		RootDirectory:         ctx.ProjectDir,
		Directory:             filepath.Dir(file),
		Name:                  name,
		Slug:                  slug.Make(name),
		Discovery:             []ProjectModuleDiscovery{{File: file}},
		Type:                  ModuleTypeDeployment,
		DeploymentSpec:        deploymentSpec,
		DeploymentType:        deploymentType,
		DeploymentEnvironment: envName,
		Language:              nil,
		Dependencies:          nil,
		Submodules:            nil,
		Files:                 ctx.Files,
		FilesByExtension:      ctx.FilesByExtension,
	}
}
