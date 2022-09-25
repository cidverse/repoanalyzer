package hugo

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

	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "hugo"))

	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "hugo", result[0].Name)

	// print result
	logger.Info("output", "result", result)
}
