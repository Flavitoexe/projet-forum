package services_client

import (
	"bluedit/forum-frontend/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetUserById(id int) (*dto.DtoUser, int, error) {
	if id == -1 {
		user := dto.DtoUser{
			IdUser: -1,
			Name:   "visitor",
			Email:  "",
		}
		return &user, http.StatusOK, nil
	}
	_client := http.Client{Timeout: time.Second * 5}

	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:3000/api/user/%v", id), nil)
	if requestErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur préparation requête - %s", requestErr.Error())
	}

	response, responseErr := _client.Do(request)
	if responseErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur envoi requête - %s", responseErr.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil,
			response.StatusCode,
			fmt.Errorf("Erreur réponse - %s", response.Status)
	}

	var user dto.DtoUser
	err := json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur décodage données - %s", err.Error())
	}

	return &user, http.StatusOK, nil
}
