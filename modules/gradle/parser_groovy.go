package gradle

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// ParseBuildGradleGroovy parses the build.gradle file and extracts the Java version and dependencies
func ParseBuildGradleGroovy(file string) (BuildGradle, error) {
	// read build.gradle
	buildGradle, err := os.ReadFile(file)
	if err != nil {
		return BuildGradle{}, fmt.Errorf("failed to open build.gradle: %v", err)
	}
	buildGradleStr := string(buildGradle)

	// languages
	languages := make(map[analyzerapi.ProjectLanguage]string)
	pluginBlock := extractPluginBlock(buildGradleStr)
	if pluginBlock != "" {
		for _, plugin := range kotlinPlugins {
			if strings.Contains(pluginBlock, fmt.Sprintf("id 'org.jetbrains.kotlin.%s'", plugin)) {
				languages[analyzerapi.LanguageKotlin] = ""
			}
		}
		if strings.Contains(pluginBlock, "id 'java'") || strings.Contains(pluginBlock, "id 'java-library'") {
			languages[analyzerapi.LanguageJava] = parseJavaVersionOrDefault(buildGradleStr, "21.0.0")
		}
	}

	// dependencies
	depRegex := regexp.MustCompile(`(implementation|compile) ['"](.*:.*):(.*)['"]`)
	depMatches := depRegex.FindAllStringSubmatch(buildGradleStr, -1)
	dependencies := make([]analyzerapi.ProjectDependency, 0, len(depMatches))
	for _, depMatch := range depMatches {
		if len(depMatch) > 2 {
			dependencies = append(dependencies, analyzerapi.ProjectDependency{Type: "maven", ID: depMatch[2], Version: depMatch[3]})
		}
	}

	return BuildGradle{
		Languages:    languages,
		Dependencies: dependencies,
	}, nil
}
