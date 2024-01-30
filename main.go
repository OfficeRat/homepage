package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func servePage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		path = "/index.html"
	}

	filePath := filepath.Join(".", "Pages", path)
	_, err := os.Stat(filePath)
	
	if err != nil {
		filePathWithExtension := filePath + ".html"
		_, err := os.Stat(filePathWithExtension)
	
		if err != nil {
			http.NotFound(w, r)
			return
		}
	
		filePath = filePathWithExtension
	}

	http.ServeFile(w, r, filePath)
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
	fmt.Printf("Server started on %s\n", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
