package services_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"bluedit/forum-backend/models"
	"bluedit/forum-frontend/helper"
)

func GetAllCategories() (*[]models.Category, int, error) {
	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Get(helper.BackendURL() + "/api/categories")
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur réponse - %s", err.Error())
	}
	defer resp.Body.Close()

	var categories []models.Category
	err = json.NewDecoder(resp.Body).Decode(&categories)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur decode données - %s", err.Error())
	}
	return &categories, http.StatusOK, err
}

func GetCategoryById(id int) (*models.Category, int, error) {
	var category models.Category
	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Get(fmt.Sprintf("%s/api/category/%v", helper.BackendURL(), id))
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur réponse - %s", err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&category)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur decode données - %s", err.Error())
	}
	return &category, http.StatusOK, nil
}

func CreateCategory(category models.Category) error {
	jsonData, err := json.Marshal(category)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Post(helper.BackendURL()+"/api/category", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erreur API statut: %d", resp.StatusCode)
	}
	return nil
}

func UpdateCategory(id string, category models.Category) error {
	jsonData, err := json.Marshal(category)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, helper.BackendURL()+"/api/category/"+id, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

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

func DeleteCategory(id string) error {
	req, err := http.NewRequest(http.MethodDelete, helper.BackendURL()+"/api/category/"+id, nil)
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
