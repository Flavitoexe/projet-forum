package services_client

import (
	"bluedit/forum-backend/models"
	"bluedit/forum-frontend/dto"
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"
)

func CreateThread(newThread dto.DtoThread) (*dto.DtoThread, int, error) {
	jsonData, err := json.Marshal(newThread)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur Marshall données : Création thread (service frontend) - %s", err.Error())
	}
	fmt.Println("jsondata : ", string(jsonData))

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Post("http://localhost:3000/api/thread", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur requête : Création thread (service frontend) - %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		fmt.Println("resp.Statuscode : ", resp.StatusCode)
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur réponse : Création thread (service frontend) - ")
	}

	var thread dto.DtoThread
	decodeErr := json.NewDecoder(resp.Body).Decode(&thread)
	if decodeErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode données - %s", decodeErr.Error())
	}

	return &thread, http.StatusOK, nil
}

func GetAllThreads() (*[]models.Thread, int, error) {
	_client := http.Client{Timeout: time.Second * 5}

	request, requestErr := http.NewRequest(http.MethodGet, "http://localhost:3000/api/threads", nil)
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

	var threadsList []models.Thread

	decodeErr := json.NewDecoder(response.Body).Decode(&threadsList)
	if decodeErr != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode données - %s", decodeErr.Error())
	}

	return &threadsList, response.StatusCode, nil
}

func GetThreadById(id int) (*dto.DtoThread, int, error) {
	_client := http.Client{Timeout: time.Second * 5}

	request, requestErr := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:3000/api/thread/%v", id), nil)
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

	var thread dto.DtoThread
	err := json.NewDecoder(response.Body).Decode(&thread)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur décodage données - %s", err.Error())
	}

	return &thread, http.StatusOK, nil
}

func GetThreadsByCategory(id int) (*[]models.Thread, int, error) {
	_client := http.Client{Timeout: 5 * time.Second}

	request, requestErr := http.NewRequest(http.MethodGet, "http://localhost:3000/api/threads", nil)
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

	var threads []models.Thread
	err := json.NewDecoder(response.Body).Decode(&threads)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode données - %s", err.Error())
	}

	var sortedThreads []models.Thread
	for _, thread := range threads {
		if thread.FkIdCategory == id {
			sortedThreads = append(sortedThreads, thread)
		}
	}

	return &sortedThreads, http.StatusOK, nil
}

func GetThreadsByPopularity() (*[]models.Thread, int, error) {
	_client := http.Client{Timeout: 5 * time.Second}

	request, requestErr := http.NewRequest(http.MethodGet, "http://localhost:3000/api/threads", nil)
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

	var threads []models.Thread
	err := json.NewDecoder(response.Body).Decode(&threads)
	if err != nil {
		return nil,
			http.StatusInternalServerError,
			fmt.Errorf("Erreur decode données - %s", err.Error())
	}

	slices.SortFunc(threads, func(a, b models.Thread) int {
		aScore := a.LikeCount + (a.MessageCount * 2)
		bScore := b.LikeCount + (b.MessageCount * 2)

		return cmp.Compare(bScore, aScore)
	})

	return &threads, http.StatusOK, nil
}

func UpdateThread(newThread dto.DtoThreadRequest) error {
	jsonData, err := json.Marshal(newThread)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:3000/api/thread/%v", newThread.IdThread), bytes.NewBuffer(jsonData))
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

func DeleteThread(idThread int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:3000/api/thread/%v", idThread), nil)
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
