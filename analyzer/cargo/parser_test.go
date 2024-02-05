package cargo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCargoFile(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected Config
	}{
		{
			name: "Valid Cargo file",
			input: `
				[package]
				name = "my-project"
				version = "0.1.0"
				authors = ["Firstname Lastname <firstname.lastname@example.com>"]
				edition = "2021"
				license = "MIT"
				description = """
				Multiline
				Description
				"""
				rust-version = "1.56"
			`,
			expected: Config{
				Package: Package{
					Name:        "my-project",
					RustVersion: "1.56",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := parseCargoFileFromByteArray([]byte(tc.input))

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, cfg)
		})
	}
}
