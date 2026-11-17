package controllers_client

import (
	"bluedit/forum-backend/models"
	"bluedit/forum-frontend/dto"
	"bluedit/forum-frontend/helper"
	"bluedit/forum-frontend/services_client"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func readIdParams(r *http.Request) (int, error) {
	return strconv.Atoi(mux.Vars(r)["id"])
}

func loggedIn(r *http.Request) (int, error) {
	cookie, cookieErr := r.Cookie("bluedit_session")
	if cookieErr != nil {
		if cookieErr == http.ErrNoCookie {
			return -1, nil
		}
		return -2, fmt.Errorf("Erreur récupération cookie")
	}

	usernameId := cookie.Value
	idStr := strings.Split(usernameId, ":")[1]
	if len(idStr) <= 0 {
		return -2, fmt.Errorf("Erreur récupération ID dans cookie")
	}

	idUser, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		return -2, fmt.Errorf("Erreur conversion ID string à int")
	}

	return idUser, nil
}

func DisplayIndex(w http.ResponseWriter, r *http.Request) {
	var dataTemp dto.TempIndexData
	dataTemp.ThreadsData.Data, dataTemp.ThreadsData.DataStatusCode, dataTemp.ThreadsData.DataErr = services_client.GetAllThreads()
	if dataTemp.ThreadsData.DataStatusCode != http.StatusOK || dataTemp.ThreadsData.DataErr != nil {
		return
	}

	dataTemp.CategoriesData.Data, dataTemp.CategoriesData.DataStatusCode, dataTemp.CategoriesData.DataErr = services_client.GetAllCategories()
	if dataTemp.CategoriesData.DataStatusCode != http.StatusOK || dataTemp.CategoriesData.DataErr != nil {
		return
	}

	idUser, idUserErr := loggedIn(r)
	if idUserErr != nil {
		return
	}
	dataTemp.UserData.Data, dataTemp.UserData.DataStatusCode, dataTemp.UserData.DataErr = services_client.GetUserById(idUser)
	if dataTemp.UserData.DataStatusCode != http.StatusOK || dataTemp.UserData.DataErr != nil {
		return
	}

	helper.RenderTemplate(w, r, "index", dataTemp)
}

func DisplayThread(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		return
	}

	var dataTemp dto.TempThreadDetailsData

	dataTemp.ThreadData.Data, dataTemp.ThreadData.DataStatusCode, dataTemp.ThreadData.DataErr = services_client.GetThreadById(id)
	if dataTemp.ThreadData.DataStatusCode != http.StatusOK || dataTemp.ThreadData.DataErr != nil {
		fmt.Println("erreur get thread")
		http.RedirectHandler("/error/404", http.StatusNotFound)
	}

	dataTemp.MessagesData.Data, dataTemp.MessagesData.DataStatusCode, dataTemp.MessagesData.DataErr = services_client.GetMessagesByThread(id)
	if dataTemp.MessagesData.DataStatusCode != http.StatusOK || dataTemp.MessagesData.DataErr != nil {
		fmt.Println("erreur getmessages : ", dataTemp.MessagesData.DataErr)
		return
	}

	dataTemp.CategoriesData.Data, dataTemp.CategoriesData.DataStatusCode, dataTemp.CategoriesData.DataErr = services_client.GetAllCategories()
	if dataTemp.CategoriesData.DataStatusCode != http.StatusOK || dataTemp.CategoriesData.DataErr != nil {
		return
	}

	idUser, idUserErr := loggedIn(r)
	if idUserErr != nil {
		return
	}
	dataTemp.UserData.Data, dataTemp.UserData.DataStatusCode, dataTemp.UserData.DataErr = services_client.GetUserById(idUser)
	if dataTemp.UserData.DataStatusCode != http.StatusOK || dataTemp.UserData.DataErr != nil {
		fmt.Println("erreur getuser")
		return
	}

	helper.RenderTemplate(w, r, "thread", dataTemp)
}

func DisplayCategory(w http.ResponseWriter, r *http.Request) {
	var dataTemp dto.TempCategoryDetailsData
	id, idErr := readIdParams(r)
	if idErr != nil {
		return
	}

	dataTemp.CategoryData.Data, dataTemp.CategoryData.DataStatusCode, dataTemp.CategoryData.DataErr = services_client.GetCategoryById(id)
	if dataTemp.CategoryData.DataStatusCode != http.StatusOK || dataTemp.CategoryData.DataErr != nil {
		return
	}

	dataTemp.ThreadsData.Data, dataTemp.ThreadsData.DataStatusCode, dataTemp.ThreadsData.DataErr = services_client.GetThreadsByCategory(id)
	if dataTemp.ThreadsData.DataStatusCode != http.StatusOK || dataTemp.ThreadsData.DataErr != nil {
		return
	}

	dataTemp.CategoriesData.Data, dataTemp.CategoriesData.DataStatusCode, dataTemp.CategoriesData.DataErr = services_client.GetAllCategories()
	if dataTemp.CategoriesData.DataStatusCode != http.StatusOK || dataTemp.CategoriesData.DataErr != nil {
		return
	}

	idUser, idUserErr := loggedIn(r)
	if idUserErr != nil {
		return
	}
	dataTemp.UserData.Data, dataTemp.UserData.DataStatusCode, dataTemp.UserData.DataErr = services_client.GetUserById(idUser)
	if dataTemp.UserData.DataStatusCode != http.StatusOK || dataTemp.UserData.DataErr != nil {
		return
	}

	helper.RenderTemplate(w, r, "category", dataTemp)
}

func DisplayCategories(w http.ResponseWriter, r *http.Request) {
	var dataTemp dto.TempCategoriesDetailsData

	dataTemp.CategoriesData.Data, dataTemp.CategoriesData.DataStatusCode, dataTemp.CategoriesData.DataErr = services_client.GetAllCategories()
	if dataTemp.CategoriesData.DataErr != nil || dataTemp.CategoriesData.DataStatusCode != http.StatusOK {
		return
	}

	idUser, idUserErr := loggedIn(r)
	if idUserErr != nil {
		return
	}
	dataTemp.UserData.Data, dataTemp.UserData.DataStatusCode, dataTemp.UserData.DataErr = services_client.GetUserById(idUser)
	if dataTemp.UserData.DataStatusCode != http.StatusOK || dataTemp.UserData.DataErr != nil {
		return
	}

	helper.RenderTemplate(w, r, "categories", dataTemp)
}

func DisplayNews(w http.ResponseWriter, r *http.Request) {
	var dataTemp dto.TempIndexData
	dataTemp.ThreadsData.Data, dataTemp.ThreadsData.DataStatusCode, dataTemp.ThreadsData.DataErr = services_client.GetAllThreads()
	if dataTemp.ThreadsData.DataStatusCode != http.StatusOK || dataTemp.ThreadsData.DataErr != nil {
		return
	}

	dataTemp.CategoriesData.Data, dataTemp.CategoriesData.DataStatusCode, dataTemp.CategoriesData.DataErr = services_client.GetAllCategories()
	if dataTemp.CategoriesData.DataStatusCode != http.StatusOK || dataTemp.CategoriesData.DataErr != nil {
		return
	}

	var sortedThreads []models.Thread

	for _, thread := range *dataTemp.ThreadsData.Data {
		if time.Since(thread.CreatedAt) <= time.Hour*48 {
			sortedThreads = append(sortedThreads, thread)
		}
	}

	dataTemp.ThreadsData.Data = &sortedThreads

	idUser, idUserErr := loggedIn(r)
	if idUserErr != nil {
		return
	}
	dataTemp.UserData.Data, dataTemp.UserData.DataStatusCode, dataTemp.UserData.DataErr = services_client.GetUserById(idUser)
	if dataTemp.UserData.DataStatusCode != http.StatusOK || dataTemp.UserData.DataErr != nil {
		return
	}

	helper.RenderTemplate(w, r, "news", dataTemp)
}

func DisplayPopulars(w http.ResponseWriter, r *http.Request) {
	var dataTemp dto.TempIndexData
	dataTemp.ThreadsData.Data, dataTemp.ThreadsData.DataStatusCode, dataTemp.ThreadsData.DataErr = services_client.GetThreadsByPopularity()
	if dataTemp.ThreadsData.DataStatusCode != http.StatusOK || dataTemp.ThreadsData.DataErr != nil {
		return
	}

	dataTemp.CategoriesData.Data, dataTemp.CategoriesData.DataStatusCode, dataTemp.CategoriesData.DataErr = services_client.GetAllCategories()
	if dataTemp.CategoriesData.DataStatusCode != http.StatusOK || dataTemp.CategoriesData.DataErr != nil {
		return
	}

	idUser, idUserErr := loggedIn(r)
	if idUserErr != nil {
		return
	}
	dataTemp.UserData.Data, dataTemp.UserData.DataStatusCode, dataTemp.UserData.DataErr = services_client.GetUserById(idUser)
	if dataTemp.UserData.DataStatusCode != http.StatusOK || dataTemp.UserData.DataErr != nil {
		return
	}

	helper.RenderTemplate(w, r, "populars", dataTemp)
}

func DisplayError(w http.ResponseWriter, r *http.Request) {
	id, idErr := readIdParams(r)
	if idErr != nil {
		return
	}

	var dataTemp dto.TempErrorDetailsData
	dataTemp.ErrorData.CategoriesData.Data, dataTemp.ErrorData.CategoriesData.DataStatusCode, dataTemp.ErrorData.CategoriesData.DataErr = services_client.GetAllCategories()

	if id == http.StatusNotFound {
		dataTemp.ErrorData.Code = http.StatusNotFound
		dataTemp.ErrorData.Message = "Désolé, y a eu un petit problème et on a pas pu trouver la ressource que vous recherchez.."
		helper.RenderTemplate(w, r, "error", dataTemp)
	}

	idUser, idUserErr := loggedIn(r)
	if idUserErr != nil {
		return
	}
	dataTemp.UserData.Data, dataTemp.UserData.DataStatusCode, dataTemp.UserData.DataErr = services_client.GetUserById(idUser)
	if dataTemp.UserData.DataStatusCode != http.StatusOK || dataTemp.UserData.DataErr != nil {
		return
	}

}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur de lecture du formulaire", http.StatusBadRequest)
		return
	}

	idThread, idErr := readIdParams(r)
	if idErr != nil {
		http.Error(w, "Erreur récupération identifiant", 500)
		return
	}

	idUser, idErr := loggedIn(r)

	reqData := dto.DtoMessage{
		Content:    r.FormValue("content"),
		FkIdThread: idThread,
		FkIdUser:   idUser,
	}

	postErr := services_client.CreateMessage(reqData)
	if postErr != nil {
		http.Error(w, postErr.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/thread/%v", idThread), http.StatusSeeOther)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur de lecture du formulaire", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")

	idThread, idErr := strconv.Atoi(mux.Vars(r)["idThread"])
	if idErr != nil {
		http.Error(w, "Erreur récupération ID thread", http.StatusBadRequest)
	}
	idMessage, idErr := strconv.Atoi(mux.Vars(r)["idMessage"])
	if idErr != nil {
		http.Error(w, "Erreur récupération ID message", http.StatusBadRequest)
		return
	}

	idUser, idErr := loggedIn(r)
	if idUser < 0 || idErr != nil {
		http.Error(w, "Erreur utilisateur pas connecté", http.StatusBadRequest)
		return
	}

	updatedMessage := dto.DtoMessageRequest{
		IdMessage: idMessage,
		Content:   content,
	}

	updateErr := services_client.UpdateMessage(updatedMessage)
	if updateErr != nil {
		http.Error(w, "Erreur mise à jour message (voir services client)", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/thread/%v", idThread), http.StatusSeeOther)
}

func UpdateThread(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur de lecture du formulaire", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	content := r.FormValue("content")
	category, categoryErr := strconv.Atoi(r.FormValue("category"))
	if categoryErr != nil {
		http.Error(w, "Erreur conversion id categorie (update thread)", http.StatusBadRequest)
		return
	}

	idThread, idErr := strconv.Atoi(mux.Vars(r)["idThread"])
	if idErr != nil {
		http.Error(w, "Erreur récupération ID thread", http.StatusBadRequest)
		return
	}

	idUser, idErr := loggedIn(r)
	if idUser < 0 || idErr != nil {
		http.Error(w, "Erreur utilisateur pas connecté", http.StatusBadRequest)
		return
	}

	updatedThread := dto.DtoThreadRequest{
		IdThread:     idThread,
		Name:         name,
		Content:      content,
		FkIdCategory: category,
	}

	updateErr := services_client.UpdateThread(updatedThread)
	if updateErr != nil {
		fmt.Println("updateErr : ", updateErr)
		http.Error(w, "Erreur mise à jour thread (voir services client)", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/thread/%v", idThread), http.StatusSeeOther)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur de lecture du formulaire", http.StatusBadRequest)
		return
	}

	idThread, idErr := strconv.Atoi(mux.Vars(r)["idThread"])
	if idErr != nil {
		http.Error(w, "Erreur récupération ID thread", http.StatusBadRequest)
	}
	idMessage, idErr := strconv.Atoi(mux.Vars(r)["idMessage"])
	if idErr != nil {
		http.Error(w, "Erreur récupération ID message", http.StatusBadRequest)
		return
	}

	idUser, idErr := loggedIn(r)
	if idUser < 0 || idErr != nil {
		http.Error(w, "Erreur utilisateur pas connecté", http.StatusBadRequest)
		return
	}

	deleteErr := services_client.DeleteMessage(idMessage)
	if deleteErr != nil {
		http.Error(w, "Erreur suppression message (voir services client)", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/thread/%v", idThread), http.StatusSeeOther)
}

func DeleteThread(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur de lecture du formulaire", http.StatusBadRequest)
		return
	}

	idThread, idErr := strconv.Atoi(mux.Vars(r)["idThread"])
	if idErr != nil {
		http.Error(w, "Erreur récupération ID thread", http.StatusBadRequest)
	}
	idCategory, idErr := strconv.Atoi(mux.Vars(r)["idCategory"])
	if idErr != nil {
		http.Error(w, "Erreur récupération ID catégorie", http.StatusBadRequest)
		return
	}

	idUser, idErr := loggedIn(r)
	if idUser < 0 || idErr != nil {
		http.Error(w, "Erreur utilisateur pas connecté", http.StatusBadRequest)
		return
	}

	deleteErr := services_client.DeleteThread(idThread)
	if deleteErr != nil {
		http.Error(w, "Erreur suppression thread (voir services client)", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/category/%v", idCategory), http.StatusSeeOther)
}

func PostThread(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur de lecture du formulaire", http.StatusBadRequest)
		return
	}

	var idCategory int
	var idCategoryErr error
	if r.FormValue("category") == "" {
		idCategory, idCategoryErr = readIdParams(r)
		if idCategoryErr != nil {
			http.Error(w, "Erreur récupération id category dans url", http.StatusBadRequest)
		}
	} else {
		idCategory, idCategoryErr = strconv.Atoi(r.FormValue("category"))
		if idCategoryErr != nil {
			http.Error(w, "Erreur récupération identifiant catégorie", 500)
			return
		}
	}

	cookie, cookieErr := r.Cookie("bluedit_session")
	if cookieErr != nil {
		http.Error(w, "Erreur récupération cookies", http.StatusBadRequest)
		return
	}

	usernameId := cookie.Value
	idStr := strings.Split(usernameId, ":")[1]
	if len(idStr) <= 0 {
		http.Error(w, "Erreur récupération ID dans cookie", http.StatusBadRequest)
	}
	id, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		http.Error(w, "Erreur conversion ID", http.StatusBadRequest)
	}

	reqData := dto.DtoThread{
		Name:         r.FormValue("title"),
		Content:      r.FormValue("content"),
		FkIdUser:     id,
		FkIdCategory: idCategory,
	}

	newId, newIdStatusCode, newIdErr := services_client.CreateThread(reqData)
	if newIdErr != nil || newIdStatusCode != http.StatusOK {
		http.Error(w, newIdErr.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/thread/%v", newId.IdThread), http.StatusSeeOther)
}
