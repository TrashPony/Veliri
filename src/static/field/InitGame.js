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

    if (JSON.parse(jsonMessage).user_ready === true) {
        var ready = document.getElementById("Ready");
        ready.value = "Ты готов!";
        ready.className = "button noActive";
        ready.onclick = null
    }

    phase = JSON.parse(jsonMessage).game_phase;
}

function InitUnit(jsonMessage) {

    DelMoveCoordinate();

    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var type = JSON.parse(jsonMessage).type_unit;
    var userOwned = JSON.parse(jsonMessage).user_owned;
    var userName  = JSON.parse(jsonMessage).user_name;
    var action = JSON.parse(jsonMessage).unit_action;
    var hp = JSON.parse(jsonMessage).hp;

    var cell = cells[x + ":" + y];
    var unit = game.add.sprite(cell.x + tileWidth / 2, cell.y + tileWidth / 2, type);

    game.physics.arcade.enable(unit);
    unit.inputEnabled = true;           // включаем ивенты на спрайт
    unit.anchor.setTo(0.35, 0.5);        // устанавливаем центр спрайта
    unit.scale.set(.32);                  // устанавливаем размер спрайта от оригинала
    unit.body.collideWorldBounds = true; // границы страницы

    unit.id = x + ":" + y;
    unit.events.onInputOver.add(mouse_over); // обрабатываем нажатие мышки
    unit.events.onInputOut.add(mouse_out);     // обрабатываем нажатие мышки

    var style;

    if (userName === userOwned) {
        style = { font: "52px Arial", fill: "#00ffff" };
        unit.events.onInputDown.add(SelectUnit);
        if (action === "false") {
            unit.tint = 0x757575; // накладывает фильтр со светом
        }
    } else {
        style = { font: "52px Arial", fill: "#ff0000" };
        unit.events.onInputDown.add(SelectTarget);
    }

    var label_score = game.add.text(x, y, "hp " + hp, style);
    unit.addChild(label_score);

    units[unit.id] = unit;
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
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;

    var cell = cells[x + ":" + y];
    var obstacle = game.add.tileSprite(cell.x, cell.y, 100, 100, 'obstacle');
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