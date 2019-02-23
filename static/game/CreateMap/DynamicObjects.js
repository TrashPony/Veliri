function CreateDynamicObjects(dynamicObject, q, r, needBackGround, coordinate) {
    if (coordinate && !coordinate.dynamicObjects) {
        coordinate.dynamicObjects = [];
    }

    let dynamicObjectSprite = {background: undefined, object: undefined, shadow: undefined};

    let xy = GetXYCenterHex(q, r);
    if (needBackGround && dynamicObject.texture_background !== '') {
        let background = game.floorLayer.create(xy.x, xy.y, dynamicObject.texture_background);
        background.scale.set((dynamicObject.background_scale / 100)/2);
        background.angle = dynamicObject.background_rotate;
        background.anchor.setTo(0.5);
        dynamicObjectSprite.background = background;
    }

    let objectSprite = game.floorObjectLayer.create(xy.x, xy.y, dynamicObject.texture_object);
    objectSprite.scale.set((dynamicObject.object_scale / 100)/2);
    objectSprite.angle = dynamicObject.object_rotate;
    objectSprite.anchor.setTo(0.5);
    dynamicObjectSprite.object = objectSprite;

    if (dynamicObject.shadow > 0) {

        let shadow = game.floorObjectLayer.create(
            xy.x + game.shadowXOffset * (dynamicObject.shadow / 100),
            xy.y + game.shadowYOffset * (dynamicObject.shadow / 100),
            dynamicObject.texture_object
        );

        shadow.scale.set((dynamicObject.object_scale / 100)/2);
        shadow.angle = dynamicObject.object_rotate;
        shadow.anchor.setTo(0.5);
        shadow.tint = 0x000000;
        shadow.alpha = 0.6;
        dynamicObjectSprite.shadow = shadow;
    }

    if (dynamicObject.dialog) {
        objectSprite.inputEnabled = true;
        objectSprite.events.onInputDown.add(function () {
            anomalyText(dynamicObject.dialog, dynamicObject.dialog.pages[0])
        });

        let paketLine;
        objectSprite.events.onInputOver.add(function () {
            paketLine = game.floorObjectSelectLineLayer.create(xy.x, xy.y, dynamicObject.texture_object);
            paketLine.anchor.setTo(0.5);
            paketLine.scale.set(0.22);
            paketLine.angle = dynamicObject.object_rotate;
            paketLine.tint = 0x00FF00;
        });
        objectSprite.events.onInputOut.add(function () {
            paketLine.destroy();
        });
        objectSprite.input.priorityID = 1;
    }

    if (coordinate) {
        coordinate.dynamicObjects.push(dynamicObjectSprite);
    }
}