package auth

import (
	"net/http"
	"html/template"
	"sessions-master"
	"encoding/gob"
)

var cookieStore = sessions.NewCookieStore([]byte("dick, mountain, sky ray")) // мало понимаю в шифрование сессии внутри указан приватный ключь шифрования

const cookieName = "MyCookie" // имя куки в браузере юзера

type sesKey int // -

const (
	login sesKey = iota // -
	id sesKey = iota // -
)

func Login(w http.ResponseWriter, r *http.Request) {
	gob.Register(sesKey(0)) // вот это отвечает за шифрование даных как я понял и это важно будет переделать вероятно когда то
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/static/login/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// берем из формы вбитые значения
		var userName string = r.Form.Get("username")
		var password string = r.Form.Get("password")
		// отправляет эти данные на проверку если прошло то возвращает пользователя и пропуск
		user, passed := CheckUserInfo(userName, password)
		if passed {
			//отправляет пользователя на получение токена подключения
			GetCookie(w , r, user)
		}
	}
}

func CheckUserInfo(userName string, userPassword string) (User, bool) {
	// тестовые in memory пользователи
	var users []User
	users = append(users, User{1,"admin", "pass1"})
	users = append(users, User{2,"user", "pass"})

	// сравнием вбитые значения со значениями на сервере
	for i := 0; i < len(users); i++{
		if users[i].name == userName && users[i].pass == userPassword{
			return users[i], true
		}
	}
	return users[0], false
}

func GetCookie(w http.ResponseWriter, r *http.Request, user User) {
	// берет сеанс из браузера пользователя
	ses, err := cookieStore.Get(r, cookieName)
	// если есть куки подписаные не правильным ключем то вылетает ошибка
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ses.Values[login] = user.name // ложит данные в сессию
	ses.Values[id] = user.id // ложит данные в сессию


	//возвращает ответ с сохранение сессии в браузере
	err = cookieStore.Save(r, w, ses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func CheckCookie(w http.ResponseWriter, r *http.Request) (string, int) {
	// берет сеанс из браузера пользователя
	ses, err := cookieStore.Get(r, cookieName)
	// если есть куки подписаные не правильным ключем то вылетает ошибка
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", 0
	}

	// смотрит пустая сессия или нет, если нет то присваивает переменной логин логин
	// ищет значение в сессии и присваивает переменной [Login] - ключь .(string) - тип данных ok - удалось ли получить

	login, ok := ses.Values[login].(string)
	id, ok := ses.Values[id].(int)

	if !ok { // если пустая то говорит что ты анонимус
		login = "anonymous"
	}

	// выводит в браузер кто ты есть
	//w.Write([]byte("Твой логин " + login + ", твой id " + strconv.Itoa(id)))
	return login, id
}

type User struct {
	id int
	name string
	pass string
}



