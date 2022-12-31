package gradle

import (
	"fmt"
	"os"
	"regexp"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// ParseBuildGradleKotlin parses the build.gradle.kts file and extracts the Java version and dependencies
func ParseBuildGradleKotlin(file string) (BuildGradle, error) {
	// Read the contents of the build.gradle.kts file
	buildGradle, err := os.ReadFile(file)
	if err != nil {
		return BuildGradle{}, fmt.Errorf("failed to open build.gradle.kts: %v", err)
	}

	// Convert the contents to a string
	buildGradleStr := string(buildGradle)

	// Extract the Java version
	javaRegex := regexp.MustCompile(`sourceCompatibility\s?=\s?(['"])?(.*)(['"])?`)
	javaMatches := javaRegex.FindStringSubmatch(buildGradleStr)
	if len(javaMatches) < 3 {
		return BuildGradle{}, fmt.Errorf("failed to extract Java version from build.gradle.kts")
	}
	javaVersion := getSemverJavaVersion(javaMatches[2])

	// Extract the dependencies
	depRegex := regexp.MustCompile(`(implementation|compile)\s*\((['"])([^:]+):([^:]+):([^:]+)(['"])\)`)
	depMatches := depRegex.FindAllStringSubmatch(buildGradleStr, -1)
	dependencies := make([]analyzerapi.ProjectDependency, 0, len(depMatches))
	for _, depMatch := range depMatches {
		if len(depMatch) > 5 {
			dependencies = append(dependencies, analyzerapi.ProjectDependency{Type: "maven", ID: depMatch[3] + ":" + depMatch[4], Version: depMatch[5]})
		}
	}

	// Return the BuildGradle struct
	return BuildGradle{
		JavaVersion:  javaVersion,
		Dependencies: dependencies,
	}, nil
}
