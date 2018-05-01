function CreateNewLobbyGame(jsonMessage) {
    var game = JSON.parse(jsonMessage).game;
    CreateLobbyMenu(game.Name, JSON.parse(jsonMessage).error, true);

    var user = Object();
    user.Name = JSON.parse(jsonMessage).user_name;
    user.Ready = " Не готов";

    CreateUserLine(user);
    CreateSelectRespawn(user.Name);
    Respawn();
}

function InitLobbyGame(jsonMessage) {
    CreateLobbyMenu(JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, false);

    if (JSON.parse(jsonMessage).error === "") {

        var user = Object();
        user.Name = JSON.parse(jsonMessage).user_name;
        user.Ready = " Не готов";

        CreateUserLine(user);
        CreateSelectRespawn(user.Name);
        Respawn();
    }
}