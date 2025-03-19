package deploymentdotenv

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeDeplyomentDotEnv(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "deploymentdotenv"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "deployment-dotenv-dev", result[0].Name)
	assert.Equal(t, "deployment-dotenv-dev", result[0].Slug)
	assert.Equal(t, analyzerapi.ModuleTypeDeployment, result[0].Type)
	assert.Equal(t, analyzerapi.DeploymentSpecDotEnv, result[0].DeploymentSpec)
	assert.Equal(t, "helm", result[0].DeploymentType)
	assert.Equal(t, "dev", result[0].DeploymentEnvironment)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
