function MapView(jsonMessage) {
    var func = function () {
        CreateLobbyGame(this.id);
    };
    var funcMouse = function () {
        MouseOverMap(this.id);
    };
    var funcOutMouse = function () {
        MouseOutMap()
    };

    var map = Object();

    map.Name = JSON.parse(jsonMessage).name_map;
    map.Copasity = JSON.parse(jsonMessage).num_of_players;

    CreateLobbyLine('Map', 'SubMenu', 'Select SubMenu', JSON.parse(jsonMessage).name_map, func, funcMouse, funcOutMouse, map, JSON.parse(jsonMessage).user_name);
}

function GameView(jsonMessage) {
    var func = function () {
        sendJoinToLobbyGame(this.id);
    };

    var newGame = Object();

    newGame.Name = JSON.parse(jsonMessage).name_game;
    newGame.Map  = JSON.parse(jsonMessage).name_map;
    newGame.Creator = JSON.parse(jsonMessage).creator;
    newGame.Copasity = JSON.parse(jsonMessage).players;
    newGame.Players = JSON.parse(jsonMessage).num_of_players;

    CreateLobbyLine('Game', 'Menu', 'Select Menu', JSON.parse(jsonMessage).name_game, func, null, null, newGame, JSON.parse(jsonMessage).user_name);
}

function NotEndGame(jsonMessage) {
    var func = function () {
        JoinToGame(this.id);
    };

    var game = Object();

    game.Name = JSON.parse(jsonMessage).name_game;
    game.Id = JSON.parse(jsonMessage).id_game;
    game.Step = JSON.parse(jsonMessage).step_game;
    game.Phase = JSON.parse(jsonMessage).phase_game;
    game.Ready = !JSON.parse(jsonMessage).ready;

    CreateLobbyLine('NotEndGame', 'SubMenu', 'Select SubMenu', JSON.parse(jsonMessage).id_game, func, null, null, game, JSON.parse(jsonMessage).user_name);
}

function NewLobbyUser(jsonMessage) {
    var UsersInLobbyTable = document.getElementById("UsersInLobby");
    var NewUser = JSON.parse(jsonMessage).game_user;

    var userTr  = document.getElementById(NewUser + "allList");

    if (!userTr) {
        var tr = document.createElement("tr");
        tr.id = NewUser + "allList";
        tr.style.textAlign = "center";
        var td = document.createElement("td");
        td.className = "Value";
        td.innerHTML = NewUser;
        tr.appendChild(td);
        UsersInLobbyTable.appendChild(tr);
    }
}

function DelLobbyUser(jsonMessage) {
    var DelUser = JSON.parse(jsonMessage).game_user;
    var userTr  = document.getElementById(DelUser + "allList");
    if (userTr) {
        userTr.parentNode.removeChild(userTr);
    }
}