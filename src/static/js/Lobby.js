var stompClient = null;
var createGame = false;
var createNameGame = "";
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
         location.href = "http://642e0559eb9c.sn.mynetname.net:8080/login";
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
                    createNameGame = this.id;
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
                    JoinToGame(this.id);
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
            createGame = true;

        }


        if (event === "DontEndGames") {
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
    }
}

function ReturnLobby() {
    location.href = "http://642e0559eb9c.sn.mynetname.net:8080/login";
}

function CreateLobbyGame(mapName) {
    var gameName = document.querySelector('input[name="NameGame"]').value;
    sendCreateLobbyGame(mapName, gameName);
}

function CreateNewGame() {
    if(createNameGame !== "") {
        sendStartNewGame(createNameGame);
    } else {
        location.href = "http://642e0559eb9c.sn.mynetname.net:8080/login";
    }
}

function JoinToGame(gameName) {
    sendJoinToGame(gameName);
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
    sock.send(JSON.stringify({
        event: "CreateLobbyGame",
        map_name: mapName,
        game_name: gameName
    }));
}

function sendJoinToGame(gameName) {
    sock.send(JSON.stringify({
        event: "JoinToGame",
        game_name: gameName
    }));
}

function sendDontEndGames () {
    sock.send(JSON.stringify({
        event: "DontEndGames"
    }));
}

function sendStartNewGame () {
    sock.send(JSON.stringify({
        event: "StartNewGame"
    }));
}