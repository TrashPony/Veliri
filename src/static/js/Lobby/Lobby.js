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
    p.className = "Maps";
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
