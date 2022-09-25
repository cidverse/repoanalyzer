package node

import (
	"github.com/cidverse/repoanalyzer/logger"
	"github.com/cidverse/repoanalyzer/util"
	"github.com/go-logr/logr/testr"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

func TestAnalyzer_AnalyzeReact(t *testing.T) {
	logger.Logger = testr.New(t)

	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "react"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// print result
	logger.Info("output", "result", result)
}
