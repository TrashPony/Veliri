function CreateNewLobbyGame(jsonMessage) {
    CreateLobbyMenu(JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, true);

    var user = Object();
    user.Name = JSON.parse(jsonMessage).user_name;
    user.Ready = " Не готов";
    user.Respawn = JSON.parse(jsonMessage).respawn_name;

    CreateLobbyLine('User', 'gameInfo', null, null, null, null, null, user, JSON.parse(jsonMessage).user_name);

    Respawn();
}

function InitLobbyGame(jsonMessage) {

    CreateLobbyMenu(JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, false);

    if (JSON.parse(jsonMessage).error === "") {
        var text = JSON.parse(jsonMessage).user_name + " Не готов";
        CreateLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).user_name, null, null, null, text, JSON.parse(jsonMessage).user_name);
        Respawn();
    }
}

function NewUser(jsonMessage) {
    var user = Object();
    user.Name = JSON.parse(jsonMessage).new_user;
    user.Ready = " Не готов";
    user.Respawn = JSON.parse(jsonMessage).respawn_name;

    CreateLobbyLine('User', 'gameInfo', null, null, null, null, null, user, JSON.parse(jsonMessage).user_name);
}

function JoinToLobby(jsonMessage) {
    var user = Object();
    user.Name = JSON.parse(jsonMessage).game_user;
    user.Ready = JSON.parse(jsonMessage).ready;
    user.Respawn = JSON.parse(jsonMessage).respawn_name;

    CreateLobbyLine('User', 'gameInfo', null, null, null, null, null, user, JSON.parse(jsonMessage).user_name);
}

function Ready(jsonMessage) {
    var error = JSON.parse(jsonMessage).error;
    var user = Object();

    if (error === "") {

        user.Name = JSON.parse(jsonMessage).game_user;
        user.Ready = JSON.parse(jsonMessage).ready;
        user.Respawn = JSON.parse(jsonMessage).respawn_name;

        CreateLobbyLine('User', 'gameInfo', null, null, null, null, null, user, JSON.parse(jsonMessage).user_name);

        if (user === JSON.parse(jsonMessage).user_name && JSON.parse(jsonMessage).respawn_name === "") {
            //CreateSelectRespawn(user, user + " Не готов ");
        }
    } else {
        //CreateSelectRespawn(user, user + " Не готов ");
    }
    Respawn();
}

function Respawn(jsonMessage) {
    //Create and append the options
    var select = document.getElementById("RespawnSelect");
    if (select) {
        var option = document.createElement("option");
        option.className = "RespawnOption";
        option.value = JSON.parse(jsonMessage).respawn;
        option.text = JSON.parse(jsonMessage).respawn_name;
        select.appendChild(option);
    }
}

function StartNewGame(jsonMessage) {
    if (JSON.parse(jsonMessage).error === "") {
        toField = true;
        var idGame = JSON.parse(jsonMessage).id_game;
        document.cookie = "idGame=" + idGame + "; path=/;";
        location.href = "http://" + window.location.host + "/field";
    } else {
        if (JSON.parse(jsonMessage).error === "Players < 2") {
            alert("Ошибка: Мало игроков для старта");
        }
        if (JSON.parse(jsonMessage).error === "error ad to DB") {
            alert("Неизвестная ошибка");
        }
        if (JSON.parse(jsonMessage).error === "PlayerNotReady") {
            alert("Ошибка: не все игроки готовы");
        }
    }
}