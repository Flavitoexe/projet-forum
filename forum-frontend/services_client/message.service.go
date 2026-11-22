package services_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"bluedit/forum-backend/models"
	"bluedit/forum-frontend/dto"
	"bluedit/forum-frontend/helper"
)

func GetAllMessages() ([]models.Message, error) {
	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Get(helper.BackendURL() + "/api/messages")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var messages []models.Message
	err = json.NewDecoder(resp.Body).Decode(&messages)
	return messages, err
}

func GetMessageById(id string) (models.Message, error) {
	var message models.Message
	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Get(helper.BackendURL() + "/api/message/" + id)
	if err != nil {
		return message, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&message)
	return message, err
}

func CreateMessage(message dto.DtoMessage) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Post(helper.BackendURL()+"/api/message", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erreur API statut: %d", resp.StatusCode)
	}

	return nil
}

func UpdateMessage(updatedMessage dto.DtoMessageRequest) error {
	jsonData, err := json.Marshal(updatedMessage)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/message/%v", helper.BackendURL(), updatedMessage.IdMessage), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Erreur réponse - %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erreur API statut: %d", resp.StatusCode)
	}
	return nil
}

func DeleteMessage(id int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/message/%v", helper.BackendURL(), id), nil)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erreur API statut: %d", resp.StatusCode)
	}
	return nil
}

func GetMessagesByThread(id int) (*[]dto.DtoMessage, int, error) {
	_client := http.Client{Timeout: time.Second * 5}

	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/messages/%v", helper.BackendURL(), id), nil)
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

	var messages []dto.DtoMessage

	err := json.NewDecoder(response.Body).Decode(&messages)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur décodage données - %s", err.Error())
	}

	return &messages, http.StatusOK, nil
}
