package cargo

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Package Package `toml:"package"`
}

type Package struct {
	Name        string `toml:"name"`
	RustVersion string `toml:"rust-version"`
}

func parseCargoFile(file string) (Config, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open Cargo.toml: %w", err)
	}

	return parseCargoFileFromByteArray(content)
}

func parseCargoFileFromByteArray(content []byte) (Config, error) {
	var cfg Config
	err := toml.Unmarshal(content, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse Cargo.toml: %w", err)
	}

	return cfg, nil
}
