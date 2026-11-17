package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	errLoad := godotenv.Load("../.env")
	if errLoad != nil {
		log.Println("Aucun fichier .env trouvé.\n\tUtilisation des variables d'environnement système.")
	}
}

func GetEnvWithDefault(key, defaultValue string) string {
	envVar, envErr := os.LookupEnv(key)
	if !envErr {
		return defaultValue
	}
	return envVar
}

func GetRequiredEnv(key string) string {
	envVar, envErr := os.LookupEnv(key)
	if !envErr || envVar == "" {
		log.Fatalf("Erreur configuration - Variable d'environnement manquante : %s", key)
	}
	return envVar
}
