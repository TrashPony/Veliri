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
	Confirm string `json:"confirm"`
}

func Registration(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/static/registration/index.html")
		t.Execute(w, nil)
	}
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var msg message
		err := decoder.Decode(&msg)
		if err != nil {
			panic(err)
		}

		if msg.Login == "" || msg.Email == "" || msg.Password == "" || msg.Confirm == "" {
			resp := response{Success: false, Error: "form is empty"}
			json.NewEncoder(w).Encode(resp)
		} else {
			if msg.Password == msg.Confirm {

				checkLogin := checkAvailableLogin(msg.Login)
				checkEmail := checkAvailableEmail(msg.Email)

				if checkLogin && checkEmail {
					SuccessRegistration(msg.Login, msg.Email, msg.Password)
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
	hashPassword, err := HashPassword(login, password)
	if err != nil {
		panic(err)
	}
	lobby.CreateUser(login, email, hashPassword)
}