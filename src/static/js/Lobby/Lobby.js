var stompClient = null;
var createGame = false;
var createNameGame = "";
var toField = false;
var sock;
var respownId;

function ConnectLobby() {
     sock = new WebSocket("ws://" + window.location.host + "/wsLobby");
     console.log("Websocket - status: " + sock.readyState);

    var date = new Date(0);
    document.cookie = "idGame=; path=/; expires=" + date.toUTCString();

     sock.onopen = function(msg) {
         console.log("CONNECTION opened..." + this.readyState);
         InitLobby();
         sendGameSelection();
         sendDontEndGamesList();
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
    location.reload();
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
    DelElements("DontEndGame");
    DelElements("Select Map");
    DelElements("Maps");
    var mapContent = document.getElementById('Games');
    var div = document.createElement('div');
    div.style.wordWrap = 'break-word';
    div.appendChild(document.createTextNode("Карты:"));
    div.className = "Maps";
    div.id = "Maps";
    mapContent.appendChild(div);

    sock.send(JSON.stringify({
            event: "MapView"
    }));
}

function MouseOverMap(id) {
    var info = document.getElementById('SelectInfo');
    info.innerHTML = "Имя карты " + id;
    var div = document.createElement('div');
    div.style.wordWrap = 'break-word';
    div.appendChild(document.createTextNode("Невероятная картинка карты! В разработке"));
    div.className = "infoMap";
    div.id = "infoImage";
    info.appendChild(div);
    var div2 = document.createElement('div');
    div2.style.wordWrap = 'break-word';
    div2.className = "infoMap";
    div2.appendChild(document.createTextNode("Описание карты, в разработке"));
    info.appendChild(div2);
}

function MouseOutMap() {
    var info = document.getElementById('SelectInfo');
    info.innerHTML = "";
    DelElements("infoMap");
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
    DelElements("Select DontEndGame");

    var gamesContent = document.getElementById('DontEndGame');

    var div = document.createElement('div');
    div.style.wordWrap = 'break-word';
    div.appendChild(document.createTextNode("Недоиграные игры:"));
    div.className= "Games";
    div.id = "DontEndGames";
    gamesContent.appendChild(div);

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
    var selectResp = document.getElementById("RespawnSelect");
    if (selectResp) {
        DelElements("RespawnSelect");
        respownId = selectResp.value
    }
    if (selectResp) {
        sock.send(JSON.stringify({
            event: "Ready",
            game_name: gameName,
            respawn: selectResp.value
        }));
    } else {
        sock.send(JSON.stringify({
            event: "Ready",
            game_name: gameName,
            respawn: respownId
        }));
    }
}

function Logout() {
    sock.send(JSON.stringify({
        event: "Logout"
    }));
}

function InitLobby() {
    sock.send(JSON.stringify({
        event: "InitLobby"
    }));
}

function Respawn() {
    DelElements("RespawnOption");
    sock.send(JSON.stringify({
        event: "Respawn"
    }));
}

function DelElements(ClassElements) {
    var SelectMap = document.getElementsByClassName(ClassElements);
    while (SelectMap.length > 0) {
        SelectMap[0].parentNode.removeChild(SelectMap[0]);
    }
}