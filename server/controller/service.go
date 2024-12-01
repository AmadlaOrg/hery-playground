package controller

// NewControllerService to set up the controller service
func NewControllerService() IController {
	return &SController{}
}
