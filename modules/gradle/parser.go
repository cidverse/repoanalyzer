package gradle

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// BuildGradle represents the contents of a build.gradle file
type BuildGradle struct {
	Languages    map[analyzerapi.ProjectLanguage]string
	Dependencies []analyzerapi.ProjectDependency
}

var kotlinPlugins = []string{"jvm", "multiplatform"}

// javaVersionToSemver converts a Java version constant to its corresponding semantic version string
func javaVersionToSemver(input string) string {
	// remove any "\r" characters
	input = strings.Replace(input, "\r", "", -1)

	// pre-Java 9 constants have a different naming convention
	for i := 1; i <= 8; i++ {
		if input == fmt.Sprintf("JavaVersion.VERSION_1_%d", i) {
			return fmt.Sprintf("%d.0.0", i)
		}
		if input == fmt.Sprintf("1.%d", i) {
			return fmt.Sprintf("%d.0.0", i)
		}
	}

	// java 9 and later constants have a different naming convention
	for i := 9; i <= 40; i++ {
		if input == fmt.Sprintf("JavaVersion.VERSION_%d", i) {
			return fmt.Sprintf("%d.0.0", i)
		}
		if input == fmt.Sprintf("%d", i) {
			return fmt.Sprintf("%d.0.0", i)
		}
	}

	return ""
}

func parseJavaVersionOrDefault(input string, defaultValue string) string {
	javaRegex := regexp.MustCompile(`sourceCompatibility\s?=\s?(['"])?(.*)(['"])?`)
	javaMatches := javaRegex.FindStringSubmatch(input)
	if len(javaMatches) < 3 {
		return defaultValue
	}

	javaVersion := javaVersionToSemver(javaMatches[2])
	if javaVersion == "" {
		return defaultValue
	}

	return javaVersion
}

func extractPluginBlock(input string) string {
	startStr := "plugins {"
	endStr := "}"

	// Extract the plugin block
	pluginBlockStart := strings.Index(input, startStr)
	pluginBlockEnd := strings.Index(input, endStr)
	if pluginBlockStart == -1 || pluginBlockEnd == -1 {
		return ""
	}
	pluginBlock := input[pluginBlockStart+len(startStr) : pluginBlockEnd]
	return pluginBlock
}
