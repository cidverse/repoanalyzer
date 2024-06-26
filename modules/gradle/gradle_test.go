package gradle

import (
	"os"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestGradleAnalyzer_AnalyzeGroovy(t *testing.T) {
	_ = os.Setenv("REPOANAYLZER_DEBUG", "true")
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "gradle-groovy"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "gradle-groovy", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemGradle, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxGradleGroovyDSL, result[0].BuildSystemSyntax)
	assert.Equal(t, "17.0.0", result[0].Language[analyzerapi.LanguageJava])
	assert.Len(t, result[0].Dependencies, 1)
	assert.Equal(t, "maven", result[0].Dependencies[0].Type)
	assert.Equal(t, "junit:junit", result[0].Dependencies[0].ID)
	assert.Equal(t, "4.12", result[0].Dependencies[0].Version)

	// submodule
	assert.Len(t, result[0].Submodules, 1)
	assert.Equal(t, "gradle-groovy-api", result[0].Submodules[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemGradle, result[0].Submodules[0].BuildSystem)
	assert.Equal(t, string(analyzerapi.BuildSystemSyntaxGradleGroovyDSL), string(result[0].Submodules[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestGradleAnalyzer_AnalyzeKotlin(t *testing.T) {
	_ = os.Setenv("REPOANAYLZER_DEBUG", "true")
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "gradle-kotlin"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "gradle-kotlin", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemGradle, result[0].BuildSystem)
	assert.Equal(t, string(analyzerapi.BuildSystemSyntaxGradleKotlinDSL), string(result[0].BuildSystemSyntax))
	assert.Equal(t, "8.0.0", result[0].Language[analyzerapi.LanguageJava])
	assert.Len(t, result[0].Dependencies, 1)

	// submodule
	assert.Len(t, result[0].Submodules, 1)
	assert.Equal(t, "gradle-kotlin-api", result[0].Submodules[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemGradle, result[0].Submodules[0].BuildSystem)
	assert.Equal(t, string(analyzerapi.BuildSystemSyntaxGradleKotlinDSL), string(result[0].Submodules[0].BuildSystemSyntax))

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
