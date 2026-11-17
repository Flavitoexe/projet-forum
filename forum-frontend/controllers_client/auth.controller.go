package controllers_client

import (
	"bluedit/forum-frontend/dto"
	"bluedit/forum-frontend/helper"
	"bluedit/forum-frontend/services_client"
	"fmt"
	"net/http"
)

type AuthController struct {
	authService *services_client.AuthService
}

func InitAuthController(authService *services_client.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) PostRegister(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur formulaire", http.StatusBadRequest)
		return
	}

	reqData := dto.RegisterRequestDto{
		Email:    r.FormValue("email"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	if err := c.authService.RegisterUser(reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur formulaire", http.StatusBadRequest)
		return
	}

	reqData := dto.LoginRequestDto{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	idUser, err := services_client.LoginUser(reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usernameId := fmt.Sprintf("%v:%v", reqData.Username, idUser)

	SetSessionCookie(w, usernameId)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	ClearSessionCookie(w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DisplayRegister(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, r, "register", nil)
}

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, r, "login", nil)
}
