package gitlabci

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeGitHubWorkflows(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "gitlabci"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "gitlab-ci", result[0].Name)
	assert.Equal(t, analyzerapi.ModuleTypeConfig, result[0].Type)
	assert.Equal(t, analyzerapi.ConfigTypeGitLabCI, result[0].ConfigType)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
