function CreateObjects(coordinate, x, y) {
    let object;

    if (coordinate.texture_object === "terrain_1") {
        object = gameObjectCreate(x, y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "terrain_2") {
        object = gameObjectCreate(x, y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "terrain_3") {
        object = gameObjectCreate(x, y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_04") {
        object = gameObjectCreate(x, y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_05") {
        object = gameObjectCreate(x, y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_06") {
        object = gameObjectCreate(x, y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_07") {
        object = gameObjectCreate(x, y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_08") {
        object = gameObjectCreate(x, y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "fallen_01") {
        object = game.floorObjectLayer.create(x - 90, y + 10, coordinate.texture_object);
    }

    if (coordinate.texture_object === "fallen_02") {
        object = game.floorObjectLayer.create(x - 95, y + 10, coordinate.texture_object);
    }

    if (coordinate.texture_object === "crater") {
        object = game.floorObjectLayer.create(x - 90, y + 17, coordinate.texture_object);
    }

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture) {

    let object = game.floorObjectLayer.create(x - 90, y + 10, texture);
    object.anchor.setTo(0, 0);

    let shadow = game.floorObjectLayer.create(x - 90 + game.shadowXOffset, y - game.shadowYOffset + 25, texture);
    shadow.anchor.setTo(0, 0);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;

    object.shadow = shadow;

    return object
}