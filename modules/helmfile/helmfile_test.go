package helmfile

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeHelmfile(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "helmfile"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "helmfile", result[0].Name)
	assert.Equal(t, analyzerapi.DeploymentSpecHelmfile, result[0].DeploymentSpec)
	assert.Equal(t, "helmfile", result[0].DeploymentType)
	assert.Nil(t, result[0].Language)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_AnalyzeHelmfileChart(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "helmfile-chart"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "example", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemHelmfile, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDefault, result[0].BuildSystemSyntax)
	assert.Nil(t, result[0].Language)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
