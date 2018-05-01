function NewUser(jsonMessage) {
    // говорит всем кто в игре кто подключился
    var user = Object();
    user.Name = JSON.parse(jsonMessage).new_user;
    user.Ready = " Не готов";
    user.Respawn = JSON.parse(jsonMessage).respawn_name;

    CreateUserLine(user);
}

function JoinToLobby(jsonMessage) {
    var user = Object();
    // дает новому игроку данные по тем кто уже внутри
    user.Name = JSON.parse(jsonMessage).game_user;

    if (user.Ready = JSON.parse(jsonMessage).ready === "false") {
        user.Ready = " Не готов";
    } else {
        user.Ready = " Готов";
    }

    user.Respawn = JSON.parse(jsonMessage).respawn_name;

    CreateUserLine(user);
}