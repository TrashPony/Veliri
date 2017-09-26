var stompClient = null;
var createGame = false;
var sock;

function ConnectLobby() {
     sock = new WebSocket("ws://" + window.location.host + "/wsLobby");
     console.log("Websocket - status: " + sock.readyState);

     sock.onopen = function(msg) {
         console.log("CONNECTION opened..." + this.readyState);
     }

     sock.onmessage = function(msg) {
         console.log("message: " + msg.data);
         ResponseLobby(msg.data);
     }

     sock.onerror = function(msg) {
         console.log("Error occured sending..." + msg.data);
     }

     sock.onclose = function(msg) {
        console.log("Disconnected - status " + this.readyState);
        location.href = "http://642e0559eb9c.sn.mynetname.net:8080/login"
     }
}

function ResponseLobby(jsonMessage) {
    if(JSON.parse(jsonMessage).event) {
        var event = JSON.parse(jsonMessage).event;

        if (event === "MapSelection") {
            if (JSON.parse(jsonMessage).response_name_map) {
                var SelectMap = document.getElementsByClassName("SelectMap");
                while (SelectMap.length > 0) {
                    SelectMap[0].parentNode.removeChild(SelectMap[0]);
                }

                var mapName = (JSON.parse(jsonMessage).response_name_map).split(':');
                var mapContent = document.getElementById('Games');

                var p = document.createElement('p');
                p.style.wordWrap = 'break-word';
                p.appendChild(document.createTextNode("Карты:"));
                mapContent.appendChild(p);

                for (var i = 0; (i + 1) < mapName.length; i++) {
                    div = document.createElement('div');
                    div.style.wordWrap = 'break-word';
                    div.className = "SelectMap";
                    div.id = mapName[i];
                    div.onclick = function () {
                        CreateNewGame(this.id);
                    };
                    div.appendChild(document.createTextNode(i + ") " + mapName[i]));
                    mapContent.appendChild(div);
                }
            }
        }

        if (event === "GameSelection") {
            //if (JSON.parse(jsonMessage).gameList) {
                var SelectGame = document.getElementsByClassName("Games");
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
                    div.className = "Games";
                    div.id = gameName[i];
                    div.onclick = function () {
                        CreateNewGame(this.id);
                    };
                    div.appendChild(document.createTextNode(i + ") Имя:     " + gameName[i] + " Карта:      " + mapName[i] + " Создатель:       " + userName[i]));
                    gameContent.appendChild(div);
                }
            //}
        }

        if (event === "CreateNewGame") {
                // удаляем старые элементы //
                del = document.getElementById("lobby");
                del.remove();
                // удаляем старые элементы //

                var div = document.createElement('div');
                div.className = "gameInfo";
                var parentElem = document.body;
                parentElem.appendChild(div);
                var button = document.createElement("input");
                button.type = "button";
                button.value = "Отменить";
                div.appendChild(button);
                createGame = true;

        }
    }
}

function CreateNewGame(mapName) {
    var gameName = document.querySelector('input[name="NameGame"]').value;
    sendCrateNewGame(mapName, gameName);
}

function sendMapSelection() {
    sock.send(JSON.stringify({
            event: "MapSelection"
    }));
}

function sendGameSelection() {
    sock.send(JSON.stringify({
            event: "GameSelection"
    }));
}

function sendCrateNewGame(mapName, gameName) {
    sock.send(JSON.stringify({
        event: "CreateNewGame",
        map_name: mapName,
        game_name: gameName
    }));
}