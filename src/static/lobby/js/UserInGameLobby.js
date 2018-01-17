function CreateNewLobbyGame(jsonMessage) {
    // создаем меню внутренниго лоби
    CreateLobbyMenu(JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, true);

    var user = Object();
    user.Name = JSON.parse(jsonMessage).user_name;
    user.Ready = " Не готов";
    user.Respawn = JSON.parse(jsonMessage).respawn_name;
    // создаем в лоби строку с пользователем
    CreateLobbyLine('User', 'gameInfo', "User", user.Name, null, null, null, user, user.Name);

    Respawn();
}

function InitLobbyGame(jsonMessage) {
    // создаем меню внутренниго лоби
    CreateLobbyMenu(JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, false);

    if (JSON.parse(jsonMessage).error === "") {

        var user = Object();
        user.Name = JSON.parse(jsonMessage).user_name;
        user.Ready = " Не готов";
        user.Respawn = JSON.parse(jsonMessage).respawn_name;

        // создаем в лоби строку с пользователем
        CreateLobbyLine('User', 'gameInfo', "User", user.Name, null, null, null, user, user.Name);
    }
}

function NewUser(jsonMessage) {

    var user = Object();
    user.Name = JSON.parse(jsonMessage).new_user;
    user.Ready = " Не готов";
    user.Respawn = JSON.parse(jsonMessage).respawn_name;

    CreateLobbyLine('User', 'gameInfo', "User", user.Name, null, null, null, user, JSON.parse(jsonMessage).user_name);
}

function JoinToLobby(jsonMessage) {
    var user = Object();

    user.Name = JSON.parse(jsonMessage).game_user;

    if (user.Ready = JSON.parse(jsonMessage).ready === "false") {
        user.Ready = " Не готов";
    } else {
        user.Ready = " Готов";
    }

    user.Respawn = JSON.parse(jsonMessage).respawn_name;

    CreateLobbyLine('User', 'gameInfo', "User", user.Name, null, null, null, user, JSON.parse(jsonMessage).user_name);
}

function Ready(jsonMessage) {
    var error = JSON.parse(jsonMessage).error;
    var ownedName = JSON.parse(jsonMessage).user_name;
    var user = Object();
    user.Name = JSON.parse(jsonMessage).game_user;

    var userRespawnCell = document.getElementById(user.Name).cells[1];
    var userReadyCell = document.getElementById(user.Name).cells[2];

    if (user.Ready = JSON.parse(jsonMessage).ready === "false") {
        user.Ready = " Не готов";
        userReadyCell.innerHTML = user.Ready;
        userReadyCell.className = "Failed";
    } else {
        user.Ready = " Готов";
        userReadyCell.innerHTML = user.Ready;
        userReadyCell.className = "Success";
    }

    user.Respawn = JSON.parse(jsonMessage).respawn_name;

    userRespawnCell.innerHTML = user.Respawn;


    if (error === "") {
        if (ownedName === user.Name && user.Respawn === "") {
            CreateSelectRespawn(user.Name);
        }
    } else {
        CreateSelectRespawn(user.Name);
    }

    Respawn();
}

function RespawnInit(jsonMessage) {
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