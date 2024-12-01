package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewModelService(t *testing.T) {
	t.Run("should return a new instance of model", func(t *testing.T) {
		service := NewModelService()
		assert.NotNil(t, service)
		assert.IsType(t, &SModel{}, service)
	})
}
