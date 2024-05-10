package mkdocs

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeMkdocs(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "mkdocs"))

	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "mkdocs", result[0].Name)
	assert.Equal(t, "mkdocs", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_AnalyzeTechdocs(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "techdocs"))

	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "techdocs", result[0].Name)
	assert.Equal(t, "mkdocs", string(result[0].BuildSystem))
	assert.Equal(t, "mkdocs-techdocs", string(result[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
