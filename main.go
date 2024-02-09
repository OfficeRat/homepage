package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func servePage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// Define a struct for the dynamic data you want to pass to your templates
	type PageData struct {
		Title string
	}
	var tmplFiles []string
	data := PageData{Title: "Default Title"}
	// Always include the main layout
	tmplFiles = append(tmplFiles, filepath.Join("templates", "index.html"))
	// Determine which content template to use based on the request
	if path == "/" || path == "/index" {
		tmplFiles = append(tmplFiles, filepath.Join("templates", "home_content.html"))
		data.Title = "Home Page"
	} else if path == "/projects" {
		tmplFiles = append(tmplFiles, filepath.Join("templates", "projects_content.html"))
		data.Title = "Projects"
	}
	// Parse the templates
	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the main layout template
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	staticDir := "./Static"
	http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))).ServeHTTP(w, r)
}

func setupRoutes() {
	err := filepath.Walk("./Pages", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			relPath, err := filepath.Rel("./Pages", path)

			if err != nil {
				return err
			}

			if relPath == "index.html" {
				http.HandleFunc("/", servePage)
			} else {
				http.HandleFunc("/"+strings.TrimSuffix(relPath, ".html"), servePage)
				http.HandleFunc("/"+strings.TrimSuffix(relPath, ".html")+".html", servePage)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through Pages folder: %v\n", err)
	}

	http.HandleFunc("/static/", staticHandler)
}

func main() {
	setupRoutes()

	PORT := ":23423"
	fmt.Printf("Server started on http://localhost%s\n", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
