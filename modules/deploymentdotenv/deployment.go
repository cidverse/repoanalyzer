package deploymentdotenv

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"golang.org/x/exp/slices"
)

type Analyzer struct {
	AllowedEnvironmentNames []string
}

func (a Analyzer) GetName() string {
	return string(analyzerapi.DeploymentSpecDotEnv)
}

func (a Analyzer) Scan(ctx analyzerapi.AnalyzerContext) []*analyzerapi.ProjectModule {
	var result []*analyzerapi.ProjectModule

	// find project files
	for _, file := range ctx.Files {
		filename := filepath.Base(file)

		if strings.HasPrefix(filename, ".env-") && filepath.Dir(file) == ctx.ProjectDir {
			envName := strings.TrimPrefix(filename, ".env-")
			if len(a.AllowedEnvironmentNames) > 0 && !slices.Contains(a.AllowedEnvironmentNames, envName) {
				continue
			}

			content, err := os.ReadFile(file)
			if err != nil {
				continue
			}
			properties := parseDotEnvFile(content)
			deploymentType := properties["DEPLOYMENT_TYPE"]
			if deploymentType == "" {
				continue
			}

			analyzerapi.AddModuleToResult(&result, analyzerapi.CreateProjectDeploymentModule(ctx, file, a.GetName(), "deployment-dotenv-"+envName, analyzerapi.DeploymentSpecDotEnv, deploymentType, envName))
		}
	}

	return result
}

func parseDotEnvFile(content []byte) map[string]string {
	result := make(map[string]string)

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		result[parts[0]] = parts[1]
	}

	return result
}
