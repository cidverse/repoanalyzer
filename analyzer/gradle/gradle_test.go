package gradle

import (
	"github.com/cidverse/repoanalyzer/logger"
	"github.com/cidverse/repoanalyzer/util"
	"github.com/go-logr/logr/testr"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestGradleAnalyzer_AnalyzeGroovy(t *testing.T) {
	logger.Logger = testr.New(t)

	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "gradle-groovy"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "gradle-groovy", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemGradle, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.GradleGroovyDSL, result[0].BuildSystemSyntax)
	assert.Nil(t, result[0].Language[analyzerapi.LanguageJava])

	// submodule
	assert.Len(t, result[0].Submodules, 1)
	assert.Equal(t, "gradle-groovy-api", result[0].Submodules[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemGradle, result[0].Submodules[0].BuildSystem)
	assert.Equal(t, string(analyzerapi.GradleGroovyDSL), string(result[0].Submodules[0].BuildSystemSyntax))
	assert.Nil(t, result[0].Submodules[0].Language[analyzerapi.LanguageJava])

	// print result
	logger.Logger.Info("output", "result", result)
}

func TestGradleAnalyzer_AnalyzeKotlin(t *testing.T) {
	logger.Logger = testr.New(t)

	ctx := analyzerapi.GetAnalyzerContext(util.GetTestDataDir(t, "gradle-kotlin"))
	analyzer := Analyzer{}
	result := analyzer.Analyze(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "gradle-kotlin", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemGradle, result[0].BuildSystem)
	assert.Equal(t, string(analyzerapi.GradleKotlinDSL), string(result[0].BuildSystemSyntax))
	assert.Nil(t, result[0].Language[analyzerapi.LanguageJava])

	// submodule
	assert.Len(t, result[0].Submodules, 1)
	assert.Equal(t, "gradle-kotlin-api", result[0].Submodules[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemGradle, result[0].Submodules[0].BuildSystem)
	assert.Equal(t, string(analyzerapi.GradleKotlinDSL), string(result[0].Submodules[0].BuildSystemSyntax))
	assert.Nil(t, result[0].Submodules[0].Language[analyzerapi.LanguageJava])

	// print result
	logger.Logger.Info("output", "result", result)
}
