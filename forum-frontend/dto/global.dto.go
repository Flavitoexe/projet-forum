package dto

import "bluedit/forum-backend/models"

// DTOs Threads
type TempThreadsData struct {
	Data           *[]models.Thread
	DataStatusCode int
	DataErr        error
}

type TempThreadData struct {
	Data           *DtoThread
	DataStatusCode int
	DataErr        error
}

type DtoThread struct {
	IdThread     int     `json:"id_thread"`
	Name         string  `json:"name"`
	Content      string  `json:"content"`
	CreatedAt    string  `json:"created_at"`
	MessageCount int     `json:"message_count"`
	LikeCount    int     `json:"like_count"`
	CategoryName string  `json:"category_name"`
	FkIdCategory int     `json:"fk_id_category"`
	Author       *string `json:"author"`
	FkIdUser     int     `json:"fk_id_user"`
	Status       int     `json:"status"`
}

// DTOs Messages
type DtoMessage struct {
	IdMessage  int     `json:"id_message"`
	Content    string  `json:"content"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	LikeCount  int     `json:"likes"`
	Author     *string `json:"author"`
	FkIdUser   int     `json:"fk_id_user"`
	FkIdThread int     `json:"fk_id_thread"`
}

type DtoMessageRequest struct {
	IdMessage int    `json:"id_message"`
	Content   string `json:"content"`
}

type DtoThreadRequest struct {
	IdThread     int    `json:"id_thread"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	FkIdCategory int    `json:"fk_id_category"`
}

type TempMessagesData struct {
	Data           *[]DtoMessage
	DataStatusCode int
	DataErr        error
}

// DTOs Catégories
type TempCategoriesData struct {
	Data           *[]models.Category
	DataStatusCode int
	DataErr        error
}

type TempCategoryData struct {
	Data           *models.Category
	DataStatusCode int
	DataErr        error
}

// DTO utilisateurs
type TempUserData struct {
	Data           *DtoUser
	DataStatusCode int
	DataErr        error
}

type DtoUser struct {
	IdUser int    `json:"id_user"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

// DTOs erreurs
type DtoError struct {
	Code           int
	Message        string
	CategoriesData TempCategoriesData
}
