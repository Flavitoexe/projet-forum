package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var listeTemplate *template.Template

func ResolveDir(relativePath string) string {
	cwd, err := os.Getwd()
	if err == nil {
		for dir := cwd; ; dir = filepath.Dir(dir) {
			candidates := []string{
				filepath.Join(dir, relativePath),
				filepath.Join(dir, "forum-frontend", relativePath),
			}
			for _, candidate := range candidates {
				if info, statErr := os.Stat(candidate); statErr == nil && info.IsDir() {
					return candidate
				}
			}

			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
		}
	}
	return relativePath
}

func BackendURL() string {
	if backendURL := os.Getenv("BACKEND_URL"); backendURL != "" {
		return strings.TrimRight(backendURL, "/")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	port = strings.TrimPrefix(port, ":")
	return "http://localhost:" + port
}

func Load() {
	templateDir := ResolveDir("templates")
	temp, tempErr := template.ParseGlob(filepath.Join(templateDir, "*.html"))
	if tempErr != nil {
		log.Fatalf("Erreur template - %s", tempErr.Error())
		return
	}
	listeTemplate = temp
	// --------------
	fmt.Println("templates name : ", listeTemplate.Name())
	// --------------
	fmt.Println("Template - chargement des templates terminé")
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var buffer bytes.Buffer

	errRender := listeTemplate.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		fmt.Println(errRender)
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	buffer.WriteTo(w)
}
