package analyzerapi

import (
	"os"
	"path/filepath"
	"strings"

	ignore "github.com/sabhiram/go-gitignore"
)

// defaultIgnoredDirectories are always excluded from the scan, regardless of
// whether they are mentioned in a .gitignore or .cidignore file.
var defaultIgnoredDirectories = map[string]struct{}{
	".git":         {},
	".svn":         {},
	".hg":          {},
	"node_modules": {},
	"vendor":       {},
}

type projectIgnoreRule struct {
	matcher *ignore.GitIgnore
	negate  bool
}

type projectIgnoreRules struct {
	rules []projectIgnoreRule
}

func (r *projectIgnoreRules) match(path string) (ignored bool, matched bool) {
	for _, rule := range r.rules {
		if rule.matcher.MatchesPath(path) {
			ignored = !rule.negate
			matched = true
		}
	}

	return ignored, matched
}

// ProjectIgnoreMatcher checks paths against the project's ignore rules.
type ProjectIgnoreMatcher struct {
	projectDir  string
	directories map[string]*projectIgnoreRules
}

// ProcessIgnoreFiles returns a matcher for the project's default ignores and
// all .gitignore/.cidignore files, including nested ones.
func ProcessIgnoreFiles(projectDir string) *ProjectIgnoreMatcher {
	return &ProjectIgnoreMatcher{
		projectDir:  projectDir,
		directories: make(map[string]*projectIgnoreRules),
	}
}

// MatchesPath checks a path against the default ignored directories and all
// .gitignore/.cidignore files from the project root up to the path's parent directory.
func (m *ProjectIgnoreMatcher) MatchesPath(absPath string, isDir bool) bool {
	relativePath, err := filepath.Rel(m.projectDir, absPath)
	if err != nil {
		return false
	}

	relativePath = filepath.ToSlash(relativePath)
	if relativePath == "." || relativePath == ".." || strings.HasPrefix(relativePath, "../") {
		return false
	}

	parts := strings.Split(relativePath, "/")
	for i, part := range parts {
		if _, ok := defaultIgnoredDirectories[part]; ok && (isDir || i < len(parts)-1) {
			return true
		}
	}

	ignored := false
	dir := m.projectDir
	for depth := 0; depth < len(parts); depth++ {
		if depth > 0 {
			dir = filepath.Join(dir, parts[depth-1])
		}

		rules := m.rulesForDirectory(dir)
		if rules == nil {
			continue
		}

		pathRelativeToDir := strings.Join(parts[depth:], "/")
		if isDir {
			// directory-only patterns need the trailing slash to match the directory itself
			pathRelativeToDir += "/"
		}
		if dirIgnored, matched := rules.match(pathRelativeToDir); matched {
			ignored = dirIgnored
		}
	}

	return ignored
}

func (m *ProjectIgnoreMatcher) rulesForDirectory(dir string) *projectIgnoreRules {
	if rules, ok := m.directories[dir]; ok {
		return rules
	}

	var lines []string
	for _, name := range []string{".gitignore", ".cidignore"} {
		content, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			continue
		}
		lines = append(lines, strings.Split(string(content), "\n")...)
	}

	rules := compileIgnoreRules(lines)
	m.directories[dir] = rules
	return rules
}

func compileIgnoreRules(lines []string) *projectIgnoreRules {
	var rules []projectIgnoreRule
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		negate := strings.HasPrefix(trimmed, "!") && !strings.HasPrefix(trimmed, `\!`)
		pattern := line
		if negate {
			pattern = trimmed[1:]
		}

		matcher := ignore.CompileIgnoreLines(pattern)
		rules = append(rules, projectIgnoreRule{matcher: matcher, negate: negate})
	}

	if len(rules) == 0 {
		return nil
	}

	return &projectIgnoreRules{rules: rules}
}
