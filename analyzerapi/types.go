package analyzerapi

var Analyzers []Scanner

// Scanner is the interface that needs to be implemented by all analyzers
type Scanner interface {
	// GetName returns the name of the analyzer
	GetName() string

	// Analyze will retrieve information about the project
	Scan(ctx AnalyzerContext) []*ProjectModule
}

// ProjectModule contains information about project modules
type ProjectModule struct {
	// RootDirectory stores the project root directory
	RootDirectory string `json:"project_dir"`

	// Directory stores the module root directory
	Directory string `json:"module_dir"`

	// Discovery stores information on how this module was discovered
	Discovery []ProjectModuleDiscovery `json:"discovery"`

	// Name stores the module name
	Name string `json:"name"`

	// Slug contains an url/folder name compatible name of the module
	Slug string `json:"slug"`

	// Type of the module
	Type ModuleType `json:"type"`

	// BuildSystem used in this project, only applies to modules of type build_system
	BuildSystem ProjectBuildSystem `json:"build_system"`

	// BuildSystemSyntax used in this project, only applies to modules of type build_system
	BuildSystemSyntax ProjectBuildSystemSyntax `json:"build_system_syntax"`

	// SpecificationType that was found, only applies to modules of type spec
	SpecificationType SpecificationType `json:"specification_type"`

	// Language of the project
	Language map[ProjectLanguage]string `json:"language"`

	// Dependencies
	Dependencies []ProjectDependency `json:"dependencies"`

	// Submodules contains information about submodules
	Submodules []*ProjectModule `json:"submodules"`

	// Files holds all project files
	Files []string `json:"files"`

	// FilesByExtension contains all files by extension
	FilesByExtension map[string][]string `json:"-"`
}

// ProjectModuleDiscovery contains info on the files used to discover the module
type ProjectModuleDiscovery struct {
	File string `json:"file"`
}

type ModuleType string

const (
	ModuleTypeBuildSystem ModuleType = "build_system" // e.g. Go Mo, Java, Python, Helm Charts, Ansible Playbooks, ...
	ModuleTypeSpec        ModuleType = "spec"         // e.g. OpenAPI, AsyncAPI, ...
	ModuleTypeDeployment  ModuleType = "deployment"   // e.g. Helm Deployment Configuration, Ansible Deployment Configuration, ...
)

type ProjectLanguage string

const (
	LanguageGolang     ProjectLanguage = "go"
	LanguageJava       ProjectLanguage = "java"
	LanguageKotlin     ProjectLanguage = "kotlin"
	LanguageJavascript ProjectLanguage = "javascript"
	LanguageTypescript ProjectLanguage = "typescript"
	LanguagePython     ProjectLanguage = "python"
	LanguagePHP        ProjectLanguage = "php"
	LanguageRust       ProjectLanguage = "rust"
	LanguageNix        ProjectLanguage = "nix"
	LanguageOpenAPI    ProjectLanguage = "openapi"
	LanguageAsyncAPI   ProjectLanguage = "asyncapi"
)

type ProjectBuildSystem string

const (
	BuildSystemDefault         ProjectBuildSystem = "default"
	BuildSystemGradle          ProjectBuildSystem = "gradle"
	BuildSystemMaven           ProjectBuildSystem = "maven"
	BuildSystemGoMod           ProjectBuildSystem = "gomod"
	BuildSystemNpm             ProjectBuildSystem = "npm"
	BuildSystemHugo            ProjectBuildSystem = "hugo"
	BuildSystemHelm            ProjectBuildSystem = "helm"
	BuildSystemHelmfile        ProjectBuildSystem = "helmfile"
	BuildSystemContainer       ProjectBuildSystem = "container"
	BuildSystemRequirementsTXT ProjectBuildSystem = "python-requirements.txt"
	BuildSystemPipfile         ProjectBuildSystem = "pipfile"
	BuildSystemSetupPy         ProjectBuildSystem = "setup.py"
	BuildSystemPoetry          ProjectBuildSystem = "poetry"
	BuildSystemMkdocs          ProjectBuildSystem = "mkdocs"
	BuildSystemComposer        ProjectBuildSystem = "composer"
	BuildSystemDotNet          ProjectBuildSystem = "dotnet"
	BuildSystemCargo           ProjectBuildSystem = "cargo"
	BuildSystemNix             ProjectBuildSystem = "nix"
	BuildSystemAnsible         ProjectBuildSystem = "ansible"
	BuildSystemQuartz          ProjectBuildSystem = "quartz"
)

type ProjectBuildSystemSyntax string

const (
	BuildSystemSyntaxDefault                ProjectBuildSystemSyntax = "default"
	BuildSystemSyntaxGradleGroovyDSL        ProjectBuildSystemSyntax = "groovy"
	BuildSystemSyntaxGradleKotlinDSL        ProjectBuildSystemSyntax = "kotlin"
	BuildSystemSyntaxDotNetSLN              ProjectBuildSystemSyntax = "sln"
	BuildSystemSyntaxDotNetSLNX             ProjectBuildSystemSyntax = "slnx"
	BuildSystemSyntaxDotNetCSProj           ProjectBuildSystemSyntax = "csproj"
	BuildSystemSyntaxContainerFile          ProjectBuildSystemSyntax = "containerfile"
	BuildSystemSyntaxContainerBuildahScript ProjectBuildSystemSyntax = "buildah-script"
	BuildSystemSyntaxMkdocsTechdocs         ProjectBuildSystemSyntax = "mkdocs-techdocs"
	BuildSystemSyntaxNixFlake               ProjectBuildSystemSyntax = "flake"
)

type SpecificationType string

const (
	SpecificationTypeOpenAPI  SpecificationType = "openapi"
	SpecificationTypeAsyncAPI SpecificationType = "asyncapi"
	SpecificationTypeRenovate SpecificationType = "renovate"
)

// ProjectDependency contains dependency information
type ProjectDependency struct {
	// Type is the dep kind
	Type string `json:"type"`

	// ID is the identifier
	ID string `json:"id"`

	// Version is the dep version
	Version string `json:"version"`
}

// AnalyzerContext holds the context to analyze projects
type AnalyzerContext struct {
	// ProjectDir holds the project directory
	ProjectDir string `json:"project_dir"`

	// Files holds all project files
	Files []string `json:"files"`

	// FilesByExtension contains all files by extension
	FilesByExtension map[string][]string `json:"files_by_extension"`
}

func (ctx *AnalyzerContext) ContainsFile(extension string) bool {
	_, ok := ctx.FilesByExtension[extension]
	return ok
}
