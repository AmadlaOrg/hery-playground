package template

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInitialize_integration is to verify that there are no errors when all the
// pieces come together: a real engine, the real template directory, and the
// template set being attached to the engine.
func TestInitialize_integration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	templateService := NewTemplateService(gin.New(), TestTmplPath)

	// Initialize wraps parseTemplates in templateMust, so a parse failure
	// surfaces as a panic rather than an error return.
	assert.NotPanics(t, func() {
		templateService.Initialize()
	})
}

// TestParseTemplates_integration is to verify that there is no errors when all the pieces come together
func TestParseTemplates_integration(t *testing.T) {
	absPath, err := filepathAbs(TestTmplPath)
	if err != nil {
		t.Fatalf("could not resolve %q: %v", TestTmplPath, err)
	}

	templateService := &STemplate{}
	got, err := templateService.parseTemplates(absPath)
	assert.NoError(t, err)
	assert.NotNil(t, got)
}
