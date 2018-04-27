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
        userTr.parentNode.removeChild(userTr);
    }

    if (event === "UserRefresh" || event === "JoinToLobby") {
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
}

function AddListSquad(jsonMessage) {
    var selectSquad = document.getElementById("listSquad");
    var squads = JSON.parse(jsonMessage).squads;

    for (var i = 0; i < squads.length; i++) {
        var squadSelect = document.createElement("option");

        squadSelect.value = squads[i].name;
        squadSelect.text = squads[i].name;
        squadSelect.id = squads[i].id;
        squadSelect.matherShip = squads[i].mather_ship;
        squadSelect.units = squads[i].units;

        squadSelect.onclick = function () {
            SelectSquad(this)
        };

        selectSquad.appendChild(squadSelect);
        selectSquad.value = "";
    }
}

function AddNewSquadInList(jsonMessage) {
    var selectSquad = document.getElementById("listSquad");
    var squad = JSON.parse(jsonMessage).squad;

    var squadSelect = document.createElement("option");

    squadSelect.value = squad.name;
    squadSelect.text = squad.name;
    squadSelect.id = squad.id;
    squadSelect.matherShip = squad.mather_ship;
    squadSelect.units = squad.units;

    squadSelect.onclick = function () {
        SelectSquad(this)
    };

    selectSquad.appendChild(squadSelect);
    selectSquad.value = squad.name;
}