package cargo

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeCargo(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "cargo"))

	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "cargo", result[0].Name)
	assert.Equal(t, "cargo", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
