package gradle

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cidverse/repoanalyzer/analyzerapi"
)

// ParseBuildGradleKotlin parses the build.gradle.kts file and extracts the Java version and dependencies
func ParseBuildGradleKotlin(file string) (BuildGradle, error) {
	// read build.gradle.kts
	buildGradle, err := os.ReadFile(file)
	if err != nil {
		return BuildGradle{}, fmt.Errorf("failed to open build.gradle.kts: %v", err)
	}
	buildGradleStr := string(buildGradle)

	// languages
	languages := make(map[analyzerapi.ProjectLanguage]string)
	pluginBlock := extractPluginBlock(buildGradleStr)
	if pluginBlock != "" {
		for _, plugin := range kotlinPlugins {
			if strings.Contains(pluginBlock, fmt.Sprintf("kotlin(\"%s\")", plugin)) || strings.Contains(pluginBlock, fmt.Sprintf("id(\"org.jetbrains.kotlin.%s\")", plugin)) {
				languages[analyzerapi.LanguageKotlin] = ""
			}
		}
		if strings.Contains(pluginBlock, "java") || strings.Contains(pluginBlock, "java-library") {
			languages[analyzerapi.LanguageJava] = parseJavaVersionOrDefault(buildGradleStr, "21.0.0")
		}
	}

	// dependencies
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
		Languages:    languages,
		Dependencies: dependencies,
	}, nil
}
