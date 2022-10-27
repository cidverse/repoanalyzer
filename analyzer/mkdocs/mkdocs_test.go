package mkdocs

import (
	"github.com/cidverse/repoanalyzer/util"
	"github.com/rs/zerolog/log"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeMkdocs(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "mkdocs"))

	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "mkdocs", result[0].Name)
	assert.Equal(t, "mkdocs", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	log.Info().Interface("result", result).Msg("output")
}

func TestAnalyzer_AnalyzeTechdocs(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "techdocs"))

	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "techdocs", result[0].Name)
	assert.Equal(t, "mkdocs", string(result[0].BuildSystem))
	assert.Equal(t, "mkdocs-techdocs", string(result[0].BuildSystemSyntax))

	// print result
	log.Info().Interface("result", result).Msg("output")
}
