function UnitMouseOver(unit) {
    unitTip(unit);

    if (game.Phase === "targeting") {
        field.send(JSON.stringify({
            event: "GetTargetZone",
            x: Number(unit.info.x),
            y: Number(unit.info.y),
            to_x: Number(unit.info.x),
            to_y: Number(unit.info.y)
        }));
    }
}

function UnitMouseOut() {
    TipOff();

    if (game.SelectLayer.children.length === 0 && game.Phase === "targeting") {
        RemoveTargetLine();
    }
}