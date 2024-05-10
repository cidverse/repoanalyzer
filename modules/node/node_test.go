package node

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

func TestAnalyzer_AnalyzeReact(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "react"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
