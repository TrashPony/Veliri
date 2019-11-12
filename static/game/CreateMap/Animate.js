function CreateAnimate(coordinate, x, y) {
    let animate;

    if (coordinate.impact) {
        return
    }
    if (coordinate.unit_overlap) {
        animate = gameAnimateObjectCreate(x, y, coordinate.animate_sprite_sheets, coordinate.scale, coordinate.shadow,
            coordinate.rotate, coordinate.animation_speed, game.floorOverObjectLayer, coordinate.animate_loop);
    } else {
        animate = gameAnimateObjectCreate(x, y, coordinate.animate_sprite_sheets, coordinate.scale, coordinate.shadow,
            coordinate.rotate, coordinate.animation_speed, game.floorObjectLayer, coordinate.animate_loop);
    }

    if (game.typeService !== "mapEditor") {
        ObjectEvents(coordinate, animate);
    }

    coordinate.objectSprite = animate;
}

function gameAnimateObjectCreate(x, y, texture, scale, needShadow, rotate, speed, group, needAnimate) {
    let shadow;
    if (needShadow) {
        shadow = group.create(x + game.shadowXOffset, y + game.shadowYOffset, texture);
        shadow.anchor.setTo(0.5, 0.5);
        shadow.scale.set((scale / 100) / 2);
        shadow.tint = 0x000000;
        shadow.alpha = 0.4;
        shadow.angle = rotate;

        if (needAnimate) {
            shadow.animations.add('objAnimate');
            shadow.animations.play('objAnimate', speed, true);
        }
    }

    let object = group.create(x, y, texture);
    object.anchor.setTo(0.5, 0.5);
    object.scale.set((scale / 100) / 2);
    object.angle = rotate;

    if (needAnimate) {
        object.animations.add('objAnimate');
        object.animations.play('objAnimate', speed, true);
    }

    object.shadow = shadow;

    return object
}