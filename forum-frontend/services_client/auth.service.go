package services_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"bluedit/forum-frontend/dto"
)

type AuthService struct {
	backendURL string
}

func InitAuthService() *AuthService {
	return &AuthService{
		backendURL: "http://localhost:3000",
	}
}

func (s *AuthService) RegisterUser(data dto.RegisterRequestDto) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.New("Erreur lors de la préparation des données")
	}

	resp, err := http.Post(s.backendURL+"/api/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.New("Impossible de joindre le serveur backend")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return errors.New("Erreur lors de l'inscription")
	}

	return nil
}

func LoginUser(data dto.LoginRequestDto) (int, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return -1, errors.New("Erreur préparation données")
	}

	response, responseErr := http.Post("http://localhost:3000/api/login", "application/json", bytes.NewBuffer(jsonData))
	if responseErr != nil {
		return -1, errors.New("Erreur envoi requête")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return -1, errors.New("Erreur login")
	}

	var loginResponse dto.LoginResponseDto
	decodeErr := json.NewDecoder(response.Body).Decode(&loginResponse)
	if decodeErr != nil {
		return -1, errors.New("Erreur decode données")
	}

	return loginResponse.IdUser, nil
}
