package controllers

import (
	"bluedit/forum-backend/dto"
	"bluedit/forum-backend/helper"
	"bluedit/forum-backend/models"
	"bluedit/forum-backend/services_api"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ThreadControllers struct {
	threadService *services_api.ThreadServices
}

func InitThreadControllers(threadService *services_api.ThreadServices) *ThreadControllers {
	return &ThreadControllers{threadService: threadService}
}

func readIdParams(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["id"])
}

func (c *ThreadControllers) Create(w http.ResponseWriter, r *http.Request) {
	var newThread dto.DtoThreadRequest
	err := json.NewDecoder(r.Body).Decode(&newThread)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur création thread (controller) - JSON invalide")
		return
	}

	threadId, threadErr := c.threadService.Create(newThread)
	if threadErr != nil {
		helper.WriteError(w, http.StatusBadRequest, threadErr.Error())
		return
	}

	thread, threadErr := c.threadService.ReadById(threadId)
	if threadErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, threadErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, thread)
}

func (c *ThreadControllers) ReadAll(w http.ResponseWriter, r *http.Request) {
	threadsList, threadsErr := c.threadService.ReadAll()
	if threadsErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, threadsErr.Error())
	}

	helper.WriteJSON(w, http.StatusOK, threadsList)
}

func (c *ThreadControllers) ReadById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur lecture thread par id (controller) - Identifiant invalide")
		return
	}

	thread, threadErr := c.threadService.ReadById(id)
	if threadErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, threadErr.Error())
		return
	}
	if thread.IdThread <= 0 {
		helper.WriteError(w, http.StatusNotFound, "Erreur lecture thread par id (controller) - Thread introuvable")
		return
	}

	helper.WriteJSON(w, http.StatusOK, thread)
}

func (c *ThreadControllers) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour thread (controller) - Identifiant invalide")
		return
	}

	var thread models.Thread
	err := json.NewDecoder(r.Body).Decode(&thread)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur mise à jour thread (controller) - JSON invalide")
		return
	}

	thread.IdThread = id

	threadErr := c.threadService.UpdateById(thread)
	if threadErr != nil {
		helper.WriteError(w, http.StatusBadRequest, threadErr.Error())
		return
	}

	updatedThread, updatedThreadErr := c.threadService.ReadById(id)
	if updatedThreadErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, updatedThreadErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, updatedThread)
}

func (c *ThreadControllers) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Erreur suppression thread (controller) - Identifiant invalide")
		return
	}

	threadErr := c.threadService.DeleteById(id)
	if threadErr != nil {
		helper.WriteError(w, http.StatusInternalServerError, threadErr.Error())
		return
	}

	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Thread supprimé",
	})
}

func (c *ThreadControllers) GetHealth(w http.ResponseWriter, r *http.Request) {
	helper.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Backend OK",
	})
}
