function UserLogin() {
    var form = document.getElementById("login");
    var formData = new FormData(form);

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
        ReadResponse(data)
    }).catch(function(error) {
        console.log(error)
    });
}

function ReadResponse(response) {
    console.log(response.success + " " + response.error);
    if (response.success) {
        location.href = "../../lobby"
    } else {
        var errorDiv = document.getElementById("error");

        if (response.error === "not allow") {
            errorDiv.innerHTML = "Не верный логин либо пароль"
        }
    }
}