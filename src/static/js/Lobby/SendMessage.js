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

function Respawn() {
    DelElements("RespawnOption");
    sock.send(JSON.stringify({
        event: "Respawn"
    }));
}

function sendDontEndGamesList () {
    DelElements("Select.SubMenu");

    sock.send(JSON.stringify({
        event: "DontEndGamesList"
    }));
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

function sendStartNewGame (gameName) {
    sock.send(JSON.stringify({
        event: "StartNewGame",
        game_name: gameName
    }));
}


function sendGameSelection() {
    sock.send(JSON.stringify({
        event: "GameView"
    }));
}

