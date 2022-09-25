package helm

import (
	"github.com/cidverse/repoanalyzer/logger"
	"github.com/cidverse/repoanalyzer/util"
	"github.com/go-logr/logr/testr"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeHugo(t *testing.T) {
	logger.Logger = testr.New(t)

	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "helm"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "mychart", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemHelm, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDefault, result[0].BuildSystemSyntax)
	assert.Nil(t, result[0].Language)

	// print result
	logger.Info("output", "result", result)
}
