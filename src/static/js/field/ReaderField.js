function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var price;
    var cell;
    var log;
    var ready;
    var info;
    var coor_id;
    var x;
    var y;
    var userOwned;

    if (event === "InitPlayer") {
        InitPlayer(jsonMessage);
    }

    if (event === "InitMap") {
        Field(JSON.parse(jsonMessage).x_map, JSON.parse(jsonMessage).y_map)
    }

    if (event === "InitUnit") {
        console.log("message: " + jsonMessage);
        InitUnit(jsonMessage);
    }

    if (event === "InitResp") {
        InitResp(jsonMessage);
    }

    if (event === "CreateUnit") {
        CreateUnit(jsonMessage);
    }

    if (event === "emptyCoordinate") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        coor_id = x + ":" + y;
        cell = document.getElementById(coor_id);
        if (cell) {
            cell.className = "fieldUnit open";
        }
    }

    if (event === "SelectCoordinateCreate") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        coor_id = x + ":" + y;
        cell = document.getElementById(coor_id);
        if (cell) {
            cell.className = "fieldUnit create";
        }
    }

    if (event === "MouseOver") {
        info = document.getElementById('unitInfo');
        info.innerHTML = "Тип Юнита: " + JSON.parse(jsonMessage).type_unit + "<br>" +
            "Владелец: " + JSON.parse(jsonMessage).user_owned + "<br>" +
            "hp: " + JSON.parse(jsonMessage).hp + "<br>" +
            "Ходил: " + JSON.parse(jsonMessage).unit_action + "<br>" +
            "Цель " + JSON.parse(jsonMessage).target + "<br>" +
            "Урон: " + JSON.parse(jsonMessage).damage + "<br>" +
            "Скорость: " + JSON.parse(jsonMessage).move_speed + "<br>" +
            "Инициатива: " + JSON.parse(jsonMessage).init + "<br>" +
            "Дальность атаки: " + JSON.parse(jsonMessage).range_attack + "<br>" +
            "Дальность обзора: " + JSON.parse(jsonMessage).range_view + "<br>" +
            "Площадь атаки: " + JSON.parse(jsonMessage).area_attack + "<br>" +
            "Тип атаки: " + JSON.parse(jsonMessage).type_attack
    }
    if (event === "Ready") {
        var error = JSON.parse(jsonMessage).error;
        phase = JSON.parse(jsonMessage).phase;
        console.log(phase);
        if (error === "") {
            ready = document.getElementById("Ready");
            var phaseBlock = document.getElementById("phase");

            if (phase === "") {
                ready.innerHTML = "Ты готов!";
                ready.style.backgroundColor = "#e1720f";
            } else {
                ready.innerHTML = "Готов!";
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
            if (error === "not units"){
                alert("У вас нет юнитов")
            }
        }
    }

    if (event === "SelectUnit") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        coor_id = x + ":" + y;
        cell = document.getElementById(coor_id);
        if (cell) {
            cell.className = "fieldUnit move";
        }
    }

    if (event === "MoveUnit") {
        move = null;

        var moveCells = document.getElementsByClassName("fieldUnit move");

        while (0 < moveCells.length) {
            if (moveCells[0]) {
                moveCells[0].className = "fieldUnit";
            }
        }

        error = JSON.parse(jsonMessage).error;
        if (error === "") {
            x = JSON.parse(jsonMessage).x;
            y = JSON.parse(jsonMessage).y;
            var toX = JSON.parse(jsonMessage).to_x;
            var toY = JSON.parse(jsonMessage).to_y;
            var type = JSON.parse(jsonMessage).type_unit;
            Move(x,y,toX,toY, type)             // TODO копировать ячейку в новую координату, а старую закрыть
        } else {

        }
    }
}

function Move(x,y,toX,toY,type) {
    var newCoor_id = toX + ":" + toY;
    var newCell = document.getElementById(newCoor_id);
    newCell.onclick = function () {
        SelectUnit(this.id)
    };
    if (type === "tank") newCell.className = "fieldUnit tank";
    if (type === "scout") newCell.className = "fieldUnit scout";
    if (type === "artillery") newCell.className = "fieldUnit artillery";

    var oldCoor_id = x + ":" + y;
    var oldCell = document.getElementById(oldCoor_id);
    oldCell.onclick = function () {
        reply_click(this.id);
    };
    oldCell.innerHTML = oldCoor_id;
    oldCell.className = "fieldUnit";
    oldCell.style.borderColor = "#404040";

    var moveCells = document.getElementsByClassName("fieldUnit open");
    while (0 < moveCells.length) {
        if (moveCells[0]) {
            moveCells[0].className = "fieldUnit";
        }
    }

    DellUnit("fieldUnit open");
    DellUnit("fieldUnit tank");
    DellUnit("fieldUnit scout");
    DellUnit("fieldUnit artillery");

    sock.send(JSON.stringify({
        event: "getPermittedCoordinates",
        id_game: Number(idGame)
    }));
}

function DellUnit(unit) {
    var moveCells = document.getElementsByClassName(unit);
    while (0 < moveCells.length) {
        if (moveCells[0]) {
            var id = moveCells[0].id;
            var Cell = document.getElementById(id);
            Cell.innerHTML = ""+id;
            Cell.className = "fieldUnit";
            Cell.style = null;
        }
    }
}