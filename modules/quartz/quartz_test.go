package quartz

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeQuartz(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "quartz"))

	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "quartz", result[0].Name)
	assert.Equal(t, "quartz", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
