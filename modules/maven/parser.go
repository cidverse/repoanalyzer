package maven

import (
	"encoding/xml"
	"fmt"
	"os"
)

// PomXML represents the structure of a Maven pom.xml file
type PomXML struct {
	XMLName         xml.Name `xml:"project"`
	ArtifactID      string   `xml:"artifactId"`
	GroupID         string   `xml:"groupId"`
	Version         string   `xml:"version"`
	Packaging       string   `xml:"packaging"`
	Name            string   `xml:"name"`
	Description     string   `xml:"description"`
	URL             string   `xml:"url"`
	PropJavaVersion string   `xml:"properties>java.version"`
	Build           struct {
		Plugins []struct {
			GroupID    string `xml:"groupId"`
			ArtifactID string `xml:"artifactId"`
			Version    string `xml:"version"`
			Executions []struct {
				Goals []string `xml:"goals>goal"`
			} `xml:"executions>execution"`
		} `xml:"plugins>plugin"`
	}
	Repositories []struct {
		ID        string `xml:"id"`
		URL       string `xml:"url"`
		Layout    string `xml:"layout"`
		Snapshots struct {
			Enabled bool `xml:"enabled"`
		} `xml:"snapshots"`
		Releases struct {
			Enabled bool `xml:"enabled"`
		} `xml:"releases"`
	} `xml:"repositories>repository"`
	Dependencies []struct {
		GroupID    string `xml:"groupId"`
		ArtifactID string `xml:"artifactId"`
		Version    string `xml:"version"`
	} `xml:"dependencies>dependency"`
}

// ParsePackageJSON will parse the package.json to evaluate its content
func ParsePackageJSON(file string) (PomXML, error) {
	var result PomXML

	// Read the contents of the pom.xml file
	pomXML, err := os.ReadFile(file)
	if err != nil {
		return PomXML{}, fmt.Errorf("failed to open pom.xml: %v", err)
	}

	// Unmarshal the XML into a PomXML struct
	if err := xml.Unmarshal(pomXML, &result); err != nil {
		return PomXML{}, fmt.Errorf("failed to parse pom.xml: %v", err)
	}

	return result, nil
}
