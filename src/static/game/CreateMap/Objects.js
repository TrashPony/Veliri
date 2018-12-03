function CreateObjects(coordinate) {
    let object;

    if (coordinate.impact) {
        return
    }

    object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object, coordinate.scale,
        coordinate.shadow, coordinate.obj_rotate, coordinate.x_offset, coordinate.y_offset);

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture, scale, needShadow, rotate, xOffset, yOffset) {

    let object = game.floorObjectLayer.create(x + xOffset, y + yOffset, texture);
    object.anchor.setTo(0.5, 0.5);
    object.scale.set(scale / 100);
    object.angle = rotate;

    if (needShadow) {
        let shadow = game.floorObjectLayer.create(x + game.shadowXOffset + xOffset, y - game.shadowYOffset + 20 + yOffset, texture);
        shadow.anchor.setTo(0.5, 0.5);
        shadow.scale.set(scale / 100);
        shadow.tint = 0x000000;
        shadow.alpha = 0.6;
        shadow.angle = rotate;

        object.shadow = shadow;
    }

    return object
}