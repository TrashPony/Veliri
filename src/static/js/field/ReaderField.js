var SelectCell = [];

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
            if (error === "not units") {
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
            var Cell = {};
            Cell.x = x;
            Cell.y = y;
            Cell.id = coor_id;
            Cell.type = cell.className;
            SelectCell.push(Cell);
            cell.className = "fieldUnit move";
        }
    }

    if (event === "OpenCoordinate") {
        console.log(jsonMessage);
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        var idCell = x + ":" + y;
        OpenUnit(idCell)
    }

    if (event === "DellCoordinate") {
        console.log(jsonMessage);
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        var idDell = x + ":" + y;
        DelUnit(idDell)
    }

    if (event === "MoveUnit") {
        var errorMove = JSON.parse(jsonMessage).error;
        if (errorMove !== null) {
            DelMoveCell()
        }
    }
}

function OpenUnit(id) {

    var classUnit = "fieldUnit open";
    if (move != null) {
        classUnit = "fieldUnit move"
    }

    var Cell = document.getElementById(id);
    Cell.className = classUnit;
    Cell.id = id;
    Cell.innerHTML = id;
    Cell.style.color = "#fbfdff";
    Cell.style.borderColor = "#404040";
    Cell.onclick = function () {
        reply_click(this.id);
    };
    Cell.onmouseover = function () {
        mouse_over(this.id);
    };
    Cell.onmouseout = function () {
        mouse_out()
    };
}

function DelUnit(id) {
    var Cell = document.getElementById(id);
    Cell.className = "fieldUnit";
    Cell.innerHTML = id;
    Cell.style.color = "#FBFDFF";
    Cell.style.borderColor = "#404040";
    Cell.onclick = function () {
        reply_click(this.id);
    };
}

function DelMoveCell() {
    move = null;

    for (var i = 0; i < SelectCell.length; i++) {
        if (SelectCell[i].type === "fieldUnit open") {
            OpenUnit(SelectCell[i].id)
        }
        if (SelectCell[i].type === "fieldUnit") {
            DelUnit(SelectCell[i].id)
        }
        delete SelectCell[i];
    }
    SelectCell = [];
}