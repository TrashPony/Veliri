function CreateObjects(coordinate, x, y) {
    let object;

    if (coordinate.texture_object === "terrain_1") {
            object = gameObjectCreate(x, y, coordinate.texture_object, -0.2, 0.2, -0.47, 0.85, 45);
    }

    if (coordinate.texture_object === "terrain_2") {
        object = gameObjectCreate(x, y, coordinate.texture_object, 0, 0.2, -0.42, 0.75, 45);
    }

    if (coordinate.texture_object === "terrain_3") {
        object = gameObjectCreate(x, y, coordinate.texture_object, 0, 0.2, -0.33, 0.95, 45);
    }

    if (coordinate.texture_object === "sand_stone_04") {
        object = gameObjectCreate(x, y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "sand_stone_05") {
        object = gameObjectCreate(x, y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "sand_stone_06") {
        object = gameObjectCreate(x, y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "sand_stone_07") {
        object = gameObjectCreate(x, y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "sand_stone_08") {
        object = gameObjectCreate(x, y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "crater") {
        object = game.floorObjectLayer.create(x - 90, y, coordinate.texture_object);
        object.inputEnabled = true;
        object.events.onInputOut.add(TipOff);
        object.events.onInputDown.add(RemoveSelect);
        object.input.pixelPerfectOver = true;
        object.input.pixelPerfectClick = true;
    }

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture, spriteAnchorX, spriteAnchorY, ShadowOffsetX, ShadowOffsetY, angle) {

    let shadow = game.floorObjectLayer.create(x - 100, y + game.shadowYOffset, texture);
    shadow.anchor.setTo(ShadowOffsetX, ShadowOffsetY);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;
    shadow.angle = angle;

    shadow.inputEnabled = true;             // включаем ивенты на спрайт
    shadow.input.pixelPerfectOver = true;
    shadow.input.pixelPerfectClick = true;

    let object = game.floorObjectLayer.create(x - 100, y, texture);
    object.inputEnabled = true;
    object.events.onInputOut.add(TipOff);
    object.events.onInputDown.add(RemoveSelect);
    object.anchor.setTo(spriteAnchorX, spriteAnchorY); // todo подгружать спрайт от невидимой границы

    shadow.inputEnabled = true;
    object.input.pixelPerfectOver = true;
    object.input.pixelPerfectClick = true;

    object.shadow = shadow;

    return object
}