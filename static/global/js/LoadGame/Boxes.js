function CreateBoxes(boxes) {
    for (let i = 0; i < boxes.length; i++) {
        CreateBox(boxes[i])
    }
}

function CreateBox(mapBox) {

    if (!game) return;

    let boxShadow;
    if (!mapBox.underground) {
        boxShadow = game.floorObjectLayer.create(mapBox.x + game.shadowXOffset, mapBox.y + game.shadowYOffset, mapBox.type);
        boxShadow.anchor.setTo(0.5);
        boxShadow.scale.set(0.1);
        boxShadow.tint = 0x000000;
        boxShadow.alpha = 0.4;
        boxShadow.angle = mapBox.rotate;
        mapBox.shadow = boxShadow;
    }

    let box = game.floorObjectLayer.create(mapBox.x, mapBox.y, mapBox.type);
    box.anchor.setTo(0.5);
    box.scale.set(0.1);
    box.angle = mapBox.rotate;
    mapBox.sprite = box;

    if (boxShadow) {
        box.shadow = boxShadow;
    }

    box.inputEnabled = true;
    box.input.pixelPerfectOver = true;
    box.input.pixelPerfectClick = true;

    box.events.onInputOver.add(function () {
        if (!box.border) {
            box.border = CreateBorder(mapBox.x, mapBox.y, mapBox.type, 20, mapBox.rotate, 0, 0, game.floorObjectLayer);
            game.floorObjectLayer.swap(box, box.border);
        } else {
            box.border.visible = true;
        }
    });

    box.events.onInputOut.add(function () {
        if (box.border) box.border.visible = false;
    });

    box.events.onInputDown.add(function () {
        for (let i in selectUnits) {
            let selectedUnit = game.units[selectUnits[i].id];
            if (selectedUnit && selectedUnit.owner_id === game.user_id && selectedUnit.body.mother_ship) {
                selectedUnit.toBox = {
                    boxID: mapBox.id,
                    to: true,
                    x: box.x,
                    y: box.y
                }
            }
        }
    });

    game.boxes.push(mapBox);
}