function CreateObjects(coordinate, x, y) {
    let object;

    if (coordinate.impact) {
        return
    }

    if (coordinate.unit_overlap) {
        object = gameObjectCreate(x, y, coordinate.texture_object, coordinate.scale,
            coordinate.shadow, coordinate.obj_rotate, coordinate.x_offset, coordinate.y_offset, game.floorOverObjectLayer);
    } else {
        object = gameObjectCreate(x, y, coordinate.texture_object, coordinate.scale,
            coordinate.shadow, coordinate.obj_rotate, coordinate.x_offset, coordinate.y_offset, game.floorObjectLayer);
    }

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture, scale, needShadow, rotate, xOffset, yOffset, group) {

    let object = group.create(x + xOffset, y + yOffset, texture);
    object.anchor.setTo(0.5, 0.5);
    object.scale.set(scale / 100);
    object.angle = rotate;


    if (needShadow) {
        let shadow = group.create(x + game.shadowXOffset + xOffset, y - game.shadowYOffset + 20 + yOffset, texture);
        shadow.anchor.setTo(0.5, 0.5);
        shadow.scale.set(scale / 100);
        shadow.tint = 0x000000;
        shadow.angle = rotate;
        shadow.alpha = 0.4;

        object.shadow = shadow;
    }

    return object
}