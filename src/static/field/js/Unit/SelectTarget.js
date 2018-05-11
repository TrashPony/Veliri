function SelectTarget(clicked) {
    if (game.input.activePointer.leftButton.isDown) {

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
}

/*function TargetUnit() {
    var targetCell = document.getElementsByClassName("aim");
    while (targetCell.length > 0) {
        targetCell[0].remove();
    }
}*/