package helmfile

import (
	"testing"
)

func TestParseHelmfileEnvironmentsContent_Simple(t *testing.T) {
	helmfileContent := `
environments:
  dev:
    values:
      - dev-values.yaml
  prod:
    values:
      - prod-values.yaml
`
	envs, err := parseHelmfileEnvironmentsContent(helmfileContent)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if _, exists := envs["dev"]; !exists {
		t.Errorf("Expected 'dev' environment to exist")
	}
	if _, exists := envs["prod"]; !exists {
		t.Errorf("Expected 'prod' environment to exist")
	}
}

func TestParseHelmfileEnvironmentsContent_GoTemplate(t *testing.T) {
	helmfileContent := `
environments:
  dev:
    values:
      - dev-values.yaml
---
releases:
  - name: {{ .Environment.Name }}-nginx
    namespace: {{ .Namespace }}
`
	envs, err := parseHelmfileEnvironmentsContent(helmfileContent)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if _, exists := envs["dev"]; !exists {
		t.Errorf("Expected 'dev' environment to exist")
	}
}

func TestParseHelmfileEnvironmentsContent_NoEnvironments(t *testing.T) {
	helmfileContent := `
releases:
  - name: my-app
    chart: my-chart
`
	_, err := parseHelmfileEnvironmentsContent(helmfileContent)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}
