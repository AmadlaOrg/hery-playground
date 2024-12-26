package server

import (
	"github.com/AmadlaOrg/hery-playground/server/controller"
	"github.com/AmadlaOrg/hery-playground/server/template"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path/filepath"
)

type IServer interface {
	Start() error
}

type SServer struct {
	controller   controller.IController
	template     template.ITemplate
	serverEngine *gin.Engine
}

var (
	filepathAbs = filepath.Abs
)

// Start is a method to start an HTTP server
func (s *SServer) Start() error {
	//s.serverEngine = ginDefault()

	// TODO: For debugging added to the config file
	// env := os.Getenv("GIN_MODE")
	gin.SetMode(gin.DebugMode)

	// Setup templates
	err := s.template.Initialize(s.serverEngine, TmplPath)
	if err != nil {
		return err
	}

	s.template.StartTemplateWatcher()

	s.router()

	// TODO: Use configuration file tor the port
	err = s.serverEngine.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}

// router will load all the API endpoints
func (s *SServer) router() {
	// 1. Set the static file support
	s.serverEngine.Static("/assets", StaticFilesPath)

	// 2. Add custom error handlers
	s.serverEngine.Use(s.errorMiddleware()) // Catch panics for 500 errors
	s.serverEngine.NoRoute(func(c *gin.Context) {
		c.Header("Cache-Control", CacheControlHeaderValue)
		c.HTML(404, "404.html", gin.H{})
	})

	// 3. Set the front end main endpoint
	s.serverEngine.GET("/", func(c *gin.Context) {
		c.Header("Cache-Control", CacheControlHeaderValue)
		c.HTML(200, "index.html", gin.H{
			"Title": "Hery Playground", // Pass additional variables if needed
		})
	})

	// 4. Example dynamic endpoint
	s.serverEngine.GET("/:collectionName", func(c *gin.Context) {
		// Retrieve the 'collectionName' from the URL
		collectionName := c.Param("collectionName")

		// Use the collectionName for your logic
		c.JSON(200, gin.H{
			"message": "Collection: " + collectionName,
		})
	})
}

// openConfig loads the server configurations
func (s *SServer) openConfig() {
	viper.SetConfigName("config")
}

// Middleware to catch panics and render 500 error page
func (s *SServer) errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log the error for debugging

				// TODO: Maybe add support logs (check how server logs works)
				//s.serverLogger.Errorf("Recovered from panic: %v", r)

				// Respond with 500 error page
				c.Header("Cache-Control", CacheControlHeaderValue)
				c.HTML(500, "500.html", gin.H{
					"Title": "Internal Server Error",
					"Error": r,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
