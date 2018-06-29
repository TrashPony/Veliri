function NewUser(jsonMessage) {
    // говорит всем кто в игре кто подключился
    let user = JSON.parse(jsonMessage).game_user;
    let userReady = JSON.parse(jsonMessage).ready;

    CreateUserLine(user, userReady);
}

function JoinToLobby(jsonMessage) {
    // дает новому игроку данные по тем кто уже внутри

    let users = JSON.parse(jsonMessage).game_users;

    for (let i = 0; i < users.length; i++) {
        CreateUserLine(users[i]);
    }
}