package gradle

import (
	"testing"
)

func TestJavaVersionToSemver(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "java 8",
			input:  "JavaVersion.VERSION_1_8",
			output: "8.0.0",
		},
		{
			name:   "java 11",
			input:  "JavaVersion.VERSION_11",
			output: "11.0.0",
		},
		{
			name:   "java 15",
			input:  "JavaVersion.VERSION_15",
			output: "15.0.0",
		},
		{
			name:   "java 8 with carriage return",
			input:  "JavaVersion.VERSION_1_8\r",
			output: "8.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := javaVersionToSemver(tt.input)

			if result != tt.output {
				t.Errorf("Expected output: %s, Actual output: %s", tt.output, result)
			}
		})
	}
}

func TestExtractPluginBlock(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name: "Kotlin-DSL plugin block",
			input: `
				plugins {
					java
					id("io.freefair.lombok").version("8.4").apply(false)
				}
				// other content
			`,
			output: `
					java
					id("io.freefair.lombok").version("8.4").apply(false)
				`,
		},
		{
			name: "Groovy plugin block",
			input: `
				plugins {
					id 'java'
					id 'org.jetbrains.kotlin.jvm' version '1.5.21'
				}
				// other content
			`,
			output: `
					id 'java'
					id 'org.jetbrains.kotlin.jvm' version '1.5.21'
				`,
		},
		{
			name: "no plugin block",
			input: `
				// no plugins block
				// other content
			`,
			output: "",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractPluginBlock(tt.input)

			if result != tt.output {
				t.Errorf("Expected output:\n%s\n\nActual output:\n%s", tt.output, result)
			}
		})
	}
}
