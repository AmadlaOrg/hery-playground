package controller

import "github.com/AmadlaOrg/hery-playground/server/model"

type IController interface{}

type SController struct {
	model model.IModel
}

// OpenHeryFile
func (s *SController) OpenHeryFile(collectionName string) {

}

// OpenHeryCache
func (s *SController) OpenHeryCache(collectionName string) {

}
