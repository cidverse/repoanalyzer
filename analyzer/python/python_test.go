package python

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

func TestAnalyzer_AnalyzePythonPipfile(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "python-pipfile"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "python-pipfile", result[0].Name)
	assert.Equal(t, "pipfile", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_AnalyzePythonPeotry(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "python-poetry"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "python-poetry", result[0].Name)
	assert.Equal(t, "poetry", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_AnalyzePythonRequirementsTXT(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "python-requirementstxt"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "python-requirementstxt", result[0].Name)
	assert.Equal(t, "python-requirements.txt", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_AnalyzePythonSetuppy(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "python-setuppy"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "python-setuppy", result[0].Name)
	assert.Equal(t, "setup.py", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
