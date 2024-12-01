package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewControllerService(t *testing.T) {
	t.Run("should return a new instance of controller", func(t *testing.T) {
		service := NewControllerService()
		assert.NotNil(t, service)
		assert.IsType(t, &SController{}, service)
	})
}
