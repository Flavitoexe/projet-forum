package controllers_client

import "net/http"

func SetSessionCookie(w http.ResponseWriter, username string) {
	http.SetCookie(w, &http.Cookie{Name: "bluedit_session",
		Value:    username,
		Path:     "/",
		MaxAge:   60 * 60 * 4,
		HttpOnly: true,
	})
}

func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: "bluedit_session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}
