function ReaderLobby(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    if (event === "InitLobby") {
        var login = document.getElementById('login');
        var userName = JSON.parse(jsonMessage).user_name;
        login.innerHTML = "Вы зашли как: " + userName;
    }

    if (event === "DisconnectLobby") {
        location.reload();
    }

    if (event === "GameRefresh") {
        DelElements("Select Menu");
    }

    if (event === "DelUser") {
        var userTr = document.getElementById(JSON.parse(jsonMessage).game_user);
        userTr.remove();
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
        console.log(jsonMessage);
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

    if (event === "GetMatherShips") {
        MatherShipsParse(jsonMessage);
    }

    if (event === "GetDetailOfUnits") {
        DetailUnitParse(jsonMessage)
    }

    if (event === "GetListSquad") {
        AddListSquad(jsonMessage);
    }

    if (event === "AddNewSquad") {
        AddNewSquadInList(jsonMessage);
    }

    if (event === "UpdateSquad") {
        UpdateSquad(jsonMessage);
    }

    if (event === "RemoveSquad") {
        RemoveSquad(jsonMessage);
    }

    if (event === "GetEquipping") {
        EquippingParse(jsonMessage);
    }

    if (event === "UnitConstructorUpdate") {
        var unit = JSON.parse(jsonMessage).unit;
        UpdateUnitInfo(unit);
    }
}