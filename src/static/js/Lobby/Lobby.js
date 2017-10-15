var stompClient = null;
var createGame = false;
var createNameGame = "";
var toField = false;
var sock;

function ConnectLobby() {
     sock = new WebSocket("ws://" + window.location.host + "/wsLobby");
     console.log("Websocket - status: " + sock.readyState);

    var date = new Date(0);
    document.cookie = "idGame=; path=/; expires=" + date.toUTCString();

     sock.onopen = function(msg) {
         console.log("CONNECTION opened..." + this.readyState);
     };
     sock.onmessage = function(msg) {
         console.log("message: " + msg.data);
         ReaderLobby(msg.data);
     };
     sock.onerror = function(msg) {
         console.log("Error occured sending..." + msg.data);
     };
     sock.onclose = function(msg) {
        console.log("Disconnected - status " + this.readyState);
        if(!toField) {
            location.href = "../../login";
        }
     };
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
        location.href = "../../login";
    }
}

function JoinToLobbyGame(gameName) {
    // удаляем старые элементы //
    var del = document.getElementById("lobby");
    del.remove();
    // удаляем старые элементы //

    var div2 = document.createElement('div');
    div2.className = "gameInfo";
    var parentElem = document.body;
    parentElem.appendChild(div2);
    var cancel = document.createElement("input");
    cancel.type = "button";
    cancel.value = "Отменить";
    cancel.onclick = ReturnLobby;
    div2.appendChild(cancel);
    var tick = document.createElement("input");
    tick.type = "button";
    tick.value = "Готов";
    tick.id = gameName;
    tick.onclick = function () {
        sendReady(this.id)
    };
    div2.appendChild(tick);

    createGame = true;
    var parentElemDiv = document.getElementsByClassName("gameInfo");

    var div3 = document.createElement('div');
    div3.className = "User";
    div3.appendChild(document.createTextNode("Подключенные Игроки"));
    parentElemDiv[0].appendChild(div3);

    sendJoinToLobbyGame(gameName);
}

function JoinToGame(idGame) {
    toField = true;
    document.cookie = "idGame=" + idGame + "; path=/;";
    location.href = "http://" + window.location.host + "/field";
}

function sendMapSelection() {

    var SelectMap = document.getElementsByClassName("Select Map");
    while (SelectMap.length > 0) {
        SelectMap[0].parentNode.removeChild(SelectMap[0]);
    }

    var mapContent = document.getElementById('Games');
    var p = document.createElement('p');
    p.style.wordWrap = 'break-word';
    p.appendChild(document.createTextNode("Карты:"));
    p.className = "Select Map";
    mapContent.appendChild(p);

    sock.send(JSON.stringify({
            event: "MapView"
    }));
}

function sendGameSelection() {
    var SelectGame = document.getElementsByClassName("GameView");
    while (SelectGame.length > 0) {
        SelectGame[0].parentNode.removeChild(SelectGame[0]);
    }

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
    var SelectGame = document.getElementsByClassName("Select Game");
    while (SelectGame.length > 0) {
        SelectGame[0].parentNode.removeChild(SelectGame[0]);
    }

    var gamesContent = document.getElementById('DontEndGame');

    var p = document.createElement('p');
    p.style.wordWrap = 'break-word';
    p.appendChild(document.createTextNode("Недоиграные игры:"));
    gamesContent.appendChild(p);

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

function sendReady (gameName) {
    sock.send(JSON.stringify({
        event: "Ready",
        game_name: gameName
    }));
}
