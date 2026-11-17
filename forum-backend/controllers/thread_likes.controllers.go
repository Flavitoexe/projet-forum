package controllers

import (
	"bluedit/forum-backend/helper"
	"bluedit/forum-backend/models"
	"bluedit/forum-backend/services_api"
	"encoding/json"
	"net/http"
)

type ThreadLikeControllers struct {
	threadLikeService *services_api.ThreadLikeServices
}

func InitThreadLikeControllers(threadLikeService *services_api.ThreadLikeServices) *ThreadLikeControllers {
	return &ThreadLikeControllers{threadLikeService: threadLikeService}
}

func (c *ThreadLikeControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newThreadLike models.ThreadLike
	err := json.NewDecoder(r.Body).Decode(&newThreadLike)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur création thread_like (controller) - JSON invalide")
		return
	}

	threadLikeId, threadLikeErr := c.threadLikeService.Create(newThreadLike)
	if threadLikeErr != nil {
		helper.WriteError(w, http.StatusBadRequest, threadLikeErr.Error())
		return
	}

	threadLike, threadLikeErr := c.threadLikeService.ReadById(threadLikeId)
	if threadLikeErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, threadLikeErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, threadLike)
}

func (c *ThreadLikeControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	threadLikesList, threadLikesErr := c.threadLikeService.ReadAll()
	if threadLikesErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, threadLikesErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, threadLikesList)
}

func (c *ThreadLikeControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur lecture thread_like par id (controller) - Identifiant invalide")
		return
	}

	threadLike, threadLikeErr := c.threadLikeService.ReadById(id)
	if threadLikeErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, threadLikeErr.Error())
		return
	}
	if threadLike.IdThreadLike <= 0 {
		helper.WriteError(w, http.StatusNotFound, "Erreur lecture thread_like par id (controller) - thread_like introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, threadLike)
}

func (c *ThreadLikeControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour thread_like (controller) - Identifiant invalide")
		return
	}

	var threadLike models.ThreadLike
	err := json.NewDecoder(r.Body).Decode(&threadLike)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour thread_like (controller) - JSON invalide")
		return
	}

	threadLike.IdThreadLike = id

	threadLikeErr := c.threadLikeService.UpdateById(threadLike)
	if threadLikeErr != nil {
		helper.WriteError(w, http.StatusBadRequest, threadLikeErr.Error())
		return
	}

	updatedThreadLike, updatedThreadLikeErr := c.threadLikeService.ReadById(id)
	if updatedThreadLikeErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, updatedThreadLikeErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedThreadLike)
}

func (c *ThreadLikeControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur suppression thread_like (controller) - Identifiant invalide")
		return
	}

	threadLikeErr := c.threadLikeService.DeleteById(id)
	if threadLikeErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, threadLikeErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "thread_like supprimé",
	})
}
