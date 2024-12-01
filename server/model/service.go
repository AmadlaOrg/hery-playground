package model

// NewModelService to set up the model service
func NewModelService() IModel {
	return &SModel{}
}
