package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServerService(t *testing.T) {
	t.Run("should return a new instance of server", func(t *testing.T) {
		service := NewServerService()
		assert.NotNil(t, service)
		assert.IsType(t, &SServer{}, service)
	})
}
