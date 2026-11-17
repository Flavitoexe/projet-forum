package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var listeTemplate *template.Template

func Load() {
	// ParseGlob se fait depuis main.go (point d'entrée de l'app)
	temp, tempErr := template.ParseGlob("./templates/*.html")
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
