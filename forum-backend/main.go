package main

import (
	"bluedit/forum-backend/app"
	"log"
	"net/http"
)

func main() {
	app := app.InitApp()
	defer app.Close()

	log.Printf("Serveur API lancé : http://localhost:3000")
	apiServeErr := http.ListenAndServe(":3000", app.RouterApi)
	if apiServeErr != nil {
		log.Fatalf("Erreur lancement serveur API - Erreur\n\t%s", apiServeErr.Error())
	}
}
