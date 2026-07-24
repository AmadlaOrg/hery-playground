package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestTmplPath is the template root relative to this package's directory.
// server.TmplPath cannot be reused here: server imports template, so importing
// it back would create a cycle.
const TestTmplPath = "../../browser/frontend/template"

func TestParseTemplates(t *testing.T) {
	tests := []struct {
		name              string
		inputTmplRootPath string
		expectNames       []string
		hasError          bool
	}{
		{
			name:              "empty path returns a walk error",
			inputTmplRootPath: "",
			hasError:          true,
		},
		{
			name:              "missing directory returns a walk error",
			inputTmplRootPath: "./does-not-exist",
			hasError:          true,
		},
		{
			name:              "template directory parses every html file",
			inputTmplRootPath: TestTmplPath,
			expectNames: []string{
				"404.html",
				"500.html",
				"dropdown.html",
				"header.html",
				"horizontal.html",
				"index.html",
				"main.html",
				"modal.html",
				"script.html",
				"style.html",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			templateService := &STemplate{}
			got, err := templateService.parseTemplates(tt.inputTmplRootPath)
			if tt.hasError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, got)

			// parseTemplates seeds the set with templateNew(""), so the root
			// template carries an empty name alongside the parsed files.
			var gotNames []string
			for _, parsed := range got.Templates() {
				if parsed.Name() == "" {
					continue
				}
				gotNames = append(gotNames, parsed.Name())
			}
			assert.ElementsMatch(t, tt.expectNames, gotNames)
		})
	}
}
