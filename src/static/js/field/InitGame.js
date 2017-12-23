var idGame;

function InitGame() {
    idGame = getCookie("idGame");
    sock.send(JSON.stringify({
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
    var xSize = JSON.parse(jsonMessage).x_map;
    var ySize =JSON.parse(jsonMessage).y_map;
    var main = document.getElementById("main");

    for (var y = 0; y < ySize; y++) {
        for (var x = 0; x < xSize; x++) {
            var div = document.createElement('div');
            div.className = "fieldUnit";
            div.id = x + ":" + y;
            div.innerHTML = x + ":" + y;
            main.appendChild(div);
        }
        var nline = document.createElement('div');
        nline.className = "nline";
        nline.innerHTML = "";
        main.appendChild(nline);
    }
}

function InitPlayer(jsonMessage) {
    var price = document.getElementById('price');
    price.innerHTML = JSON.parse(jsonMessage).player_price;
    var step = document.getElementById('step');
    step.innerHTML = JSON.parse(jsonMessage).game_step;
    var phaseGame = document.getElementById('phase');
    phaseGame.innerHTML = JSON.parse(jsonMessage).game_phase;

    if (JSON.parse(jsonMessage).user_ready === "true") {
        var ready = document.getElementById("Ready");
        ready.innerHTML = "Ты готов!";
        ready.style.backgroundColor = "#e1720f"
    }

    phase = JSON.parse(jsonMessage).game_phase;
}

function InitUnit(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var type = JSON.parse(jsonMessage).type_unit;
    var hp = JSON.parse(jsonMessage).hp;
    var action = JSON.parse(jsonMessage).unit_action;
    var userOwned = JSON.parse(jsonMessage).user_owned;

    var clicked_id = x + ":" + y;
    var cell = document.getElementById(clicked_id);

    if (type === "tank") cell.className = "fieldUnit tank";
    if (type === "scout") cell.className = "fieldUnit scout";
    if (type === "artillery") cell.className = "fieldUnit artillery";

    cell.onmouseover = function () {
        mouse_over(this.id);
    };
    cell.onmouseout = function () {
        mouse_out();
    };

    cell.innerHTML = "hp: " + hp;

    if (JSON.parse(jsonMessage).user_name === userOwned) {
        cell.style.color = "#fbfdff";
        cell.style.borderColor = "#fbfdff";
        cell.onclick = function () {
            SelectUnit(this.id)
        };
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
    }
}

function InitStructure(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
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
    }
}

function InitObstacle(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var coor_id = x + ":" + y;
    var cell = document.getElementById(coor_id);
    cell.className = "fieldUnit obstacle"
}

function ReadyReader(jsonMessage) {
    var error = JSON.parse(jsonMessage).error;
    phase = JSON.parse(jsonMessage).phase;

    if (error === "") {
        var ready = document.getElementById("Ready");
        var phaseBlock = document.getElementById("phase");

        if (phase === "") {
            ready.value = "Ты готов!";
            ready.style.backgroundColor = "#e1720f";
        } else {
            ready.value = "Готов!";
            if (phase === "move") {
                ready.style.backgroundColor = "#A8ADE1";
            }
            if (phase === "targeting") {
                ready.style.backgroundColor = "#E1C7A6";
            }
            if (phase === "attack") {
                ready.style.backgroundColor = "#E12D27";
            }
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