function SelectDigger(coordinates, slot, typeSlot) {
    for (let i in coordinates) {
        if (coordinates.hasOwnProperty(i) && coordinates[i]) {
            let xy = GetXYCenterHex(coordinates[i].q, coordinates[i].r);
            let select = game.floorSelectLineLayer.create(xy.x, xy.y, 'selectEmpty');
            select.alpha = 0.5;
            select.scale.setTo(0.5);
            select.anchor.setTo(0.5);

            select.inputEnabled = true;
            select.events.onInputDown.add(function () {
                global.send(JSON.stringify({
                    event: "useDigger",
                    q: coordinates[i].q,
                    r: coordinates[i].r,
                    slot: slot,
                    type_slot: typeSlot,
                }));
            });

            select.events.onInputOver.add(function () {
                select.animations.add('select');
                select.animations.play('select', 5, true);
            });
            select.events.onInputOut.add(function () {
                select.animations.getAnimation('select').stop(false);
                select.animations.frame = 0;
            });
            select.input.priorityID = 1;
        }
    }
}