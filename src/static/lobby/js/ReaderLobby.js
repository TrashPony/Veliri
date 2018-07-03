function ReaderLobby(jsonMessage) {
    let event = JSON.parse(jsonMessage).event;
    if (event === "InitLobby") {
        let login = document.getElementById('login');
        let userName = JSON.parse(jsonMessage).user_name;
        login.innerHTML = "Вы зашли как: " + userName;
    }

    if (event === "DisconnectLobby") {
        location.reload();
    }

    if (event === "GameRefresh") {
        DelElements("Select Menu");
    }

    if (event === "DelUser") {
        let userTr = document.getElementById(JSON.parse(jsonMessage).game_user);
        if (userTr !== null) {
            userTr.remove();
        }
    }

    if (event === "UserRefresh") {
        JoinToLobby(jsonMessage);
    }

    if (event === "GameView") {
        GameView(jsonMessage);
    }

    if (event === "DontEndGamesList") {
        NotEndGame(jsonMessage);
    }

    if (event === "MapView") {
        MapView(jsonMessage);
    }

    if (event === "NewUser") {
        NewUser(jsonMessage);
    }

    if (event === "CreateLobbyGame") {
        CreateNewLobbyGame(jsonMessage);
    }

    if (event === "initLobbyGame") {
        InitLobbyGame(jsonMessage);
    }

    if (event === "Respawn") {
        RespawnInit(jsonMessage);
    }

    if (event === "StartNewGame") {
        StartNewGame(jsonMessage);
    }

    if (event === "Ready") {
        Ready(jsonMessage);
    }

    if (event === "NewLobbyUser") {
        NewLobbyUser(jsonMessage);
    }

    if (event === "DelLobbyUser") {
        DelLobbyUser(jsonMessage);
    }

    if (event === "GetEquipping") {
        console.log(jsonMessage);
        EquippingParse(jsonMessage);
    }
}