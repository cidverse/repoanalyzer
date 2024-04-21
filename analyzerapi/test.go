package analyzerapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetTestDataDir(t *testing.T, dir string) string {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	return filepath.Join(filepath.Dir(cwd), "..", "testdata", dir)
}
