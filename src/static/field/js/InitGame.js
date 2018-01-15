var idGame;

function InitGame() {
    idGame = getCookie("idGame");
    field.send(JSON.stringify({
        event: "InitGame",
        id_game: Number(idGame)
    }));
}

function getCookie(name) {
    var matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
}

function FieldCreate(jsonMessage) {
    var x = JSON.parse(jsonMessage).x_map;
    var y = JSON.parse(jsonMessage).y_map;
    Game(Number(x), Number(y)) // создаем окно игры размером х:у
}

function InitPlayer(jsonMessage) {
    var price = document.getElementById('price');
    price.innerHTML = JSON.parse(jsonMessage).player_price;
    var step = document.getElementById('step');
    step.innerHTML = JSON.parse(jsonMessage).game_step;
    var phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = JSON.parse(jsonMessage).game_phase;
    MY_ID = JSON.parse(jsonMessage).user_name;

    if (JSON.parse(jsonMessage).user_ready === true) {
        var ready = document.getElementById("Ready");
        ready.value = "Ты готов!";
        ready.className = "button noActive";
        ready.onclick = null
    }

    phase = JSON.parse(jsonMessage).game_phase;
}

function InitObstacle(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;

    var cell = cells[x + ":" + y];
    var obstacle = game.add.tileSprite(cell.x, cell.y, 100, 100, 'obstacle');
    obstacle.inputEnabled = true;
    obstacle.events.onInputOut.add(mouse_out);
}

function ReadyReader(jsonMessage) {
    var error = JSON.parse(jsonMessage).error;
    phase = JSON.parse(jsonMessage).phase;

    if (error === "") {
        var ready = document.getElementById("Ready");
        var phaseBlock = document.getElementById("phase");

        if (phase === "") {
            ready.value = "Ты готов!";
            ready.className = "button noActive";
            ready.onclick = null;
        } else {
            ready.value = "Готов!";
            ready.className = "button";
            ready.onclick = function () { Ready(); };

            phaseBlock.innerHTML = JSON.parse(jsonMessage).phase;
            var cells = document.getElementsByClassName("fieldUnit create");
            while (0 < cells.length) {
                if (cells[0]) {
                    cells[0].className = "fieldUnit open";
                }
            }
        }
    } else {
        if (error === "not units") {
            alert("У вас нет юнитов")
        }
    }
}