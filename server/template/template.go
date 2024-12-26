package template

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type ITemplate interface {
	Initialize()
	StartTemplateWatcher()
}
type STemplate struct {
	serverEngine    *gin.Engine
	templatePath    string
	templateAbsPath string
}

var (
	filepathAbs  = filepath.Abs
	templateMust = template.Must
	templateNew  = template.New
	filepathWalk = filepath.Walk
	filepathExt  = filepath.Ext
)

// Initialize attaches templates to the server engine
func (s *STemplate) Initialize() {
	// Parse all HTML files in the directory, including subdirectories
	templates := templateMust(s.parseTemplates(s.templateAbsPath))
	s.serverEngine.SetHTMLTemplate(templates)
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

func (s *STemplate) parseTemplates(tmplRootPath string) (*template.Template, error) {
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

func (s *STemplate) StartTemplateWatcher() {
	root, err := filepathAbs(s.templateAbsPath)
	if err != nil {
		//return err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {

		}
	}(watcher)

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write | fsnotify.Create | fsnotify.Remove | fsnotify.Rename) {
					log.Println("modified file:", event.Name)
					s.Initialize()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = addWatcherRecursively(watcher, root)
	if err != nil {
		return
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func addWatcherRecursively(watcher *fsnotify.Watcher, root string) error {
	return filepathWalk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}

		// Add only directories to the watcher
		if info.IsDir() {
			err := watcher.Add(path)
			if err != nil {
				log.Printf("Failed to add watcher for %s: %v\n", path, err)
				return err
			}
			log.Printf("Watching directory: %s\n", path)
		}
		return nil
	})
}
