function SelectDigger(unitID, numberSlot, type, equip) {

    let unit = game.units[unitID];
    let graphics = game.add.graphics(0, 0);
    unit.selectDiggerLine = {graphics: graphics, radius: equip.equip.radius};
    game.floorObjectLayer.add(graphics);

    game.input.onDown.add(function () {
        dontMove = true;

        let x = (game.input.mousePointer.x + game.camera.x) / game.camera.scale.x;
        let y = (game.input.mousePointer.y + game.camera.y) / game.camera.scale.y;

        UnselectDigger();

        global.send(JSON.stringify({
            event: "useDigger",
            x: Math.round(x),
            y: Math.round(y),
            slot: Number(numberSlot),
            type_slot: type,
        }));
    });
}