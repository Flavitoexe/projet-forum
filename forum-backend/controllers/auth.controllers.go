package controllers

import (
	"bluedit/forum-backend/auth"
	"bluedit/forum-backend/dto"
	"bluedit/forum-backend/helper"
	"bluedit/forum-backend/services_api"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthControllers struct {
	service *services_api.AuthService
}

func AuthController(authService *services_api.AuthService) *AuthControllers {
	return &AuthControllers{service: authService}
}

func (c *AuthControllers) Register(w http.ResponseWriter, r *http.Request) {
	var data dto.RegisterRequestDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	response, responseErr := c.service.Register(data)
	if responseErr != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiants invalides")
		return
	}

	helper.WriteJSON(w, http.StatusOK, response)
}

func (c *AuthControllers) Login(w http.ResponseWriter, r *http.Request) {
	var data dto.LoginRequestDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	response, err := c.service.Login(data)
	fmt.Println("response controller : ", response)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Identifiants invalides")
		return
	}

	fmt.Println("pas d'erreur, ok")
	helper.WriteJSON(w, http.StatusOK, response)
}

func (c *AuthControllers) Me(w http.ResponseWriter, r *http.Request) {
	claim, ok := r.Context().Value("user").(*auth.Claim)
	if !ok {
		helper.WriteError(w, http.StatusUnauthorized, "Unauthorized")
	}

	message := "Hello " + claim.Role + claim.UserId

	helper.WriteJSON(w, http.StatusOK, dto.ResponseDto{
		Code:    http.StatusOK,
		Message: message,
	})
}
