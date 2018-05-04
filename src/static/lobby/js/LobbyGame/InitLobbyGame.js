function CreateNewLobbyGame(jsonMessage) {
    var game = JSON.parse(jsonMessage).game;
    CreateLobbyMenu(game.Name, JSON.parse(jsonMessage).error, true);

    var user = JSON.parse(jsonMessage).game.Creator;

    CreateUserLine(user);
    CreateSelectRespawn(user.Name);
    Respawn();
}

function InitLobbyGame(jsonMessage) {
    CreateLobbyMenu(JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, false);

    var user = JSON.parse(jsonMessage).user;

    CreateUserLine(user);
    CreateSelectRespawn(user.Name);
    Respawn();

    var users = JSON.parse(jsonMessage).game_users;

    for (var i = 0; i < users.length; i++) {
        CreateUserLine(users[i]);
    }
}