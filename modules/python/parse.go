package python

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type PyProject struct {
	Project ProjectSection `toml:"project"`
	Tool    ToolSection    `toml:"tool"`
}

type ProjectSection struct {
	Name           string   `toml:"name"`
	Version        string   `toml:"version"`
	Description    string   `toml:"description"`
	Readme         string   `toml:"readme"`
	Authors        []string `toml:"authors"`
	RequiresPython string   `toml:"requires-python"`
}

type ToolSection struct {
	Poetry map[string]interface{} `toml:"poetry"`
	UV     map[string]interface{} `toml:"uv"`
}

func readPyProjectFile(path string) (PyProject, error) {
	var data PyProject

	content, err := os.ReadFile(path)
	if err != nil {
		return PyProject{}, fmt.Errorf("failed to read file: %w", err)
	}

	if _, err = toml.Decode(string(content), &data); err != nil {
		return PyProject{}, fmt.Errorf("failed to parse TOML: %w", err)
	}

	return data, nil
}
