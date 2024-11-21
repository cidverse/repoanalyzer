package dotnet

import (
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_AnalyzeVisualStudioSolution(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "dotnet"))

	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "dotnet", result[0].Name)
	assert.Equal(t, analyzerapi.ModuleTypeBuildSystem, result[0].Type)
	assert.Equal(t, analyzerapi.BuildSystemDotNet, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDotNetCSProj, result[0].BuildSystemSyntax)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_AnalyzeVisualStudioSolutionNested(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "dotnet-nested"))

	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Equal(t, "dotnet-nested", result[0].Name)
	assert.Equal(t, analyzerapi.ModuleTypeBuildSystem, result[0].Type)
	assert.Equal(t, analyzerapi.BuildSystemDotNet, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDotNetSLN, result[0].BuildSystemSyntax)
	assert.Len(t, result[0].Submodules, 1)
	assert.Equal(t, "dotnet-nested-hello-world", result[0].Submodules[0].Name)
	assert.Equal(t, analyzerapi.ModuleTypeBuildSystem, result[0].Submodules[0].Type)
	assert.Equal(t, analyzerapi.BuildSystemDotNet, result[0].Submodules[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDotNetCSProj, result[0].Submodules[0].BuildSystemSyntax)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
