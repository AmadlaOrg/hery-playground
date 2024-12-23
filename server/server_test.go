package server

import (
	"github.com/stretchr/testify/assert"
	"html/template"
	"path/filepath"
	"testing"
)

func TestParseTemplates(t *testing.T) {
	tests := []struct {
		name                 string
		inputTmplRootPath    string
		internalTemplateNew  func(name string) *template.Template
		internalFilepathWalk func(root string, fn filepath.WalkFunc) error
		internalFilepathExt  func(path string) string
		expectTmpl           *template.Template
		expertErr            error
		hasError             bool
	}{
		{
			name: "empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serverService := &SServer{}
			got, err := serverService.parseTemplates(tt.inputTmplRootPath)
			if tt.hasError {
				assert.Error(t, err)
				assert.ErrorContains(t, tt.expertErr, err.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectTmpl, got)
		})
	}
}
