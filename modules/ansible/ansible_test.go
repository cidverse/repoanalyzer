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
	assert.Equal(t, "ansible-playbook-a-playbook-yml", result[0].ID)
	assert.Equal(t, "playbook-a", result[0].Name)
	assert.Equal(t, analyzerapi.ModuleTypeBuildSystem, result[0].Type)
	assert.Equal(t, analyzerapi.BuildSystemAnsible, result[0].BuildSystem)
	assert.Equal(t, analyzerapi.BuildSystemSyntaxDefault, result[0].BuildSystemSyntax)

	assert.Equal(t, []analyzerapi.ProjectRepository{
		{Type: analyzerapi.RepositoryTypeGalaxy, Name: "galaxy.ansible.com", URL: "https://galaxy.ansible.com"},
		{Type: analyzerapi.RepositoryTypeGit, Name: "org.role", URL: "https://github.com/org/role.git"},
		{Type: analyzerapi.RepositoryTypeGit, Name: "org.other", URL: "ssh://git@github.com/org/other-role.git"},
		{Type: analyzerapi.RepositoryTypeGit, Name: "https://github.com/org/coll.git", URL: "https://github.com/org/coll.git"},
		{Type: analyzerapi.RepositoryTypeGalaxy, Name: "https://galaxy.example.com", URL: "https://galaxy.example.com"},
	}, result[0].Repositories)

	// submodule
	assert.Len(t, result[0].Submodules, 0)

	// print result
	for i, item := range result {
		t.Logf("result[%d]: %+v", i, *item)
	}
}
