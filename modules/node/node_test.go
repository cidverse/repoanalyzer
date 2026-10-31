package node

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyzer_AnalyzeReact(t *testing.T) {
	ctx := analyzerapi.GetAnalyzerContext(analyzerapi.GetTestDataDir(t, "react"))
	analyzer := Analyzer{}
	result := analyzer.Scan(ctx)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}

func TestAnalyzer_SkipsUnnamedPackages(t *testing.T) {
	dir := t.TempDir()
	app := writeNodeTestFile(t, dir, "package.json", `{"name":"my-app"}`)
	unnamed := writeNodeTestFile(t, dir, "tools/package.json", `{"dependencies":{}}`)
	blankName := writeNodeTestFile(t, dir, "blank/package.json", `{"name":"   "}`)

	files := []string{app, unnamed, blankName}
	ctx := analyzerapi.AnalyzerContext{
		ProjectDir:       dir,
		Files:            files,
		FilesByExtension: map[string][]string{"json": files},
	}

	result := Analyzer{}.Scan(ctx)

	require.Len(t, result, 1)
	assert.Equal(t, "my-app", result[0].Name)
	assert.Equal(t, "my-app", result[0].Slug)
}

func writeNodeTestFile(t *testing.T, root, relativePath, content string) string {
	t.Helper()

	path := filepath.Join(root, filepath.FromSlash(relativePath))
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0755))
	require.NoError(t, os.WriteFile(path, []byte(content), 0644))
	return path
}
