package node

import (
	"encoding/json"
	"fmt"
	"os"
)

type PackageJSON struct {
	Name         string
	Version      string
	Dependencies map[string]string
	Scripts      map[string]string
}

// ParsePackageJSON will parse the package.json to evaluate its content
func ParsePackageJSON(file string) (PackageJSON, error) {
	var result PackageJSON

	// Read the contents of the package.json file
	pkgJSON, err := os.ReadFile(file)
	if err != nil {
		return PackageJSON{}, fmt.Errorf("failed to open package.json: %v", err)
	}

	// Unmarshal the JSON into a PackageStruct struct
	if err := json.Unmarshal(pkgJSON, &result); err != nil {
		return PackageJSON{}, fmt.Errorf("failed to parse package.json: %v", err)
	}

	return result, nil
}
