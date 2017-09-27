package webSocket

import (
	"net/http"
	"../auth"
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	var login string
	var id int

	login, id = auth.CheckCookie(w, r) // берем из куки данные по логину и ид пользовтеля

	if login == "" || id == 0  || login == "anonymous" {
		println("Соеденение не разрешено: не авторизован")
		//http.Redirect(w, r, "http://www.google.com", 401)
		return // если человек не авторизован то ему не разрешается соеденение
	}
	println(r.URL.Path)
	if r.URL.Path == "/wsLobby" {
		ReadLobbySocket(login, id, w, r)
	}

	if r.URL.Path == "/wsField" {
		ReadFieldSocket(login, id, w, r)
	}
}
