package controllers

import (
	"bluedit/forum-backend/dto"
	"bluedit/forum-backend/helper"
	"bluedit/forum-backend/models"
	"bluedit/forum-backend/services_api"
	"encoding/json"
	"fmt"
	"net/http"
)

type MessageControllers struct {
	messageService *services_api.MessageServices
}

func InitMessageControllers(messageService *services_api.MessageServices) *MessageControllers {
	return &MessageControllers{messageService: messageService}
}

func (c *MessageControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newMessage dto.DtoMessageRequest
	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur création message (controller) - JSON invalide")
		return
	}

	messageId, messageErr := c.messageService.Create(newMessage)
	if messageErr != nil {
		helper.WriteError(w, http.StatusBadRequest, messageErr.Error())
		return
	}

	message, messageErr := c.messageService.ReadById(messageId)
	if messageErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messageErr.Error())
	}

	fmt.Println("ReadById ok, WriteJSON ok avec message")
	helper.WriteJSON(w, http.StatusOK, message)
}

func (c *MessageControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	messagesList, messagesErr := c.messageService.ReadAll()
	if messagesErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messagesErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, messagesList)
}

func (c *MessageControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur lecture message par id (controller) - Identifiant invalide")
		return
	}

	message, messageErr := c.messageService.ReadById(id)
	if messageErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messageErr.Error())
		return
	}
	if message.IdMessage <= 0 {
		helper.WriteError(w, http.StatusNotFound, "Erreur lecture message par id (controller) - Message introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, message)
}

func (c *MessageControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour message (controller) - Identifiant invalide")
		return
	}

	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour message (controller) - JSON invalide")
		return
	}

	message.IdMessage = id

	messageErr := c.messageService.UpdateById(message)
	if messageErr != nil {
		helper.WriteError(w, http.StatusBadRequest, messageErr.Error())
		return
	}

	updatedMessage, updatedMessageErr := c.messageService.ReadById(id)
	if updatedMessageErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, updatedMessageErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedMessage)
}

func (c *MessageControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur suppression message (controller) - Identifiant invalide")
		return
	}

	messageErr := c.messageService.DeleteById(id)
	if messageErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messageErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Message supprimé",
	})
}

func (c *MessageControllers) ReadByThread(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur lecture messages par thread (controller) - Identifiant invalide")
		return
	}

	messages, messagesErr := c.messageService.ReadByThread(id)
	if messagesErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, messagesErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, messages)
}
