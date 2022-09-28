package node

import (
	"github.com/cidverse/repoanalyzer/util"
	"github.com/rs/zerolog/log"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

func TestAnalyzer_AnalyzeReact(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "react"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// print result
	log.Info().Interface("result", result).Msg("output")
}
