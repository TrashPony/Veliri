function ReaderLobby(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var user;
    var func;
    var text;
    var textButton;
    var funcButton;
    var funcMouse;
    var funcOutMouse;

    if (event === "InitLobby") {
        var login = document.getElementById('login');
        var userName = JSON.parse(jsonMessage).user_name;
        login.innerHTML = "Вы зашли как: " + userName;
    }

    if (event === "DisconnectLobby") {
        location.reload()
    }
    if (event === "GameRefresh") {
        DelElements("Select Game");
    }
    if (event === "DelUser"){
        DelElements("User List");
    }

    if (event === "UserRefresh" || event === "JoinToLobby") {
        if (JSON.parse(jsonMessage).ready === "true") {
            text = JSON.parse(jsonMessage).game_user + " Готов!";
        } else {
            text = JSON.parse(jsonMessage).game_user + " Не готов";
        }
        func = null;
        CreateLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).game_user, func, null, null, text);
    }

    if (event === "GameView") {
        func = function () { JoinToLobbyGame(this.id); };
        text = "Имя: " + JSON.parse(jsonMessage).name_game + ", Карта: " + JSON.parse(jsonMessage).name_map + ", Создатель: " +
            JSON.parse(jsonMessage).creator + ", Игроков:" + JSON.parse(jsonMessage).players + "/" + JSON.parse(jsonMessage).num_of_players;
        CreateLobbyLine('Games list', 'Select Game', JSON.parse(jsonMessage).name_game, func, null, null, text);
    }

    if (event === "DontEndGamesList") {
        func = function () { JoinToGame(this.id); };
        text = "Имя: " + JSON.parse(jsonMessage).name_game + " id: " + JSON.parse(jsonMessage).id_game + " Шаг: " +
            JSON.parse(jsonMessage).step_game + " Фаза: " + JSON.parse(jsonMessage).phase_game + " Мой ход: " + (!JSON.parse(jsonMessage).ready);
        CreateLobbyLine('DontEndGames', 'Select DontEndGame', JSON.parse(jsonMessage).id_game, func, null, null, text);
    }

    if (event === "MapView") {
        func = function () {
            CreateLobbyGame(this.id);
        };
        funcMouse = function () {
            MouseOverMap(this.id);
        };
        funcOutMouse = function () {
            MouseOutMap()
        };
        text = "Имя: " + JSON.parse(jsonMessage).name_map + " Максимум игроков:" + JSON.parse(jsonMessage).num_of_players;
        CreateLobbyLine('Maps', 'Select Map', JSON.parse(jsonMessage).name_map, func, funcMouse, funcOutMouse, text);
    }

    if (event === "NewUser") {
        func = null;
        text = JSON.parse(jsonMessage).new_user + " Не готов";
        CreateLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).new_user, func, null, null, text);
    }

    if (event === "CreateLobbyGame") {
        textButton = "Начать";
        funcButton = CreateNewGame;
        CreateLobbyMenu(textButton, funcButton, null, JSON.parse(jsonMessage).error);
    }

    if (event === "initLobbyGame") {
        textButton = "Готов";
        funcButton = function () {
            sendReady(this.id)
        };
        CreateLobbyMenu(textButton, funcButton, JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error);
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

    if (event === "Ready"){
        var ready;
        if (JSON.parse(jsonMessage).ready === "true") {
            ready = "Готов!"
        } else {
            ready = "Не готов"
        }
        user = JSON.parse(jsonMessage).game_user;
        var userBlock = document.getElementById(user);
        userBlock.innerHTML = user + " " + ready;
    }
}

function CreateLobbyMenu(textButton, funcButton, id, error) {
    if (error === "") {
        DelElements("NotGameLobby");
        var gameInfo = document.createElement('div');
        gameInfo.className = "gameInfo";
        gameInfo.id = "gameInfo";
        var parentElem = document.getElementById("lobby");
        parentElem.appendChild(gameInfo);
        var cancel = document.createElement("input");
        cancel.type = "button";
        cancel.value = "Отменить";
        cancel.onclick = ReturnLobby;
        gameInfo.appendChild(cancel);
        var button = document.createElement("input");
        button.type = "button";
        button.value = textButton;
        button.onclick = funcButton;
        button.id = id;
        gameInfo.appendChild(button);
        createGame = true;
        var parentElemDiv = document.getElementsByClassName("gameInfo");
        var div3 = document.createElement('div');
        div3.className = "User";
        div3.appendChild(document.createTextNode("Подключенные Игроки"));
        parentElemDiv[0].appendChild(div3);
    } else {
        if (error === "lobby is full") {
            alert("Игра полная")
        }
        if (error === "unknown error") {
            alert("unknown error")
        }
    }
}

function CreateLobbyLine(gameContent, className, id, func, funcMouse, funcOutMouse, text) {
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
}
