package server

import "github.com/gin-gonic/gin"

type IServer interface{}

type SServer struct {
	serverEngine *gin.Engine
}

func (s *SServer) Start() {
	s.serverEngine = gin.Default()
	s.router()
	err := s.serverEngine.Run(":8080")
	if err != nil {
		return
	}
}

func (s *SServer) router() {
	s.serverEngine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
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
