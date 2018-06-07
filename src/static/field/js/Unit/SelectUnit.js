function SelectUnit(unit) {

    RemoveSelect();

    if (!unit.info.action && game.user.name === unit.info.owner) {
        CreateUnitSubMenu(unit);
    }

    field.send(JSON.stringify({
        event: "SelectUnit",
        x: Number(unit.info.x),
        y: Number(unit.info.y)
    }));
}