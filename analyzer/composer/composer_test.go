package composer

import (
	"testing"

	"github.com/cidverse/repoanalyzer/util"
	"github.com/rs/zerolog/log"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeComposer(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "composer"))

	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "composer", result[0].Name)
	assert.Equal(t, "composer", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))
	assert.Equal(t, "8.0.0", result[0].Language[analyzerapi.LanguagePHP])

	// print result
	log.Info().Interface("result", result).Msg("output")
}
