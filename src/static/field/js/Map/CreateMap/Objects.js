function CreateObjects(coordinate) {
    var object;

    if (coordinate.texture_object === "terrain_1") {
        object = gameObjectCreate(coordinate.x, coordinate.y, coordinate.texture_object, 0, 0.2, -0.33, 0.70, 45);
    }

    if (coordinate.texture_object === "terrain_2") {
        object = gameObjectCreate(coordinate.x, coordinate.y, coordinate.texture_object, 0, 0.2, -0.42, 0.75, 45);
    }

    if (coordinate.texture_object === "terrain_3") {
        object = gameObjectCreate(coordinate.x, coordinate.y, coordinate.texture_object, 0, 0, -0.5, 0.8, 45);
    }

    if (coordinate.texture_object === "sand_stone_04") {
        object = gameObjectCreate(coordinate.x, coordinate.y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "sand_stone_05") {
        object = gameObjectCreate(coordinate.x, coordinate.y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "sand_stone_06") {
        object = gameObjectCreate(coordinate.x, coordinate.y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "sand_stone_07") {
        object = gameObjectCreate(coordinate.x, coordinate.y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "sand_stone_08") {
        object = gameObjectCreate(coordinate.x, coordinate.y, coordinate.texture_object, 0, 0.2, -0.1, 0.35, 10);
    }

    if (coordinate.texture_object === "crater") {
        object = game.floorObjectLayer.create(coordinate.x * game.tileSize, coordinate.y * game.tileSize, coordinate.texture_object);
        object.inputEnabled = true;
        object.events.onInputOut.add(TipOff);
        object.events.onInputDown.add(RemoveSelect);
        object.input.pixelPerfectOver = true;
        object.input.pixelPerfectClick = true;
    }

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture, spriteAnchorX, spriteAnchorY, ShadowOffsetX, ShadowOffsetY, angle) {

    var shadow = game.floorObjectLayer.create((x * game.tileSize), (y * game.tileSize) + game.shadowYOffset, texture);
    shadow.anchor.setTo(ShadowOffsetX, ShadowOffsetY);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;
    shadow.angle = angle;

    shadow.inputEnabled = true;             // включаем ивенты на спрайт
    shadow.input.pixelPerfectOver = true;
    shadow.input.pixelPerfectClick = true;

    var object = game.floorObjectLayer.create(x * game.tileSize, y * game.tileSize, texture);
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