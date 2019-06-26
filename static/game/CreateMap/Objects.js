function CreateObject(coordinate, x, y) {
    let object;

    if (coordinate.impact) {
        return
    }

    if (coordinate.unit_overlap) {
        object = gameObjectCreate(x, y, coordinate.texture_object, coordinate.scale, coordinate.shadow, coordinate.obj_rotate,
            coordinate.x_offset, coordinate.y_offset, game.floorOverObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset,
            coordinate.shadow_intensity);
    } else {
        object = gameObjectCreate(x, y, coordinate.texture_object, coordinate.scale, coordinate.shadow, coordinate.obj_rotate,
            coordinate.x_offset, coordinate.y_offset, game.floorObjectLayer, coordinate.x_shadow_offset, coordinate.y_shadow_offset,
            coordinate.shadow_intensity);
    }

    coordinate.objectSprite = object;
}

function gameObjectCreate(x, y, texture, scale, needShadow, rotate, xOffset, yOffset, group, xShadowOffset, yShadowOffset, shadowIntensity) {
    let shadow;

    if (needShadow) {
        shadow = group.create(x + xOffset + game.shadowXOffset + xShadowOffset, y + yOffset + game.shadowYOffset + yShadowOffset, texture);
        shadow.anchor.setTo(0.5, 0.5);
        shadow.scale.set((scale / 100) / 2);
        shadow.tint = 0x000000;
        shadow.angle = rotate;
        shadow.alpha = shadowIntensity / 100;
    }

    let object = group.create(x + xOffset, y + yOffset, texture);
    object.anchor.setTo(0.5, 0.5);
    object.scale.set((scale / 100) / 2);
    object.angle = rotate;

    object.shadow = shadow;

    return object
}