package gradle

import (
	"fmt"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// BuildGradle represents the contents of a build.gradle file
type BuildGradle struct {
	JavaVersion  string
	Dependencies []analyzerapi.ProjectDependency
}

// getSemverJavaVersion converts a Java version constant to its corresponding semantic version string.
func getSemverJavaVersion(input string) string {
	// Remove any "\r" characters
	input = strings.Replace(input, "\r", "", -1)

	// Pre-Java 9 constants have a different naming convention
	for i := 1; i <= 8; i++ {
		input = strings.Replace(input, fmt.Sprintf("1.%d", i), fmt.Sprintf("%d.0.0", i), 1)
		input = strings.Replace(input, fmt.Sprintf("JavaVersion.VERSION_1_%d", i), fmt.Sprintf("%d.0.0", i), 1)
	}

	// Java 9 and later constants have a different naming convention
	for i := 9; i <= 40; i++ {
		input = strings.Replace(input, fmt.Sprintf("%d", i), fmt.Sprintf("%d.0.0", i), 1)
		input = strings.Replace(input, fmt.Sprintf("JavaVersion.VERSION_%d", i), fmt.Sprintf("%d.0.0", i), 1)
	}

	return input
}
