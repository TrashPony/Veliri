function CreateObjects(coordinate, x, y) {
    let object;

    if (coordinate.impact) {
        return
    }

    object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object, coordinate.scale, coordinate.shadow);

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture, scale, needShadow) {

    let object = game.floorObjectLayer.create(x, y, texture);
    object.anchor.setTo(0.5, 0.5);
    object.scale.set(scale/100);

    if (needShadow) {
        let shadow = game.floorObjectLayer.create(x + game.shadowXOffset, y - game.shadowYOffset + 20, texture);
        shadow.anchor.setTo(0.5, 0.5);
        shadow.scale.set(scale/100);
        shadow.tint = 0x000000;
        shadow.alpha = 0.6;

        object.shadow = shadow;
    }

    return object
}