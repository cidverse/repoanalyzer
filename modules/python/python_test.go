package python

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

func TestAnalyzer_AnalyzePythonPipfile(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "python-pipfile"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

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

func TestAnalyzer_AnalyzePythonRequirementsTXT(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "python-requirementstxt"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

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
	result := analyzer.Scan(ctx)

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

func TestAnalyzer_AnalyzePythonPoetry(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "python-poetry"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "test-library-python-poetry", result[0].Name)
	assert.Equal(t, "pyproject-poetry", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))
	assert.Len(t, result[0].Dependencies, 3)
	assert.Equal(t, "pytest", result[0].Dependencies[0].ID)
	assert.Equal(t, "8.3.5", result[0].Dependencies[0].Version)
	assert.Equal(t, "dev", result[0].Dependencies[0].Scope)
	assert.Equal(t, "pytest-cov", result[0].Dependencies[1].ID)
	assert.Equal(t, "6.1.1", result[0].Dependencies[1].Version)
	assert.Equal(t, "dev", result[0].Dependencies[1].Scope)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_AnalyzePythonUV(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "python-uv"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "test-library-python-uv", result[0].Name)
	assert.Equal(t, "pyproject-uv", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))
	assert.Len(t, result[0].Dependencies, 2)
	assert.Equal(t, "pytest", result[0].Dependencies[0].ID)
	assert.Equal(t, "8.3.5", result[0].Dependencies[0].Version)
	assert.Equal(t, "dev", result[0].Dependencies[0].Scope)
	assert.Equal(t, "pytest-cov", result[0].Dependencies[1].ID)
	assert.Equal(t, "6.1.1", result[0].Dependencies[1].Version)
	assert.Equal(t, "dev", result[0].Dependencies[1].Scope)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
