package analyzerapi

var Analyzers []Scanner

// Scanner is the interface that needs to be implemented by all analyzers
type Scanner interface {
	// GetName returns the name of the analyzer
	GetName() string

	// Scan will retrieve information about the project
	Scan(ctx AnalyzerContext) []*ProjectModule
}

// ProjectModule contains information about project modules
type ProjectModule struct {
	ID                    string                     `json:"id"`                     // ID is a unique identifier for the module
	RootDirectory         string                     `json:"project_dir"`            // RootDirectory stores the project root directory
	Directory             string                     `json:"module_dir"`             // Directory stores the module root directory
	Discovery             []ProjectModuleDiscovery   `json:"discovery"`              // Discovery stores information on how this module was discovered
	Name                  string                     `json:"name"`                   // Name stores the module name
	Slug                  string                     `json:"slug"`                   // Slug contains an url/folder name compatible name of the module
	Type                  ModuleType                 `json:"type"`                   // Type of the module
	BuildSystem           ProjectBuildSystem         `json:"build_system"`           // BuildSystem used in this project, only applies to modules of type build_system
	BuildSystemSyntax     ProjectBuildSystemSyntax   `json:"build_system_syntax"`    // BuildSystemSyntax used in this project, only applies to modules of type build_system
	SpecificationType     SpecificationType          `json:"specification_type"`     // SpecificationType that was found, only applies to modules of type spec
	ConfigType            ConfigType                 `json:"config_type"`            // ConfigType that was found, only applies to modules of type config
	DeploymentSpec        DeploymentSpec             `json:"deployment_spec"`        // DeploymentSpec is the type of spec describing the deployment, e.g. dotenv, ...
	DeploymentType        string                     `json:"deployment_type"`        // DeploymentType is the type of deployment, e.g. ansible, helm, kustomize, ...
	DeploymentEnvironment string                     `json:"deployment_environment"` // DeploymentEnvironment is the environment the deployment is for, e.g. staging, production, ...
	Language              map[ProjectLanguage]string `json:"language"`               // Language of the project
	Dependencies          []ProjectDependency        `json:"dependencies"`           // Dependencies
	Submodules            []*ProjectModule           `json:"submodules"`             // Submodules contains information about submodules
	Files                 []string                   `json:"files"`                  // Files holds all project files
	FilesByExtension      map[string][]string        `json:"-"`                      // FilesByExtension contains all files by extension
}

// ProjectModuleDiscovery contains info on the files used to discover the module
type ProjectModuleDiscovery struct {
	File string `json:"file"`
}

type ModuleType string

const (
	ModuleTypeBuildSystem ModuleType = "build_system" // e.g. Go Mo, Java, Python, Helm Charts, Ansible Playbooks, ...
	ModuleTypeSpec        ModuleType = "spec"         // e.g. OpenAPI, AsyncAPI, ...
	ModuleTypeConfig      ModuleType = "config"       // e.g. .env, .gitlab-ci.yml, ...
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
	BuildSystemPyprojectPoetry ProjectBuildSystem = "pyproject-poetry"
	BuildSystemPyprojectUV     ProjectBuildSystem = "pyproject-uv"
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
)

type ConfigType string

const (
	ConfigTypeRenovate       ConfigType = "renovate"
	ConfigTypeGitHubWorkflow ConfigType = "github-workflow"
	ConfigTypeGitLabCI       ConfigType = "gitlab-ci"
)

type DeploymentSpec string

const (
	DeploymentSpecDotEnv   DeploymentSpec = "deployment-dotenv"
	DeploymentSpecHelmfile DeploymentSpec = "deployment-helmfile"
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
