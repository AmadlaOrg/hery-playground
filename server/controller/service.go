package controller

import "github.com/AmadlaOrg/hery-playground/server/model"

// NewControllerService to set up the controller service
func NewControllerService() IController {
	return &SController{
		model: model.NewModelService(),
	}
}
