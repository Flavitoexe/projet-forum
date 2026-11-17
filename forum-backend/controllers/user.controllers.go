package controllers

import (
	"bluedit/forum-backend/helper"
	"bluedit/forum-backend/models"
	"bluedit/forum-backend/services_api"
	"encoding/json"
	"net/http"
)

type UserControllers struct {
	userService *services_api.UserServices
}

func InitUserControllers(userService *services_api.UserServices) *UserControllers {
	return &UserControllers{userService: userService}
}

func (c *UserControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur création utilisateur (controller) - JSON invalide")
		return
	}

	userId, userErr := c.userService.Create(newUser)
	if userErr != nil {
		helper.WriteError(w, http.StatusBadRequest, userErr.Error())
		return
	}

	user, userErr := c.userService.ReadById(userId)
	if userErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, userErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, user)
}

func (c *UserControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	usersList, usersErr := c.userService.ReadAll()
	if usersErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, usersErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, usersList)
}

func (c *UserControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur lecture utilisateur par id (controller) - Identifiant invalide")
		return
	}

	user, userErr := c.userService.ReadById(id)
	if userErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, userErr.Error())
		return
	}
	if user.IdUser <= 0 {
		helper.WriteError(w, http.StatusNotFound, "Erreur lecture utilisateur par id (controller) - utilisateur introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, user)
}

func (c *UserControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour utilisateur (controller) - Identifiant invalide")
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour utilisateur (controller) - JSON invalide")
		return
	}

	user.IdUser = id

	userErr := c.userService.UpdateById(user)
	if userErr != nil {
		helper.WriteError(w, http.StatusBadRequest, userErr.Error())
		return
	}

	updatedUser, updatedUserErr := c.userService.ReadById(id)
	if updatedUserErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, updatedUserErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedUser)
}

func (c *UserControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur suppression utilisateur (controller) - Identifiant invalide")
		return
	}

	userErr := c.userService.DeleteById(id)
	if userErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, userErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Utilisateur supprimé",
	})
}
