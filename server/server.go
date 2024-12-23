package server

import (
	"github.com/AmadlaOrg/hery-playground/server/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"html/template"
	"os"
	"path/filepath"
)

type IServer interface {
	Start() error
}

type SServer struct {
	controller   controller.IController
	serverEngine *gin.Engine
}

var (
	ginDefault   = gin.Default
	filepathAbs  = filepath.Abs
	templateMust = template.Must
	templateNew  = template.New
	filepathWalk = filepath.Walk
	filepathExt  = filepath.Ext
)

// Start is a method to start an HTTP server
func (s *SServer) Start() error {
	s.serverEngine = ginDefault()

	// TODO: For debugging added to the config file
	// env := os.Getenv("GIN_MODE")
	gin.SetMode(gin.DebugMode)

	err := s.templateSetup(TmplPath)
	if err != nil {
		return err
	}
	s.router()

	// TODO: Use configuration file tor the port
	err = s.serverEngine.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}

// templateSetup attaches templates to the server engine
/*func (s *SServer) templateSetup(path string) error {
	absPath, err := filepathAbs(path)
	if err != nil {
		return err
	}

	templates := templateMust(s.parseTemplates(absPath))
	s.serverEngine.SetHTMLTemplate(templates)

	return nil
}*/

// templateSetup attaches templates to the server engine
func (s *SServer) templateSetup(path string) error {
	absPath, err := filepathAbs(path)
	if err != nil {
		return err
	}

	// Parse all HTML files in the directory, including subdirectories
	templates := templateMust(s.parseTemplates(absPath))
	s.serverEngine.SetHTMLTemplate(templates)

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

// parseTemplates walks through the directory and parses all .html files
/*func (s *SServer) parseTemplates(tmplRootPath string) (*template.Template, error) {
	tmpl := templateNew("")
	err := filepathWalk(tmplRootPath, func(path string, info os.FileInfo, err error) error { // Change to os.FileInfo
		if err != nil {
			return err
		}
		if !info.IsDir() && filepathExt(path) == ".html" { // Correct reference to IsDir
			_, err = tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return tmpl, err
}*/

// parseTemplates walks through the directory and parses all .html files
/*func (s *SServer) parseTemplates(tmplRootPath string) (*template.Template, error) {
	// Start with an empty template set
	tmpl := template.New("")
	err := filepathWalk(tmplRootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Only parse .html files, skip directories
		if !info.IsDir() && filepathExt(path) == ".html" {
			// Parse the file and add it to the template set
			println(path)
			_, err = tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return tmpl, err
}*/

func (s *SServer) parseTemplates(tmplRootPath string) (*template.Template, error) {
	tmpl := templateNew("")
	err := filepathWalk(tmplRootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepathExt(path) == ".html" {
			// TODO: Use a debug function (check flyaway project)
			println("Parsing template:", path)
			_, err = tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	// TODO: Move to debug
	if err == nil {
		// Log the names of registered templates
		for _, t := range tmpl.Templates() {
			// TODO: Use a debug function (check flyaway project)
			println("Registered template:", t.Name())
		}
	}
	return tmpl, err
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
