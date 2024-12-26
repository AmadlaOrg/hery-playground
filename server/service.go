package server

import (
	"github.com/AmadlaOrg/hery-playground/server/controller"
	"github.com/AmadlaOrg/hery-playground/server/template"
	"github.com/gin-gonic/gin"
)

var (
	ginDefault = gin.Default
)

// NewServerService to set up the server service
func NewServerService() IServer {
	return &SServer{
		controller:   controller.NewControllerService(),
		template:     template.NewTemplateService(TmplPath),
		serverEngine: ginDefault(),
	}
}
