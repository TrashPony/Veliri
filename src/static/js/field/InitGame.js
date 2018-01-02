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
    Game(Number(x), Number(y)) // создаем карту размером х:у
}

function InitPlayer(jsonMessage) {
    var price = document.getElementById('price');
    price.innerHTML = JSON.parse(jsonMessage).player_price;
    var step = document.getElementById('step');
    step.innerHTML = JSON.parse(jsonMessage).game_step;
    var phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = JSON.parse(jsonMessage).game_phase;

    if (JSON.parse(jsonMessage).user_ready === true) {
        var ready = document.getElementById("Ready");
        ready.value = "Ты готов!";
        ready.className = "button noActive";
        ready.onclick = null
    }

    phase = JSON.parse(jsonMessage).game_phase;
}

function InitUnit(jsonMessage) {

    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var type = JSON.parse(jsonMessage).type_unit;
    var userOwned = JSON.parse(jsonMessage).user_owned;
    var userName  = JSON.parse(jsonMessage).user_name;

    var cell = cells[x + ":" + y];
    var scout = game.add.tileSprite(cell.x, cell.y, 100, 100, type);
    scout.inputEnabled = true; // включаем ивенты на спрайт
    scout.id = x + ":" + y;
    scout.events.onInputOver.add(mouse_over, this); // обрабатываем нажатие мышки
    scout.events.onInputOut.add(mouse_out, this);     // обрабатываем нажатие мышки

    if ( userName === userOwned) {
        scout.events.onInputDown.add(SelectUnit, this);
    }

    cells[scout.id] = scout;
    /*DelMoveCoordinate();

    var hp = JSON.parse(jsonMessage).hp;
    var action = JSON.parse(jsonMessage).unit_action;

    cell.innerHTML = "hp: " + hp;

    if ( userName === userOwned) {
        cell.style.color = "#fbfdff";
        cell.style.borderColor = "#fbfdff";

        if (action === "false") {
            cell.style.filter = "brightness(50%)";
        } else {
            cell.style.filter = "brightness(100%)";
        }
    } else {
        cell.style.color = "#FF0117";
        cell.style.borderColor = "#FF0117";
        cell.onclick = function () {
            SelectTarget(this.id)
        };
    }*/
}

function InitStructure(jsonMessage) {
    /*var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var type = JSON.parse(jsonMessage).type_structure;
    var user = JSON.parse(jsonMessage).user_owned;
    var coor_id = x + ":" + y;
    var cell = document.getElementById(coor_id);

    cell.onmouseover = function () {
        mouse_over(this.id);
    };
    cell.onmouseout = function () {
        mouse_out();
    };

    if (type === "respawn") {
        if (user === JSON.parse(jsonMessage).user_name) {
            cell.style.color = "#fbfdff";
            cell.style.borderColor = "#fbfdff";
            cell.className = "fieldUnit respawn";
            cell.innerHTML = "Resp: " + JSON.parse(jsonMessage).user_name;
        } else {
            cell.className = "fieldUnit respawn";
            cell.innerHTML = "Resp: " + JSON.parse(jsonMessage).user_name;
            cell.style.color = "#FF0117";
            cell.style.borderColor = "#FF0117";
        }
    }*/
}

function InitObstacle(jsonMessage) {
    /*var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var coor_id = x + ":" + y;
    var cell = document.getElementById(coor_id);
    cell.className = "fieldUnit obstacle"*/
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