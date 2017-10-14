function ReaderLobby(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;

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
        del = document.getElementById("lobby");
        del.remove();
        // удаляем старые элементы //

        var div = document.createElement('div');
        div.className = "gameInfo";
        var parentElem = document.body;
        parentElem.appendChild(div);
        var button1 = document.createElement("input");
        button1.type = "button";
        button1.value = "Отменить";
        button1.onclick = ReturnLobby;
        div.appendChild(button1);
        var button2 = document.createElement("input");
        button2.type = "button";
        button2.value = "Начать";
        button2.onclick = CreateNewGame;
        div.appendChild(button2);
        var div3 = document.createElement('div');
        div3.appendChild(document.createTextNode("Игроки"));
        createGame = true;

        var div3 = document.createElement('div');
        div3.className = "Select Game";
        div3.appendChild(document.createTextNode("Подключенные Игроки"));
        div.appendChild(div3);
    }

    if (event === "DontEndGamesList") {
        var gamesContent = document.getElementById('DontEndGame');

        div = document.createElement('div');
        div.style.wordWrap = 'break-word';
        div.className = "Select Game";
        div.id = JSON.parse(jsonMessage).name_game;
        div.onclick = function () {
            JoinToGame(this.id);
        };
        div.appendChild(document.createTextNode("Имя: " + JSON.parse(jsonMessage).name_game + " id: " + JSON.parse(jsonMessage).id_game + " Шаг: " +
            JSON.parse(jsonMessage).step_game + " Фаза: " + JSON.parse(jsonMessage).phase_game + " Мой ход: " + (!JSON.parse(jsonMessage).ready)));
        gamesContent.appendChild(div);

    }

    if (event === "Joiner") {
        var users = (JSON.parse(jsonMessage).name_user_2).split(':');
        // удаляем старые элементы //
        del = document.getElementById("lobby");
        del.remove();
        // удаляем старые элементы //

        var div2 = document.createElement('div');
        div2.className = "gameInfo";
        var parentElem = document.body;
        parentElem.appendChild(div2);
        var button1 = document.createElement("input");
        button1.type = "button";
        button1.value = "Отменить";
        //button1.onclick = ReturnLobby;
        div2.appendChild(button1);
        var button2 = document.createElement("input");
        button2.type = "button";
        button2.value = "Готов";  // пока никак не работает потом надо будет сделать)
        div2.appendChild(button2);

        createGame = true;
        var parentElemDiv = document.getElementsByClassName("gameInfo");

        var div3 = document.createElement('div');
        div3.className = "Select Game";
        div3.appendChild(document.createTextNode("Подключенные Игроки"));
        parentElemDiv[0].appendChild(div3);

        for (var i = 0; i < users.length; i++) {
            div = document.createElement('div');
            div.style.wordWrap = 'break-word';
            div.className = "Select Game";
            div.id = users[i];
            div.appendChild(document.createTextNode(i + ") " + users[i]));
            parentElemDiv[0].appendChild(div);
        }
    }
    if (event === "JoinToLobbyGame") {
        var parentElem = document.getElementsByClassName("gameInfo");

        var user = JSON.parse(jsonMessage).name_user;
        div = document.createElement('div');
        div.style.wordWrap = 'break-word';
        div.className = "Select Game";
        div.id = user;
        div.appendChild(document.createTextNode(user));
        parentElem[0].appendChild(div);
    }

    if (event === "StartNewGame") {
        toField = true;
        var idGame = JSON.parse(jsonMessage).name_map;
        document.cookie = "idGame=" + idGame + "; path=/;";
        location.href = "http://" + window.location.host + "/field";
    }
}