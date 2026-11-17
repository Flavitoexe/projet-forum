package models

import "time"

type User struct {
	IdUser    int       `json:"id_user"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
	IsBanned  bool      `json:"is_banned"`
}
