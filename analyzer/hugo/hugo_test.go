package hugo

import (
	"github.com/cidverse/repoanalyzer/util"
	"github.com/rs/zerolog/log"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeHugo(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "hugo"))

	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "hugo", result[0].Name)

	// print result
	log.Info().Interface("result", result).Msg("output")
}
