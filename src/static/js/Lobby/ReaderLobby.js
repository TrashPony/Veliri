function ReaderLobby(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var user;
    var func;
    var text;
    var textButton;
    var funcButton;
    var funcMouse;
    var funcOutMouse;

    if (event === "InitLobby") {
        var login = document.getElementById('login');
        var userName = JSON.parse(jsonMessage).user_name;
        login.innerHTML = "Вы зашли как: " + userName;
    }

    if (event === "DisconnectLobby") {
        location.reload()
    }
    if (event === "GameRefresh") {
        DelElements("Select Game");
    }
    if (event === "DelUser") {
        DelElements("User List");
    }

    if (event === "UserRefresh" || event === "JoinToLobby") {
        if (JSON.parse(jsonMessage).ready === "true") {
            text = JSON.parse(jsonMessage).game_user + " Готов! Респаун: " + JSON.parse(jsonMessage).respawn_name;
            CreateLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).game_user, null, null, null, text, "");
        } else {
            text = JSON.parse(jsonMessage).game_user + " Не готов";
            CreateLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).game_user, null, null, null, text, JSON.parse(jsonMessage).user_name);
        }
    }

    if (event === "GameView") {
        func = function () {
            sendJoinToLobbyGame(this.id);
        };

        var newGame = new Object();

        newGame.Name = JSON.parse(jsonMessage).name_game;
        newGame.Map  = JSON.parse(jsonMessage).name_map;
        newGame.Creator = JSON.parse(jsonMessage).creator;
        newGame.Copasity = JSON.parse(jsonMessage).players;
        newGame.Players = JSON.parse(jsonMessage).num_of_players;

        CreateLobbyLine('Game', 'Menu', 'Select Menu', JSON.parse(jsonMessage).name_game, func, null, null, newGame, JSON.parse(jsonMessage).user_name);
    }

    if (event === "DontEndGamesList") {
        func = function () {
            JoinToGame(this.id);
        };

        var game = new Object();

        game.Name = JSON.parse(jsonMessage).name_game;
        game.Id = JSON.parse(jsonMessage).id_game;
        game.Step = JSON.parse(jsonMessage).step_game;
        game.Phase = JSON.parse(jsonMessage).phase_game;
        game.Ready = !JSON.parse(jsonMessage).ready;

        CreateLobbyLine('NotEndGame', 'SubMenu', 'Select SubMenu', JSON.parse(jsonMessage).id_game, func, null, null, game, JSON.parse(jsonMessage).user_name);
    }

    if (event === "MapView") {
        func = function () {
            CreateLobbyGame(this.id);
        };
        funcMouse = function () {
            MouseOverMap(this.id);
        };
        funcOutMouse = function () {
            MouseOutMap()
        };

        var map = new Object();

        map.Name = JSON.parse(jsonMessage).name_map;
        map.Copasity = JSON.parse(jsonMessage).num_of_players;

        CreateLobbyLine('Map', 'SubMenu', 'Select SubMenu', JSON.parse(jsonMessage).name_map, func, funcMouse, funcOutMouse, map, JSON.parse(jsonMessage).user_name);
    }

    if (event === "NewUser") {
        text = JSON.parse(jsonMessage).new_user + " Не готов";
        CreateLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).new_user, null, null, null, text, JSON.parse(jsonMessage).user_name);
    }

    if (event === "CreateLobbyGame") {

        textButton = "Начать";
        funcButton = CreateNewGame;

        CreateLobbyMenu(textButton, funcButton, JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, true);

        text = JSON.parse(jsonMessage).user_name + " Не готов";

        CreateInLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).user_name, null, null, null, text, JSON.parse(jsonMessage).user_name);

        Respawn();

    }

    if (event === "initLobbyGame") {
        textButton = "Готов";
        funcButton = function () {
            sendReady(this.id)
        };
        CreateLobbyMenu(textButton, funcButton, JSON.parse(jsonMessage).name_game, JSON.parse(jsonMessage).error, false);
        if (JSON.parse(jsonMessage).error === "") {
            text = JSON.parse(jsonMessage).user_name + " Не готов";
            CreateLobbyLine('gameInfo', 'User List', JSON.parse(jsonMessage).user_name, null, null, null, text, JSON.parse(jsonMessage).user_name);
            Respawn();
        }
    }

    if (event === "Respawn") {
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

    if (event === "StartNewGame") {
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

    if (event === "Ready") {
        var ready;
        var error = JSON.parse(jsonMessage).error;
        user = JSON.parse(jsonMessage).game_user;
        if (error === "") {
            if (JSON.parse(jsonMessage).ready === "true") {
                ready = "Готов! Респаун: " + JSON.parse(jsonMessage).respawn_name;
            } else {
                ready = " Не готов ";
            }
            var userBlock = document.getElementById(user);
            userBlock.innerHTML = user + " " + ready;
            if (user === JSON.parse(jsonMessage).user_name && JSON.parse(jsonMessage).respawn_name === "") {
                CreateSelectRespawn(user, user + " Не готов ");
            }
        } else {
            CreateSelectRespawn(user, user + " Не готов ");
        }
        Respawn();
    }
}

function CreateInLobbyLine(gameContent, className, id, func, funcMouse, funcOutMouse, text, owned) {
    var list = document.getElementById(gameContent);
    var div = document.createElement('div');
    div.style.wordWrap = 'break-word';
    div.className = className;
    div.id = id;
    div.onclick = func;
    div.onmouseover = funcMouse;
    div.onmouseout = funcOutMouse;
    div.appendChild(document.createTextNode(text));
    list.appendChild(div);

    if (id === owned) {
        CreateSelectRespawn(id, text)
    }
}