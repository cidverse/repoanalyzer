package nix

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeNixFlake(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "nix-flake"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "nix-flake", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemNix, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxNixFlake, result[0].BuildSystemSyntax)
	assert.Contains(t, result[0].Language, analyzerapi.LanguageNix)
	assert.Equal(t, "0.0.0", result[0].Language["nix"])

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
