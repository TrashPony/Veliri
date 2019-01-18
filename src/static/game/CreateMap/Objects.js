function CreateObjects(coordinate) {
    let object;

    if (coordinate.impact) {
        return
    }

    if (coordinate.unit_overlap) {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object, coordinate.scale,
            coordinate.shadow, coordinate.obj_rotate, coordinate.x_offset, coordinate.y_offset, game.floorOverObjectLayer);
    } else {
        object = gameObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.texture_object, coordinate.scale,
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
        let shadow = group.create(0, 0, texture);
        shadow.anchor.setTo(0.5, 0.5);
        shadow.scale.set(scale / 100);
        shadow.tint = 0x000000;
        shadow.angle = rotate;

        game.bmdShadow.draw(shadow, x + game.shadowXOffset + xOffset, y - game.shadowYOffset + 20 + yOffset);
        shadow.destroy();
    }

    return object
}