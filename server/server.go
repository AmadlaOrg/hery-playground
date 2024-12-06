package server

import (
	"github.com/AmadlaOrg/hery-playground/server/controller"
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Start() error
}

type SServer struct {
	controller   controller.IController
	serverEngine *gin.Engine
}

func (s *SServer) Start() error {
	s.serverEngine = gin.Default()
	s.router()
	err := s.serverEngine.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}

func (s *SServer) router() {

	// env := os.Getenv("GIN_MODE")

	s.serverEngine.Static("/assets", "./browser/frontend/assets")

	s.serverEngine.GET("/", func(c *gin.Context) {
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.File("./browser/frontend/index.html")
	})

	s.serverEngine.GET("/:collectionName", func(c *gin.Context) {
		// Retrieve the 'collectionName' from the URL
		collectionName := c.Param("collectionName")

		// Use the collectionName for your logic
		c.JSON(200, gin.H{
			"message": "Collection: " + collectionName,
		})
	})
}
