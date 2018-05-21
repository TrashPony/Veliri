function SelectUnit(unit) {

    field.send(JSON.stringify({
        event: "SelectUnit",
        unit_id: Number(unit.info.id)
    }));

    /*var x = JSON.parse(jsonMessage).x;
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

            SelectMoveCoordinate(x,y)
        } else {
            if (phase === "targeting") {
                /*coordinate = x + ":" + y;
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
    }*/
}

/*function SelectUnit(unit) {
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
}*/