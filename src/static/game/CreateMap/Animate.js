function CreateAnimate(coordinate) {
    let animate;

    if (coordinate.impact) {
        return
    }

    animate = gameAnimateObjectCreate(coordinate.sprite.x, coordinate.sprite.y, coordinate.animate_sprite_sheets, coordinate.scale,
        coordinate.shadow, coordinate.obj_rotate);

    coordinate.objectSprite = animate;
}

function gameAnimateObjectCreate(x, y, texture, scale, needShadow, rotate) {

    let object = game.floorObjectLayer.create(x, y, texture);
    object.anchor.setTo(0.5, 0.5);
    object.scale.set(scale / 100);
    object.angle = rotate;

    object.animations.add('objAnimate');
    object.animations.play('objAnimate', 15, true);

    if (needShadow) {
        let shadow = game.floorObjectLayer.create(x + game.shadowXOffset, y - game.shadowYOffset + 20, texture);
        shadow.anchor.setTo(0.5, 0.5);
        shadow.scale.set(scale / 100);
        shadow.tint = 0x000000;
        shadow.alpha = 0.6;
        shadow.angle = rotate;

        shadow.animations.add('objAnimate');
        shadow.animations.play('objAnimate', 15, true);

        object.shadow = shadow;
    }

    return object
}