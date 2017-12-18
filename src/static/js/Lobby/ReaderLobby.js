function ReaderLobby(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var func;
    var text;
    var textButton;
    var funcButton;

    if (event === "InitLobby") {
        var login = document.getElementById('login');
        var userName = JSON.parse(jsonMessage).user_name;
        login.innerHTML = "Вы зашли как: " + userName;
    }

    if (event === "DisconnectLobby") {
        location.reload()
    }
    if (event === "GameRefresh") {
        DelElements("Select Menu");
    }
    if (event === "DelUser") {
        DelElements("User List");
    }

    if (event === "UserRefresh" || event === "JoinToLobby") {

        var user = new Object();
        user.Name = JSON.parse(jsonMessage).game_user;
        user.Ready = JSON.parse(jsonMessage).ready;
        user.Respawn = JSON.parse(jsonMessage).respawn_name;

        CreateLobbyLine('User', 'gameInfo', null, null, null, null, null, user, JSON.parse(jsonMessage).user_name);

    }

    if (event === "GameView") {
        GameView(jsonMessage)
    }

    if (event === "DontEndGamesList") {
        NotEndGame(jsonMessage)
    }

    if (event === "MapView") {
        MapView(jsonMessage);
    }

    if (event === "NewUser") {

        var user = new Object();
        user.Name = JSON.parse(jsonMessage).new_user;
        user.Ready = " Не готов";
        user.Respawn = JSON.parse(jsonMessage).respawn_name;

        CreateLobbyLine('User', 'gameInfo', null, null, null, null, null, user, JSON.parse(jsonMessage).user_name);
    }

    if (event === "CreateLobbyGame") {

        textButton = "Начать";
        funcButton = CreateNewGame;

        CreateLobbyMenu(textButton, funcButton, JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, true);

        var user = new Object();
        user.Name = JSON.parse(jsonMessage).user_name;
        user.Ready = " Не готов";
        user.Respawn = JSON.parse(jsonMessage).respawn_name;

        CreateLobbyLine('User', 'gameInfo', null, null, null, null, null, user, JSON.parse(jsonMessage).user_name);

        Respawn();

    }

    if (event === "initLobbyGame") {
        textButton = "Готов";
        funcButton = function () {
            sendReady(this.id)
        };
        CreateLobbyMenu(textButton, funcButton, JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, false);
        if (JSON.parse(jsonMessage).error === "") {
            text = JSON.parse(jsonMessage).user_name + " Не готов";
            CreateLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).user_name, null, null, null, text, JSON.parse(jsonMessage).user_name);
            Respawn();
        }
    }

    if (event === "Respawn") {
        //Create and append the options
        var select = document.getElementById("RespawnSelect");
        if (select) {
            var option = document.createElement("option");
            option.className = "RespawnOption";
            option.value = JSON.parse(jsonMessage).respawn;
            option.text = JSON.parse(jsonMessage).respawn_name;
            select.appendChild(option);
        }
    }

    if (event === "StartNewGame") {
        if (JSON.parse(jsonMessage).error === "") {
            toField = true;
            var idGame = JSON.parse(jsonMessage).id_game;
            document.cookie = "idGame=" + idGame + "; path=/;";
            location.href = "http://" + window.location.host + "/field";
        } else {
            if (JSON.parse(jsonMessage).error === "Players < 2") {
                alert("Ошибка: Мало игроков для старта");
            }
            if (JSON.parse(jsonMessage).error === "error ad to DB") {
                alert("Неизвестная ошибка");
            }
            if (JSON.parse(jsonMessage).error === "PlayerNotReady") {
                alert("Ошибка: не все игроки готовы");
            }
        }
    }

    if (event === "Ready") {
        var error = JSON.parse(jsonMessage).error;
        var user = new Object();

        if (error === "") {

            user.Name = JSON.parse(jsonMessage).game_user;
            user.Ready = JSON.parse(jsonMessage).ready;
            user.Respawn = JSON.parse(jsonMessage).respawn_name;

            CreateLobbyLine('User', 'gameInfo', null, null, null, null, null, user, JSON.parse(jsonMessage).user_name);

            if (user === JSON.parse(jsonMessage).user_name && JSON.parse(jsonMessage).respawn_name === "") {
               //CreateSelectRespawn(user, user + " Не готов ");
            }
        } else {
            //CreateSelectRespawn(user, user + " Не готов ");
        }
        Respawn();
    }
}

function CreateInLobbyLine(gameContent, className, id, func, funcMouse, funcOutMouse, text, owned) {
    var list = document.getElementById(gameContent);
    var div = document.createElement('div');
    div.style.wordWrap = 'break-word';
    div.className = className;
    div.id = id;
    div.onclick = func;
    div.onmouseover = funcMouse;
    div.onmouseout = funcOutMouse;
    div.appendChild(document.createTextNode(text));
    list.appendChild(div);

    if (id === owned) {
        CreateSelectRespawn(id, text)
    }
}