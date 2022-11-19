package maven

import (
	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/cidverse/repoanalyzer/util"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyzer_AnalyzeMaven(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "maven"))

	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "maven", result[0].Name)
	assert.Equal(t, "maven", string(result[0].BuildSystem))
	assert.Equal(t, string(analyzerapi.BuildSystemSyntaxDefault), string(result[0].BuildSystemSyntax))

	// print result
	log.Info().Interface("result", result).Msg("output")
}
