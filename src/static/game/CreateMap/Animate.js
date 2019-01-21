function CreateAnimate(coordinate, x, y) {
    let animate;

    if (coordinate.impact) {
        return
    }
    if (coordinate.unit_overlap) {
        animate = gameAnimateObjectCreate(x, y, coordinate.animate_sprite_sheets, coordinate.scale, coordinate.shadow,
            coordinate.obj_rotate, coordinate.animation_speed, coordinate.x_offset, coordinate.y_offset, game.floorObjectLayer);
    } else {
        animate = gameAnimateObjectCreate(x, y, coordinate.animate_sprite_sheets, coordinate.scale, coordinate.shadow,
            coordinate.obj_rotate, coordinate.animation_speed, coordinate.x_offset, coordinate.y_offset, game.floorOverObjectLayer);
    }

    coordinate.objectSprite = animate;
}

function gameAnimateObjectCreate(x, y, texture, scale, needShadow, rotate, speed, xOffset, yOffset, group) {

    let object = game.floorObjectLayer.create(x + xOffset, y + yOffset, texture);
    object.anchor.setTo(0.5, 0.5);
    object.scale.set((scale / 100) / 2);
    object.angle = rotate;

    object.animations.add('objAnimate');
    object.animations.play('objAnimate', speed, true);

    if (needShadow) {
        let shadow = game.floorObjectLayer.create(x + game.shadowXOffset + xOffset, y - game.shadowYOffset + 20 + yOffset, texture);
        shadow.anchor.setTo(0.5, 0.5);
        shadow.scale.set((scale / 100) / 2);
        shadow.tint = 0x000000;
        shadow.alpha = 0.6;
        shadow.angle = rotate;

        shadow.animations.add('objAnimate');
        shadow.animations.play('objAnimate', speed, true);

        object.shadow = shadow;
    }

    return object
}