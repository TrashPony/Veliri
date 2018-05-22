function SelectUnit(unit) {

    RemoveSelect();

    field.send(JSON.stringify({
        event: "SelectUnit",
        x: Number(unit.info.x),
        y: Number(unit.info.y)
    }));
}
