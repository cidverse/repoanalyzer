package node

import (
	"testing"

	"github.com/cidverse/repoanalyzer/util"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

func TestAnalyzer_AnalyzeReact(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "react"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
