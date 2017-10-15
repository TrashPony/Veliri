function ReaderLobby(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var parentElem;
    var cancel;
    var Start;
    var Players;
    var user;

    if (event === "MapView") {

        var mapContent = document.getElementById('Games');
        div = document.createElement('div');
        div.style.wordWrap = 'break-word';
        div.className = "Select Map";
        div.id = JSON.parse(jsonMessage).name_map;
        div.onclick = function () {
            CreateLobbyGame(this.id);
        };
        div.appendChild(document.createTextNode(JSON.parse(jsonMessage).name_map));
        mapContent.appendChild(div);
    }

    if (event === "GameView") {
        var gameContent = document.getElementById('Games list');
        div = document.createElement('div');
        div.style.wordWrap = 'break-word';
        div.className = "Select Game";
        div.id = JSON.parse(jsonMessage).name_game;
        div.onclick = function () {
            JoinToLobbyGame(this.id);
        };
        div.appendChild(document.createTextNode("Имя: " + JSON.parse(jsonMessage).name_game + " Карта: " + JSON.parse(jsonMessage).name_map + " Создатель: " + JSON.parse(jsonMessage).creator));
        gameContent.appendChild(div);
    }

    if (event === "CreateLobbyGame") {
        // удаляем старые элементы //
        var del = document.getElementById("lobby");
        del.remove();
        // удаляем старые элементы //

        var div = document.createElement('div');
        div.className = "gameInfo";
        parentElem = document.body;
        parentElem.appendChild(div);
        cancel = document.createElement("input");
        cancel.type = "button";
        cancel.value = "Отменить";
        cancel.onclick = ReturnLobby;
        div.appendChild(cancel);
        Start = document.createElement("input");
        Start.type = "button";
        Start.value = "Начать";
        Start.onclick = CreateNewGame;
        div.appendChild(Start);
        Players = document.createElement('div');
        Players.appendChild(document.createTextNode("Игроки"));
        createGame = true;

        Players = document.createElement('div');
        Players.className = "User";
        Players.appendChild(document.createTextNode("Подключенные Игроки"));
        div.appendChild(Players);
    }

    if (event === "DontEndGamesList") {
        var gamesContent = document.getElementById('DontEndGame');

        div = document.createElement('div');
        div.style.wordWrap = 'break-word';
        div.className = "Select Game";
        div.id = JSON.parse(jsonMessage).id_game;
        div.onclick = function () {
            JoinToGame(this.id);
        };
        div.appendChild(document.createTextNode("Имя: " + JSON.parse(jsonMessage).name_game + " id: " + JSON.parse(jsonMessage).id_game + " Шаг: " +
            JSON.parse(jsonMessage).step_game + " Фаза: " + JSON.parse(jsonMessage).phase_game + " Мой ход: " + (!JSON.parse(jsonMessage).ready)));
        gamesContent.appendChild(div);

    }

    if (event === "JoinToLobby") {
        var gameInfo = document.getElementsByClassName("gameInfo");
        user = JSON.parse(jsonMessage).game_user;
        var ready;
        if (JSON.parse(jsonMessage).ready === "true") {
            ready = "Готов!"
        } else {
            ready = "Не готов"
        }
        div = document.createElement('div');
        div.style.wordWrap = 'break-word';
        div.className = "User List";
        div.id = user;
        div.appendChild(document.createTextNode(user + " " + ready));
        gameInfo[0].appendChild(div);
    }

    if (event === "NewUser") {
        parentElem = document.getElementsByClassName("gameInfo");
        user = JSON.parse(jsonMessage).new_user;
        div = document.createElement('div');
        div.style.wordWrap = 'break-word';
        div.className = "User List";
        div.id = user;
        div.appendChild(document.createTextNode(user+ " Не готов"));
        parentElem[0].appendChild(div);
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
