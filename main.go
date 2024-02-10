package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func servePage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	type PageData struct {
		Title string
	}

	var tmplFiles []string
	data := PageData{Title: "Default Title"}

	tmplFiles = append(tmplFiles, filepath.Join("templates", "index.html"))

	switch path {
	case "/", "/index":
		tmplFiles = append(tmplFiles, filepath.Join("templates", "home_content.html"))
		data.Title = "Home Page"
	case "/projects":
		tmplFiles = append(tmplFiles, filepath.Join("templates", "projects_content.html"))
		data.Title = "Projects"
	}

	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	staticDir := "./Static"
	http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))).ServeHTTP(w, r)
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", servePage)
	mux.HandleFunc("/static/", staticHandler)
	return mux
}

func main() {
	mux := setupRoutes()

	PORT := ":23424"
	fmt.Printf("Server started on http://localhost%s\n", PORT)

	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
