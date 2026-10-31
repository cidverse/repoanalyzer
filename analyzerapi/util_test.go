package analyzerapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAnalyzerContext_RespectsNestedGitignore(t *testing.T) {
	dir := t.TempDir()
	writeTestFile(t, dir, "keep.txt", "keep")
	writeTestFile(t, dir, ".opencode/.gitignore", "node_modules\ngenerated\n")
	writeTestFile(t, dir, ".opencode/.cidignore", "cache\n")
	writeTestFile(t, dir, ".opencode/keep.js", "keep")
	writeTestFile(t, dir, ".opencode/generated/skip.js", "ignored")
	writeTestFile(t, dir, ".opencode/cache/skip.js", "ignored")
	writeTestFile(t, dir, ".opencode/node_modules/pkg/package.json", `{"name":"pkg"}`)

	ctx := GetAnalyzerContext(dir)

	assert.Contains(t, ctx.Files, filepath.Join(dir, "keep.txt"))
	assert.Contains(t, ctx.Files, filepath.Join(dir, ".opencode", "keep.js"))
	assertNoPathContaining(t, ctx.Files, "generated")
	assertNoPathContaining(t, ctx.Files, "cache")
	assertNoPathContaining(t, ctx.Files, "node_modules")
}

func TestGetAnalyzerContext_IgnoresNodeModulesByDefault(t *testing.T) {
	dir := t.TempDir()
	writeTestFile(t, dir, "package.json", `{"name":"app"}`)
	writeTestFile(t, dir, "src/app.js", "keep")
	writeTestFile(t, dir, "node_modules/pkg/package.json", `{"name":"pkg"}`)

	ctx := GetAnalyzerContext(dir)

	assert.Contains(t, ctx.Files, filepath.Join(dir, "src", "app.js"))
	assertNoPathContaining(t, ctx.Files, "node_modules")
}

func TestGetAnalyzerContext_DirectoryOnlyIgnoreSkipsNestedRules(t *testing.T) {
	dir := t.TempDir()
	writeTestFile(t, dir, ".gitignore", "build/\n")
	writeTestFile(t, dir, "build/.gitignore", "!keep.txt\n")
	writeTestFile(t, dir, "build/keep.txt", "ignored")

	ctx := GetAnalyzerContext(dir)

	assertNoPathContaining(t, ctx.Files, "build")
}

func TestGetAnalyzerContext_RespectsNestedNegation(t *testing.T) {
	dir := t.TempDir()
	writeTestFile(t, dir, ".gitignore", "*.log\n")
	writeTestFile(t, dir, "sub/.gitignore", "!important.log\n")
	writeTestFile(t, dir, "sub/important.log", "keep")
	writeTestFile(t, dir, "sub/other.log", "ignored")

	ctx := GetAnalyzerContext(dir)

	assert.Contains(t, ctx.Files, filepath.Join(dir, "sub", "important.log"))
	assert.NotContains(t, ctx.Files, filepath.Join(dir, "sub", "other.log"))
}

func writeTestFile(t *testing.T, root, relativePath, content string) {
	t.Helper()

	path := filepath.Join(root, filepath.FromSlash(relativePath))
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0755))
	require.NoError(t, os.WriteFile(path, []byte(content), 0644))
}

func assertNoPathContaining(t *testing.T, files []string, value string) {
	t.Helper()

	for _, file := range files {
		assert.Falsef(t, strings.Contains(filepath.ToSlash(file), value), "unexpected file %s", file)
	}
}
