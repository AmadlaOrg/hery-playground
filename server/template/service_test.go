package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTemplateService(t *testing.T) {
	t.Run("should return a new instance of template", func(t *testing.T) {
		service := NewTemplateService()
		assert.NotNil(t, service)
		assert.IsType(t, &STemplate{}, service)
	})
}
