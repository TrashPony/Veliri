function SelectUnit(unit) {
    if (move !== null) {
        DelMoveCoordinate();
    }

    var targetCell = document.getElementsByClassName("aim");
    while (targetCell.length > 0) {
        targetCell[0].remove();
    }

    var xy = unit.id.split(":");
    var x = xy[0];
    var y = xy[1];

    if (phase === "move") {
        move = unit;
    }

    if (phase === "targeting") {
        target = unit;
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

            var cell = cells[x + ":" + y];
            cell.tint = 0xb5b5ff;

            if (cell) {
                moveCell.push(cell); // кладем выделеные ячейки в масив что бы потом удалить
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

function SelectTarget(clicked) {
    var xy = clicked.id.split(":");

    var x = xy[0];
    var y = xy[1];
    var unit;
    var unit_x;
    var unit_y;

    if (phase === "targeting" && target !== null) {
        unit = target.id.split(":");
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

    if (phase === "move" && move !== null) {
        unit = move.id.split(":");
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

    if (phase === "Init" && typeUnit !== null && typeUnit !== undefined) {
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