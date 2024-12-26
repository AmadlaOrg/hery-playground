package template

import (
	"github.com/gin-gonic/gin"
	"log"
)

// NewTemplateService to set up the model service
func NewTemplateService(serverEngine *gin.Engine, templatePath string) ITemplate {
	absPath, err := filepathAbs(templatePath)
	if err != nil {
		log.Println("error:", err)
	}

	return &STemplate{
		serverEngine:    serverEngine,
		templatePath:    templatePath,
		templateAbsPath: absPath,
	}
}
