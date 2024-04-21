package hugo

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeHugo(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "hugo"))

	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "hugo", result[0].Name)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
