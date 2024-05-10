package maven

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeMaven(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "maven"))

	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "maven", result[0].Name)
	assert.Equal(t, "maven", string(result[0].BuildSystem))
	assert.Equal(t, string(analyzerapi.BuildSystemSyntaxDefault), string(result[0].BuildSystemSyntax))
	assert.Equal(t, "1.8", result[0].Language[analyzerapi.LanguageJava])
	assert.Len(t, result[0].Dependencies, 1)
	assert.Equal(t, "maven", result[0].Dependencies[0].Type)
	assert.Equal(t, "junit:junit", result[0].Dependencies[0].ID)
	assert.Equal(t, "4.12", result[0].Dependencies[0].Version)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
