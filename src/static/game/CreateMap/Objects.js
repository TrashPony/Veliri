function CreateObjects(coordinate, x, y) {
    let object;

    if (coordinate.impact) {
        return
    }

    if (coordinate.texture_object === "terrain_1") {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "terrain_2") {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "terrain_3") {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_04") {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_05") {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_06") {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_07") {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "sand_stone_08") {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
    }

    if (coordinate.texture_object === "fallen_01") {
        object = game.floorObjectLayer.create(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
        object.anchor.setTo(0.5, 0.5);
        object.scale.set(0.4);
    }

    if (coordinate.texture_object === "fallen_02") {
        object = game.floorObjectLayer.create(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
        object.anchor.setTo(0.5, 0.5);
        object.scale.set(0.4);
    }

    if (coordinate.texture_object === "crater") {
        object = game.floorObjectLayer.create(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
        object.anchor.setTo(0.5, 0.5);
        object.scale.set(0.4);
    }

    if (coordinate.texture_object === "base") {
        object = game.floorObjectLayer.create(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object);
        object.anchor.setTo(0.5, 0.5);
        object.scale.set(0.4);
    }

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture) {

    let object = game.floorObjectLayer.create(x, y, texture);
    object.anchor.setTo(0.5, 0.5);
    object.scale.set(0.4);

    let shadow = game.floorObjectLayer.create(x + game.shadowXOffset, y - game.shadowYOffset + 20, texture);
    shadow.anchor.setTo(0.5, 0.5);
    shadow.scale.set(0.4);
    shadow.tint = 0x000000;
    shadow.alpha = 0.6;

    object.shadow = shadow;

    return object
}