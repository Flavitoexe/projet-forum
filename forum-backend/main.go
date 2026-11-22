package main

import (
	"bluedit/forum-backend/app"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	app := app.InitApp()
	defer app.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	frontPort := os.Getenv("FRONT_PORT")
	if frontPort == "" {
		frontPort = ":8080"
	} else if !strings.HasPrefix(frontPort, ":") {
		frontPort = ":" + frontPort
	}

	go func() {
		log.Printf("Serveur client lancé : http://localhost%s", frontPort)
		if err := http.ListenAndServe(frontPort, app.FrontendRouter); err != nil {
			log.Printf("Erreur lancement serveur client - Erreur\n\t%s", err.Error())
		}
	}()

	log.Printf("Serveur API lancé : http://localhost:%s", port)
	apiServeErr := http.ListenAndServe(":"+port, app.RouterApi)
	if apiServeErr != nil {
		log.Fatalf("Erreur lancement serveur API - Erreur\n\t%s", apiServeErr.Error())
	}
}
