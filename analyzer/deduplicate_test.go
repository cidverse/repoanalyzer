package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cidverse/repoanalyzer/analyzerapi"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_DeduplicateRspressOverNode(t *testing.T) {
	// copy fixture to a temp dir, the analyzer ignores directories containing "testdata"
	fixtureDir := filepath.Join("..", "testdata", "rspress-node")
	tmpDir := t.TempDir()
	assert.NoError(t, copyDir(fixtureDir, tmpDir))

	rs := NewAnalyzer()
	out := rs.Scan(tmpDir)

	// only the rspress module should remain, the npm module is deduplicated
	var buildSystems []string
	for _, module := range out {
		buildSystems = append(buildSystems, string(module.BuildSystem))
	}

	assert.Contains(t, buildSystems, string(analyzerapi.BuildSystemRspress))
	assert.NotContains(t, buildSystems, string(analyzerapi.BuildSystemNpm))
}

func copyDir(src string, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		content, err := os.ReadFile(srcPath)
		if err != nil {
			return err
		}
		if err := os.WriteFile(dstPath, content, 0o644); err != nil {
			return err
		}
	}
	return nil
}
