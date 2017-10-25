var idGame;

function InitGame() {
    idGame = getCookie("idGame");
    sendInitGame(idGame);
}

function getCookie(name) {
    var matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
}

function Field(xSize,ySize) {
    var main = document.getElementById("main");
    main.style.boxShadow = "25px 25px 20px rgba(0,0,0,0.5)";

    for (var y = 0; y < ySize; y++) {
        for (var x = 0; x < xSize; x++) {
            var div = document.createElement('div');
            div.className = "fieldUnit";
            div.id = x + ":" + y;
            div.innerHTML = x + ":" + y;
            div.onclick = function () {
                reply_click(this.id);
            };
            div.onmouseover = function () {
                mouse_over(this.id);
            };
            div.onmouseout = function () {
                mouse_out()
            };
            main.appendChild(div);
        }
        var nline = document.createElement('div');
        nline.className = "nline";
        nline.innerHTML = "";
        main.appendChild(nline);
    }
}

function sendInitGame(idGame) {
    sock.send(JSON.stringify({
        event: "InitGame",
        id_game: idGame
    }));
}

function InitPlayer(jsonMessage) {
    var price = document.getElementsByClassName('fieldInfo price');
    price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;
    var step = document.getElementsByClassName('fieldInfo step');
    step[0].innerHTML = "Ход № " + JSON.parse(jsonMessage).game_step;
    var phaseGame = document.getElementsByClassName('fieldInfo phase');
    phaseGame[0].innerHTML = "Фаза: " + JSON.parse(jsonMessage).game_phase;

    if (JSON.parse(jsonMessage).user_ready === "true") {
        ready = document.getElementById("Ready");
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
    cell.onclick = function () {
        SelectUnit(this.id)
    };
    cell.innerHTML = "hp: " + hp;

    if (JSON.parse(jsonMessage).user_name === userOwned) {
        cell.style.color = "#fbfdff";
        cell.style.borderColor = "#fbfdff";
    } else {
        cell.style.color = "#FF0117";
        cell.style.borderColor = "#FF0117";
    }
}

function InitResp(jsonMessage) {
    var x = JSON.parse(jsonMessage).respawn_x;
    var y = JSON.parse(jsonMessage).respawn_y;
    var coor_id = x + ":" + y;
    var cell = document.getElementById(coor_id);
    cell.className = "fieldUnit respawn";
    cell.innerHTML = "Resp: " + JSON.parse(jsonMessage).user_name;
}

function CreateUnit(jsonMessage) {
    if (JSON.parse(jsonMessage).error_type === "") {
        var userOwned = JSON.parse(jsonMessage).user_owned;
        var x = JSON.parse(jsonMessage).x;
        var y = JSON.parse(jsonMessage).y;

        var coor_id = x + ":" + y;
        var cell = document.getElementById(coor_id);
        cell.onclick = function () {
            SelectUnit(this.id)
        };
        if (typeUnit === "tank") cell.className = "fieldUnit tank";
        if (typeUnit === "scout") cell.className = "fieldUnit scout";
        if (typeUnit === "artillery") cell.className = "fieldUnit artillery";
        var price = document.getElementsByClassName('fieldInfo price');
        price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;

        if (JSON.parse(jsonMessage).user_name === userOwned) {
            cell.style.color = "#fbfdff";
            cell.style.borderColor = "#fbfdff";
        } else {
            cell.style.color = "#FF0117";
            cell.style.borderColor = "#FF0117";
        }

    } else {
        var cells = document.getElementsByClassName("fieldUnit create");
        var log = document.getElementById('fieldLog');
        while (0 < cells.length) {
            if (cells[0]) {
                cells[0].className = "fieldUnit open";
            }
        }
        if (JSON.parse(jsonMessage).error_type === "busy") {
            log.innerHTML = "Место занято"
        }

        if (JSON.parse(jsonMessage).error_type === "no many") {
            log.innerHTML = "Нет денег"
        }
        if (JSON.parse(jsonMessage).error_type === "not allow") {
            log.innerHTML = "Не разрешено"
        }
    }
    typeUnit = null;
}