package python

import (
	"testing"

	"github.com/cidverse/repoanalyzer/util"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

func TestAnalyzer_AnalyzePythonPipfile(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "python-pipfile"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "python-pipfile", result[0].Name)
	assert.Equal(t, "pipfile", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	log.Info().Interface("result", result).Msg("output")
}

func TestAnalyzer_AnalyzePythonPeotry(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "python-poetry"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "python-poetry", result[0].Name)
	assert.Equal(t, "poetry", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	log.Info().Interface("result", result).Msg("output")
}

func TestAnalyzer_AnalyzePythonRequirementsTXT(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "python-requirementstxt"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "python-requirementstxt", result[0].Name)
	assert.Equal(t, "python-requirements.txt", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	log.Info().Interface("result", result).Msg("output")
}

func TestAnalyzer_AnalyzePythonSetuppy(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "python-setuppy"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "python-setuppy", result[0].Name)
	assert.Equal(t, "setup.py", string(result[0].BuildSystem))
	assert.Equal(t, "default", string(result[0].BuildSystemSyntax))

	// print result
	log.Info().Interface("result", result).Msg("output")
}
