function UserRegistration() {
    var form = document.getElementById("formNewUser");
    var formData = new FormData(form);

    fetch('http://' + window.location.host + '/registration',
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            credentials: "same-origin",
            body: JSON.stringify({
                username: formData.get("username"),
                email: formData.get("email"),
                password: formData.get("password"),
                confirm: formData.get("confirm")
            })
        }).then(function(response) {
        // Стоит проверить код ответа.
        if (!response.ok) {
            // Сервер вернул код ответа за границами диапазона [200, 299]
            return Promise.reject(new Error(
                'Response failed: ' + response.status + ' (' + response.statusText + ')'
            ));
        }

        // Далее будем использовать только JSON из тела ответа.
        return response.json();
    }).then(function(data) {
        ReadResponse(data, formData)
    }).catch(function(error) {
        console.log(error)
    });
}

function ReadResponse(response, formData) {
    console.log(response.success + " " + response.error);
    if (response.success) {
        fetch('http://' + window.location.host + '/login',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                credentials: "same-origin",
                body: JSON.stringify({
                    username: formData.get("username"),
                    password: formData.get("password")
                })
            }).then(function(response) {
            if (!response.ok) {
                return Promise.reject(new Error(
                    'Response failed: ' + response.status + ' (' + response.statusText + ')'
                ));
            }
            return response.json();
        }).then(function() {
            var button = document.getElementById("regButton");
            button.value = "Перейти в лобби";
            button.onclick = function () {
                location.href = "../../lobby"
            }
        });
    } else {
        var errorDiv = document.getElementById("error");

        if (response.error === "form is empty") {
            errorDiv.innerHTML = "Не все формы заполнены"
        }
        if (response.error === "login busy") {
            errorDiv.innerHTML = "Этот логин уже занят"
        }
        if (response.error === "email busy") {
            errorDiv.innerHTML = "Эта почта уже используется"
        }
        if (response.error === "password error") {
            errorDiv.innerHTML = "Пароли не совпадают"
        }
    }
}