package gradle

import (
	"fmt"
	"os"
	"regexp"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// ParseBuildGradleGroovy parses the build.gradle file and extracts the Java version and dependencies
func ParseBuildGradleGroovy(file string) (BuildGradle, error) {
	// Read the contents of the build.gradle file
	buildGradle, err := os.ReadFile(file)
	if err != nil {
		return BuildGradle{}, fmt.Errorf("failed to open build.gradle: %v", err)
	}

	// Convert the contents to a string
	buildGradleStr := string(buildGradle)

	// Extract the Java version
	javaRegex := regexp.MustCompile(`sourceCompatibility\s?=\s?['"]?(.*)['"]?`)
	javaMatches := javaRegex.FindStringSubmatch(buildGradleStr)
	if len(javaMatches) < 2 {
		return BuildGradle{}, fmt.Errorf("failed to extract Java version from build.gradle")
	}
	javaVersion := getSemverJavaVersion(javaMatches[1])

	// Extract the dependencies
	depRegex := regexp.MustCompile(`(implementation|compile) ['"](.*:.*):(.*)['"]`)
	depMatches := depRegex.FindAllStringSubmatch(buildGradleStr, -1)
	dependencies := make([]analyzerapi.ProjectDependency, 0, len(depMatches))
	for _, depMatch := range depMatches {
		if len(depMatch) > 2 {
			dependencies = append(dependencies, analyzerapi.ProjectDependency{Type: "maven", ID: depMatch[2], Version: depMatch[3]})
		}
	}

	// Return the BuildGradle struct
	return BuildGradle{
		JavaVersion:  javaVersion,
		Dependencies: dependencies,
	}, nil
}
