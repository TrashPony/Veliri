var stompClient = null;
var createGame = false;
var createNameGame = "";
var toField = false;
var sock;

function ConnectLobby() {
     sock = new WebSocket("ws://" + window.location.host + "/wsLobby");
     console.log("Websocket - status: " + sock.readyState);

     sock.onopen = function(msg) {
         console.log("CONNECTION opened..." + this.readyState);
     };
     sock.onmessage = function(msg) {
         console.log("message: " + msg.data);
         ResponseLobby(msg.data);
     };
     sock.onerror = function(msg) {
         console.log("Error occured sending..." + msg.data);
     };
     sock.onclose = function(msg) {
        console.log("Disconnected - status " + this.readyState);
        if(!toField) {
            location.href = "http://642e0559eb9c.sn.mynetname.net:8080/login";
        }
     };
}

function ResponseLobby(jsonMessage) {
    if (JSON.parse(jsonMessage).event) {
        var event = JSON.parse(jsonMessage).event;

        if (event === "MapView") {
            var SelectMap = document.getElementsByClassName("Select Map");
            while (SelectMap.length > 0) {
                SelectMap[0].parentNode.removeChild(SelectMap[0]);
            }

            var mapName = (JSON.parse(jsonMessage).response_name_map).split(':');
            var mapContent = document.getElementById('Games');

            var p = document.createElement('p');
            p.style.wordWrap = 'break-word';
            p.appendChild(document.createTextNode("Карты:"));
            p.className = "Select Map";
            mapContent.appendChild(p);

            for (var i = 0; (i + 1) < mapName.length; i++) {
                div = document.createElement('div');
                div.style.wordWrap = 'break-word';
                div.className = "Select Map";
                div.id = mapName[i];
                div.onclick = function () {
                    CreateLobbyGame(this.id);
                };
                div.appendChild(document.createTextNode(i + ") " + mapName[i]));
                mapContent.appendChild(div);
            }
        }

        if (event === "GameView") {
            var SelectGame = document.getElementsByClassName("GameView");
            while (SelectGame.length > 0) {
                SelectGame[0].parentNode.removeChild(SelectGame[0]);
            }

            var gameName = (JSON.parse(jsonMessage).response_name_game).split(':');
            var mapName = (JSON.parse(jsonMessage).response_name_map).split(':');
            var userName = (JSON.parse(jsonMessage).response_name_user).split(':');
            var gameContent = document.getElementById('Games list');

            for (var i = 0; (i + 1) < gameName.length; i++) {
                div = document.createElement('div');
                div.style.wordWrap = 'break-word';
                div.className = "Select Game";
                div.id = gameName[i];
                div.onclick = function () {
                    JoinToLobbyGame(this.id);
                };
                div.appendChild(document.createTextNode(i + ") Имя:     " + gameName[i] + " Карта:      " + mapName[i] + " Создатель:       " + userName[i]));
                gameContent.appendChild(div);
            }
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
            var SelectGame = document.getElementsByClassName("Select Game");
            while (SelectGame.length > 0) {
                SelectGame[0].parentNode.removeChild(SelectGame[0]);
            }

            var gameNames = (JSON.parse(jsonMessage).response_name_game).split(':');
            var gamesContent = document.getElementById('DontEndGame');

            var p = document.createElement('p');
            p.style.wordWrap = 'break-word';
            p.appendChild(document.createTextNode("Недоиграные игры:"));
            p.className = "Select Game";
            gamesContent.appendChild(p);

            for (var i = 0; (i + 1) < gameNames.length; i++) {
                div = document.createElement('div');
                div.style.wordWrap = 'break-word';
                div.className = "Select Game";
                div.id = gameNames[i];
                div.onclick = function () {
                    JoinToGame(this.id);
                };
                div.appendChild(document.createTextNode(i + ") " + gameNames[i]));
                gamesContent.appendChild(div);
            }
        }

        if (event === "Joiner") {
            var users = (JSON.parse(jsonMessage).response_name_user_2).split(':');
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

            var user = JSON.parse(jsonMessage).response_name_user;
            div = document.createElement('div');
            div.style.wordWrap = 'break-word';
            div.className = "Select Game";
            div.id = user;
            div.appendChild(document.createTextNode(user));
            parentElem[0].appendChild(div);
        }

        if (event === "StartNewGame") {
            toField = true;
            location.href = "http://" + window.location.host + "/field";
        }
    }
}

function ReturnLobby() {
    location.href = "http:/" + window.location.host + "/login";
}

function CreateLobbyGame(mapName) {
    var gameName = document.querySelector('input[name="NameGame"]').value;
    sendCreateLobbyGame(mapName, gameName);
}

function CreateNewGame() {
    if(createNameGame !== "") {
        sendStartNewGame(createNameGame);
    } else {
        location.href = "http:/" + window.location.host + "/login";
    }
}

function JoinToLobbyGame(gameName) {
    sendJoinToLobbyGame(gameName);
}

function JoinToGame(gameName) {
    sendJoinToGame(gameName)
}

function sendJoinToGame(gameName) {
    sock.send(JSON.stringify({
        event: "JoinToGame",
        game_name: gameName
    }));
}
function sendMapSelection() {
    sock.send(JSON.stringify({
            event: "MapView"
    }));
}

function sendGameSelection() {
    sock.send(JSON.stringify({
            event: "GameView"
    }));
}

function sendCreateLobbyGame(mapName, gameName) {
    createNameGame = gameName;
    sock.send(JSON.stringify({
        event: "CreateLobbyGame",
        map_name: mapName,
        game_name: gameName
    }));
}

function sendJoinToLobbyGame(gameName) {
    sock.send(JSON.stringify({
        event: "JoinToLobbyGame",
        game_name: gameName
    }));
}

function sendDontEndGamesList () {
    sock.send(JSON.stringify({
        event: "DontEndGamesList"
    }));
}

function sendStartNewGame (gameName) {
    sock.send(JSON.stringify({
        event: "StartNewGame",
        game_name: gameName
    }));
}