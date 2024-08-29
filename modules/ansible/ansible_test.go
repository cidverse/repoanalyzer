package ansible

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnsibleAnalyzer_Analyze(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	ctx := analyzerapi.GetAnalyzerContext(filepath.Join(filepath.Dir(cwd), "..", "testdata", "ansible"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// module
	assert.Len(t, result, 1)
	assert.Len(t, result[0].Discovery, 1)
	assert.Equal(t, "playbook-a", result[0].Name)
	assert.Equal(t, analyzerapi.BuildSystemAnsible, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDefault, result[0].BuildSystemSyntax)

	// submodule
	assert.Len(t, result[0].Submodules, 0)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
