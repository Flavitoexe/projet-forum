package controllers

import (
	"bluedit/forum-backend/helper"
	"bluedit/forum-backend/models"
	"bluedit/forum-backend/services_api"
	"encoding/json"
	"net/http"
)

type CategoryControllers struct {
	categoryService *services_api.CategoryServices
}

func InitCategoryControllers(categoryService *services_api.CategoryServices) *CategoryControllers {
	return &CategoryControllers{categoryService: categoryService}
}

func (c *CategoryControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur création catégorie (controller) - JSON invalide")
		return
	}

	categoryId, categoryErr := c.categoryService.Create(newCategory)
	if categoryErr != nil {
		helper.WriteError(w, http.StatusBadRequest, categoryErr.Error())
		return
	}

	category, categoryErr := c.categoryService.ReadById(categoryId)
	if categoryErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, categoryErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, category)
}

func (c *CategoryControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	categoriesList, categoriesErr := c.categoryService.ReadAll()
	if categoriesErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, categoriesErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, categoriesList)
}

func (c *CategoryControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur lecture catégorie par id (controller) - Identifiant invalide")
		return
	}

	category, categoryErr := c.categoryService.ReadById(id)
	if categoryErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, categoryErr.Error())
		return
	}
	if category.IdCategory <= 0 {
		helper.WriteError(w, http.StatusNotFound, "Erreur lecture catégorie par id (controller) - Catégorie introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, category)
}

func (c *CategoryControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour catégorie (controller) - Identifiant invalide")
		return
	}

	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour catégorie (controller) - JSON invalide")
		return
	}

	category.IdCategory = id

	categoryErr := c.categoryService.UpdateById(category)
	if categoryErr != nil {
		helper.WriteError(w, http.StatusBadRequest, categoryErr.Error())
		return
	}

	updatedCategory, updatedCategoryErr := c.categoryService.ReadById(id)
	if updatedCategoryErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, updatedCategoryErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedCategory)
}

func (c *CategoryControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur suppression catégorie (controller) - Identifiant invalide")
		return
	}

	categoryErr := c.categoryService.DeleteById(id)
	if categoryErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, categoryErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Catégorie supprimé",
	})
}
