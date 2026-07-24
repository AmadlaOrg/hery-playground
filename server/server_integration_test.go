package server

import (
	"github.com/AmadlaOrg/hery-playground/server/template"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: The path needs to have "../" +  added. This is there because the path is not the same when testing vs when it is running.
// TODO: Move the setting of the ABS path in the main.go via New Service function call. And then correct hsi part.

// TestTmplPath_integration verifies the TmplPath the server hands to the
// template service points at a directory that actually parses. Template
// parsing itself is covered in the template package; this guards the wiring.
func TestTmplPath_integration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	templateService := template.NewTemplateService(gin.New(), "../"+TmplPath)

	assert.NotPanics(t, func() {
		templateService.Initialize()
	})
}
