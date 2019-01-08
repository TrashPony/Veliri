function CreateBoxes(boxes) {

    game.boxes = [];

    for (let i = 0; i < boxes.length; i++) {
        CreateBox(boxes[i])
    }
}

function CreateBox(mapBox) {

    game.boxes.push(mapBox);

    if (game.map.OneLayerMap.hasOwnProperty(mapBox.q) && game.map.OneLayerMap.hasOwnProperty(mapBox.r)) {

        let xy = GetXYCenterHex(mapBox.q, mapBox.r);
        let box = game.floorObjectLayer.create(xy.x, xy.y, mapBox.type);
        box.anchor.setTo(0.5);
        box.scale.set(0.2);
        box.angle = mapBox.rotate;

        if (!mapBox.underground) {
            let boxShadow = game.floorObjectLayer.create(xy.x + game.shadowXOffset, xy.y + game.shadowYOffset, mapBox.type);
            boxShadow.anchor.setTo(0.5);
            boxShadow.scale.set(0.2);
            boxShadow.tint = 0x000000;
            boxShadow.alpha = 0.6;
            boxShadow.angle = mapBox.rotate;
            box.shadow = boxShadow;

            mapBox.sprite = box;
            mapBox.shadow = boxShadow;
        }

        box.inputEnabled = true;
        box.input.pixelPerfectOver = true;
        box.input.pixelPerfectClick = true;

        let boxLine;
        box.events.onInputOver.add(function () {
            boxLine = game.floorObjectSelectLineLayer.create(xy.x, xy.y, mapBox.type);
            boxLine.anchor.setTo(0.5);
            boxLine.scale.set(0.22);
            boxLine.tint = 0x00FF00;
            boxLine.angle = mapBox.rotate;
        });

        box.events.onInputOut.add(function () {
            boxLine.destroy();
        });

        box.events.onInputDown.add(function () {
            game.squad.toBox = {
                boxID: mapBox.id,
                to: true,
                x: box.x,
                y: box.y
            }
        });
    }
}