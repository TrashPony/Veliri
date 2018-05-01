function Respawn() {
    DelElements("RespawnOption");
    lobby.send(JSON.stringify({
        event: "Respawn"
    }));
}

function sendDontEndGamesList () {
    DelElements("Select.SubMenu");

    lobby.send(JSON.stringify({
        event: "DontEndGamesList"
    }));
}

function Logout() {
    lobby.send(JSON.stringify({
        event: "Logout"
    }));
}

function InitLobby() {
    lobby.send(JSON.stringify({
        event: "InitLobby"
    }));
}

function sendCreateLobbyGame(mapId, gameName) {
    createNameGame = gameName;
    lobby.send(JSON.stringify({
        event: "CreateLobbyGame",
        map_id: Number(mapId),
        game_name: gameName
    }));
}

function sendJoinToLobbyGame(gameName) {
    lobby.send(JSON.stringify({
        event: "JoinToLobbyGame",
        game_name: gameName
    }));
}

function sendStartNewGame (gameName) {
    lobby.send(JSON.stringify({
        event: "StartNewGame",
        game_name: gameName
    }));
}


function sendGameSelection() {
    lobby.send(JSON.stringify({
        event: "GameView"
    }));
}

