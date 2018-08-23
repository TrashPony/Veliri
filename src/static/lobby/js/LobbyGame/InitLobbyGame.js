function CreateNewLobbyGame(jsonMessage) {
    let game = JSON.parse(jsonMessage).game;
    console.log(jsonMessage);
    new Promise((resolve) => {
        CreateLobbyMenu(game.Name, JSON.parse(jsonMessage).error, true);
        return resolve();
    }).then(
        () => {
            lobby.send(
                JSON.stringify({
                    event: "GetSquad"
                })
            );

            let user = JSON.parse(jsonMessage).user_name;
            let users = JSON.parse(jsonMessage).game.Users;

            for (let name in users) {
                if (users.hasOwnProperty(name)) {
                    if (users[name].Respawn) {
                        CreateUserLine(name, users[name].LobbyReady, users[name].Respawn.id);
                    } else {
                        CreateUserLine(name, users[name].LobbyReady, "");
                    }
                    if (name === user) {
                        CreateSelectRespawn(name);
                        Respawn();
                    }
                }
            }
        }
    );
}

function InitLobbyGame(jsonMessage) {
    new Promise((resolve) => {
        CreateLobbyMenu(JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, false);
        return resolve();
    }).then(
        () => {
            lobby.send(
                JSON.stringify({
                    event: "GetSquad"
                })
            );

            let user = JSON.parse(jsonMessage).user_name;
            let users = JSON.parse(jsonMessage).game_users;

            for (let name in users) {
                if (users.hasOwnProperty(name)) {
                    if (users[name].Respawn) {
                        CreateUserLine(name, users[name].LobbyReady, users[name].Respawn.id);
                    } else {
                        CreateUserLine(name, users[name].LobbyReady, "");
                    }
                    if (name === user) {
                        CreateSelectRespawn(name);
                        Respawn();
                    }
                }
            }
        }
    );
}