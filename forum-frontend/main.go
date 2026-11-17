package main

import (
	"bluedit/forum-frontend/app"
	"log"
	"net/http"
)

func main() {
	app := app.InitApp()
	defer app.Close()

	log.Printf("Serveur client lancé : http://localhost:8080")
	clientServeErr := http.ListenAndServe(":8080", app.RouterClient)
	if clientServeErr != nil {
		log.Fatalf("Erreur lancement serveur client - Erreur\n\t%s", clientServeErr.Error())
	}
}
