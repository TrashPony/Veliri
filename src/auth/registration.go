package auth

import (
	"../lobby"
	"net/http"
	"html/template"
	"encoding/json"
	"errors"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/static/registration/registration.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// берем из формы вбитые значения
		login	 := r.Form.Get("username")
		password := r.Form.Get("password")
		confirm  := r.Form.Get("confirm_password")
		email    := r.Form.Get("email")

		if confirm == password{

			checkLogin := checkAvailableLogin(login)
			checkEmail := checkAvailableEmail(email)

			if checkLogin && checkEmail {
				SuccessRegistration(login, email, password)
			} else {
				if !checkLogin {
					err := errors.New("login busy") // error "этот логин уже занят"
					json.NewEncoder(w).Encode(err)
				}
				if !checkEmail {
					err := errors.New("email busy") //error "этот e-mail уже занят"
					json.NewEncoder(w).Encode(err)
				}
			}
		} else {
			err := errors.New("password error") //error "пароли не совпадают"
			json.NewEncoder(w).Encode(err)
		}
	}
}

func checkAvailableLogin(login string)(checkLogin bool)  {

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