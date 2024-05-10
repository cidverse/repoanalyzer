package apispec

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeOpenAPI(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "openapi3"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "openapi3", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemDefault, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDefault, result[0].BuildSystemSyntax)
	assert.Contains(t, result[0].Language, analyzerapi.LanguageOpenAPI)
	assert.Equal(t, "3.0.3", result[0].Language["openapi"])

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_AnalyzeAsyncApi(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "asyncapi"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "asyncapi", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemDefault, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDefault, result[0].BuildSystemSyntax)
	assert.Contains(t, result[0].Language, analyzerapi.LanguageAsyncAPI)
	assert.Equal(t, "3.0.0", result[0].Language["asyncapi"])

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
