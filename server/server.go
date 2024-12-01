package server

import "github.com/gin-gonic/gin"

type IServer interface{}

type SServer struct {
	serverEngine *gin.Engine
}

func (s *SServer) Start() {
	s.serverEngine = gin.Default()
	s.Router()
	err := s.serverEngine.Run(":8080")
	if err != nil {
		return
	}
}

func (s *SServer) Router() {
	s.serverEngine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
}
