function NewUser(jsonMessage) {
    // говорит всем кто в игре кто подключился
    var user = JSON.parse(jsonMessage).user;

    CreateUserLine(user);
}

function JoinToLobby(jsonMessage) {
    // дает новому игроку данные по тем кто уже внутри
    var users = JSON.parse(jsonMessage).game_users;

    for (var i = 0; i < users.length; i++) {
        CreateUserLine(users[i]);
    }
}