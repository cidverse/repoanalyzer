package python

import (
	"fmt"
	"os"
	"regexp"

	"github.com/BurntSushi/toml"
)

type PyProject struct {
	Project          ProjectSection      `toml:"project"`
	Tool             ToolSection         `toml:"tool"`
	DependencyGroups map[string][]string `toml:"dependency-groups"`
}

type ProjectSection struct {
	Name           string           `toml:"name"`
	Version        string           `toml:"version"`
	Description    string           `toml:"description"`
	Readme         string           `toml:"readme"`
	Authors        []AuthorsSection `toml:"authors"`
	RequiresPython string           `toml:"requires-python"`
}

type AuthorsSection struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

type ToolSection struct {
	Poetry PoetrySection `toml:"poetry"`
}

type PoetrySection struct {
	Name           string                        `toml:"name"`
	Version        string                        `toml:"version"`
	Description    string                        `toml:"description"`
	Readme         string                        `toml:"readme"`
	Authors        []string                      `toml:"authors"`
	RequiresPython string                        `toml:"requires-python"`
	Dependencies   map[string]string             `toml:"dependencies"`
	Groups         map[string]PoetryGroupSection `toml:"group"`
}

type PoetryGroupSection struct {
	Dependencies map[string]string `toml:"dependencies"`
}

func readPyProjectFile(path string) (PyProject, error) {
	var data PyProject

	content, err := os.ReadFile(path)
	if err != nil {
		return PyProject{}, fmt.Errorf("failed to read file: %w", err)
	}

	if err = toml.Unmarshal(content, &data); err != nil {
		return PyProject{}, fmt.Errorf("failed to parse TOML: %w", err)
	}

	return data, nil
}

var depSplitRegex = regexp.MustCompile(`([<>=!~]=?)`)

func splitDependency(dep string) (name, version string) {
	loc := depSplitRegex.FindStringIndex(dep)
	if loc == nil {
		return dep, ""
	}

	return dep[:loc[0]], dep[loc[0]:]
}
