var SelectCell = [];

function ReadResponse(jsonMessage) {
    var event = JSON.parse(jsonMessage).event;
    var cell;
    var info;
    var coor_id;
    var x;
    var y;
    var idDell;

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

    if (event === "Ready") {
        ReadyReader(jsonMessage);
    }

    if (event === "SelectUnit") {
        Select(jsonMessage);
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

    if (event === "OpenCoordinate") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        var idCell = x + ":" + y;
        OpenUnit(idCell)
    }

    if (event === "DellCoordinate") {
        x = JSON.parse(jsonMessage).x;
        y = JSON.parse(jsonMessage).y;
        idDell = x + ":" + y;
        DelUnit(idDell)
    }

    if (event === "MouseOver") {
        info = document.getElementById('unitInfo');
        if (JSON.parse(jsonMessage).target !== "") {
            var xy = JSON.parse(jsonMessage).target.split(":");
            x = xy[0];
            y = xy[1];
            var idTarget = x + ":" + y;
            targetCell = document.getElementById(idTarget);

            var div = document.createElement('div');
            div.className = "aim mouse";
            targetCell.appendChild(div);
        }

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

    if (event === "MoveUnit") {
        var errorMove = JSON.parse(jsonMessage).error;
        var action = JSON.parse(jsonMessage).unit_action;
        if (action === "false") {
            x = JSON.parse(jsonMessage).x;
            y = JSON.parse(jsonMessage).y;
            idDell = x + ":" + y;
            cell = document.getElementById(idDell);
            cell.style.filter = "brightness(50%)";
        }

        if (errorMove !== null) {
            DelMoveCell()
        }
    }

    if (event === "TargetUnit") {
        var targetCell = document.getElementsByClassName("aim");
        while (targetCell.length > 0) {
            targetCell[0].remove();
        }
    }
}

function Select(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var errorSelect = JSON.parse(jsonMessage).error;
    var phase = JSON.parse(jsonMessage).phase;
    var coor_id;
    var cell;
    var Cell;

    if (errorSelect === "") {
        if (phase === "move") {
            coor_id = x + ":" + y;
            cell = document.getElementById(coor_id);
            if (cell) {
                Cell = {};
                Cell.x = x;
                Cell.y = y;
                Cell.id = coor_id;
                Cell.type = cell.className;
                SelectCell.push(Cell);
                cell.style.filter = "brightness(85%)";
                cell.className = "fieldUnit move";
            }
        } else {
            if (phase === "targeting") {
                coor_id = x + ":" + y;
                cell = document.getElementById(coor_id);
                if (cell) {
                    Cell = {};
                    Cell.x = x;
                    Cell.y = y;
                    Cell.id = coor_id;
                    SelectCell.push(Cell);

                    var div = document.createElement('div');
                    div.className = "aim";
                    cell.appendChild(div);
                }
            }
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
    Cell.style.filter = "brightness(100%)";

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
    Cell.style.filter = "brightness(100%)";

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