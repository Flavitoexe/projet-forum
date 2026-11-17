package controllers

import (
	"bluedit/forum-backend/helper"
	"bluedit/forum-backend/models"
	"bluedit/forum-backend/services_api"
	"encoding/json"
	"net/http"
)

type MessageLikeControllers struct {
	messageLikeService *services_api.MessageLikeServices
}

func InitMessageLikeControllers(messageLikeService *services_api.MessageLikeServices) *MessageLikeControllers {
	return &MessageLikeControllers{messageLikeService: messageLikeService}
}

func (c *MessageLikeControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newMessageLike models.MessageLike
	err := json.NewDecoder(r.Body).Decode(&newMessageLike)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur création message_like (controller) - JSON invalide")
		return
	}

	messageLikeId, messageLikeErr := c.messageLikeService.Create(newMessageLike)
	if messageLikeErr != nil {
		helper.WriteError(w, http.StatusBadRequest, messageLikeErr.Error())
		return
	}

	messageLike, messageLikeErr := c.messageLikeService.ReadById(messageLikeId)
	if messageLikeErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messageLikeErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, messageLike)
}

func (c *MessageLikeControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	messageLikesList, messageLikesErr := c.messageLikeService.ReadAll()
	if messageLikesErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messageLikesErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, messageLikesList)
}

func (c *MessageLikeControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur lecture message_like par id (controller) - Identifiant invalide")
		return
	}

	messageLike, messageLikeErr := c.messageLikeService.ReadById(id)
	if messageLikeErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messageLikeErr.Error())
		return
	}
	if messageLike.IdMessageLike <= 0 {
		helper.WriteError(w, http.StatusNotFound, "Erreur lecture message_like par id (controller) - message_like introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, messageLike)
}

func (c *MessageLikeControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour message_like (controller) - Identifiant invalide")
		return
	}

	var messageLike models.MessageLike
	err := json.NewDecoder(r.Body).Decode(&messageLike)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour message_like (controller) - JSON invalide")
		return
	}

	messageLike.IdMessageLike = id

	messageLikeErr := c.messageLikeService.UpdateById(messageLike)
	if messageLikeErr != nil {
		helper.WriteError(w, http.StatusBadRequest, messageLikeErr.Error())
		return
	}

	updatedmessageLike, updatedmessageLikeErr := c.messageLikeService.ReadById(id)
	if updatedmessageLikeErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, updatedmessageLikeErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedmessageLike)
}

func (c *MessageLikeControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur suppression message_like (controller) - Identifiant invalide")
		return
	}

	messageLikeErr := c.messageLikeService.DeleteById(id)
	if messageLikeErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messageLikeErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Message_like supprimé",
	})
}
