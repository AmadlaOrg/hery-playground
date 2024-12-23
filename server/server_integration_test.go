package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: The path needs to have "../" +  added. This is there because the path is not the same when testing vs when it is running.
// TODO: Move the setting of the ABS path in the main.go via New Service function call. And then correct hsi part.

func TestTemplateSetup_integration(t *testing.T) {
	s := SServer{}
	err := s.templateSetup("../" + TmplPath)
	assert.NoError(t, err)
}

// TestParseTemplates_integration is to verify that there is no errors when all the pieces come together
func TestParseTemplates_integration(t *testing.T) {
	absPath, err := filepathAbs("../" + TmplPath)
	if err != nil {
		t.Fatalf("")
	}

	s := SServer{}
	_, err = s.parseTemplates(absPath)
	assert.NoError(t, err)
}
