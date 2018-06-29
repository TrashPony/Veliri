function CreateNewLobbyGame(jsonMessage) {
    let game = JSON.parse(jsonMessage).game;
    CreateLobbyMenu(game.Name, JSON.parse(jsonMessage).error, true);

    let user = JSON.parse(jsonMessage).user_name;

    let users = JSON.parse(jsonMessage).game.Users;

    for (let name in users) {
        if (users.hasOwnProperty(name)) {
            console.log(name);
            CreateUserLine(name, users[name].Ready);

            if (name === user) {
                CreateSelectRespawn(name);
                Respawn();
            }
        }
    }
}

function InitLobbyGame(jsonMessage) {
    CreateLobbyMenu(JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, false);

    let user = JSON.parse(jsonMessage).user_name;

    let users = JSON.parse(jsonMessage).game_users;

    for (let name in users) {
        if (users.hasOwnProperty(name)) {
            console.log(name);
            CreateUserLine(name, users[name].Ready);

            if (name === user) {
                CreateSelectRespawn(name);
                Respawn();
            }
        }
    }
}