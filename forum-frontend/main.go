package main

import (
	"bluedit/forum-frontend/app"
	"log"
	"net/http"
	"os"
)

func main() {
	app := app.InitApp()
	defer app.Close()

	port := os.Getenv("FRONT_PORT")
	if port == "" {
		port = ":8080"
	}
	log.Printf("Serveur client lancé : http://localhost%s", port)
	clientServeErr := http.ListenAndServe(port, app.RouterClient)
	if clientServeErr != nil {
		log.Fatalf("Erreur lancement serveur client - Erreur\n\t%s", clientServeErr.Error())
	}
}
