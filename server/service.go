package server

import "github.com/AmadlaOrg/hery-playground/server/controller"

// NewServerService to set up the server service
func NewServerService() IServer {
	return &SServer{
		controller: controller.NewControllerService(),
	}
}
