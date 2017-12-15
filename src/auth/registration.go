package auth

import (
	"../lobby"
	"net/http"
	"html/template"
	"encoding/json"
)

type response struct  {
	Success bool	 `json:"success"`
	Error   string	 `json:"error"`
}

type message struct {
	Login string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Confirm string `json:"confirm_password"`
}

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/static/registration/registration.html")
		t.Execute(w, nil)
	}
	if r.Method == "POST" {
		r.ParseForm()
		// берем из формы вбитые значения
		login := r.Form.Get("username")
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		confirm := r.Form.Get("confirm_password")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if login == "" || email == "" || password == "" || confirm == "" {
			resp := response{Success: false, Error: "form is empty"}
			json.NewEncoder(w).Encode(resp)
		} else {
			if confirm == password {

				checkLogin := checkAvailableLogin(login)
				checkEmail := checkAvailableEmail(email)

				if checkLogin && checkEmail {
					//SuccessRegistration(login, email, password)
					resp := response{Success: true, Error: ""}
					json.NewEncoder(w).Encode(resp)
				} else {
					if !checkLogin {
						resp := response{Success: false, Error: "login busy"} // error "этот логин уже занят"
						json.NewEncoder(w).Encode(resp)
					}
					if !checkEmail {
						resp := response{Success: false, Error: "email busy"}
						json.NewEncoder(w).Encode(resp)
					}
				}
			} else {
				resp := response{Success: false, Error: "password error"}
				json.NewEncoder(w).Encode(resp)
			}
		}
	}
}

func checkAvailableLogin(login string)(checkLogin bool)  {
	// TODO неверно делает сравнение

	user := lobby.GetUsers("WHERE name='" + login + "'")

	if user.Name != "" {
		checkLogin = false
	} else {
		checkLogin = true
	}

	return
}

func checkAvailableEmail(email string)(checkEmail bool)  {

	user := lobby.GetUsers("WHERE mail='" + email + "'")

	if user.Mail != "" {
		checkEmail = false
	} else {
		checkEmail = true
	}

	return
}

func SuccessRegistration(login, email, password string)  {
	lobby.CreateUser(login, email, password)
}