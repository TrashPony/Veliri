function CreateUnit(jsonMessage) {
    if (JSON.parse(jsonMessage).error_type === "") {
        var price = document.getElementsByClassName('fieldInfo price');
        price[0].innerHTML = "Твои Деньги: " + JSON.parse(jsonMessage).player_price;
    } else {
        var log = document.getElementById('fieldLog');

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

    var cells = document.getElementsByClassName("fieldUnit create");
    while (0 < cells.length) {
        if (cells[0]) {
            cells[0].className = "fieldUnit open";
        }
    }
    typeUnit = null;
}

function SelectUnit(id) {
    if (move !== null) {
        DelMoveCoordinate();
    }

    var targetCell = document.getElementsByClassName("aim");
    while (targetCell.length > 0) {
        targetCell[0].remove();
    }

    var xy = id.split(":");
    var x = xy[0];
    var y = xy[1];

    if (phase === "move") {
        move = id;
    }

    if (phase === "targeting") {
        target = id;
    }

    field.send(JSON.stringify({
        event: "SelectUnit",
        x: Number(x),
        y: Number(y)
    }));
}

function setUnitAction(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var errorSelect = JSON.parse(jsonMessage).error;
    var phase = JSON.parse(jsonMessage).phase;
    var coordinate;
    var cell;
    var Cell;

    if (errorSelect === "") {
        if (phase === "move") {

            var buttonSkip = document.getElementById("SkipButton");
            buttonSkip.className = "button";
            buttonSkip.onclick = function () {
                var unit = move.split(":");
                var unit_x = unit[0];
                var unit_y = unit[1];
                field.send(JSON.stringify({
                    event: "SkipMoveUnit",
                    x: Number(unit_x),
                    y: Number(unit_y)
                }));
            };

            coordinate = x + ":" + y;
            cell = document.getElementById(coordinate);
            if (cell) {
                Cell = {};
                Cell.x = x;
                Cell.y = y;
                Cell.id = coordinate;
                Cell.type = cell.className;
                SelectCell.push(Cell);
                cell.style.filter = "brightness(85%)";
                cell.className = "fieldUnit move";
            }
        } else {
            if (phase === "targeting") {
                coordinate = x + ":" + y;
                cell = document.getElementById(coordinate);
                if (cell) {
                    Cell = {};
                    Cell.x = x;
                    Cell.y = y;
                    Cell.id = coordinate;
                    SelectCell.push(Cell);

                    var div = document.createElement('div');
                    div.className = "aim";
                    cell.appendChild(div);
                }
            }
        }
    }
}

function SelectTarget(clicked_id) {
    var xy = clicked_id.split(":");

    var x = xy[0];
    var y = xy[1];
    var unit;
    var unit_x;
    var unit_y;

    if(phase === "targeting" && target !== null) {
        unit = target.split(":");
        unit_x = unit[0];
        unit_y = unit[1];

        field.send(JSON.stringify({
            event: "TargetUnit",
            x: Number(unit_x),
            y: Number(unit_y),
            target_x: Number(x),
            target_y: Number(y)
        }));
    } else {
        target = null;
    }

    if(phase === "move" && move !== null) {
        unit = move.split(":");
        unit_x = unit[0];
        unit_y = unit[1];
        field.send(JSON.stringify({
            event: "MoveUnit",
            x: Number(unit_x),
            y: Number(unit_y),
            to_x: Number(x),
            to_y: Number(y)
        }));
    } else {
        move = null;
    }

    if(phase === "Init" && typeUnit !== null && typeUnit !== undefined) {
        field.send(JSON.stringify({
            event: "CreateUnit",
            type_unit: typeUnit,
            id_game: Number(idGame),
            x: Number(x),
            y: Number(y)
        }));
    } else {
        typeUnit = null;
    }
}

function MoveUnit(jsonMessage) {
    var errorMove = JSON.parse(jsonMessage).error;
    var action = JSON.parse(jsonMessage).unit_action;
    if (action === "false") {
        var x = JSON.parse(jsonMessage).x;
        var y = JSON.parse(jsonMessage).y;
        var idDell = x + ":" + y;
        var cell = document.getElementById(idDell);
        cell.style.filter = "brightness(50%)";
    }

    if (errorMove !== null) {
        DelMoveCoordinate()
    }
}

function TargetUnit() {
    var targetCell = document.getElementsByClassName("aim");
    while (targetCell.length > 0) {
        targetCell[0].remove();
    }
}

function AttackUnit(jsonMessage) {
    var attackX = JSON.parse(jsonMessage).x;
    var attackY = JSON.parse(jsonMessage).y;
    var toX = JSON.parse(jsonMessage).to_x;
    var toY = JSON.parse(jsonMessage).to_y;

    var attackID = attackX + ":" + attackY;
    var cell = document.getElementById(attackID);
    cell.innerHTML = "ПЫЩЬ1";

    var targetID = toX + ":" + toY;
    var targeCell = document.getElementById(targetID);
    targeCell.innerHTML = "БдфЩь!";
}