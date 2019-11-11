function FreeMoveEvacuation(jsonData) {
    if (!(game && game.bases && game.bases[jsonData.base_id] && game.bases[jsonData.base_id].transports)) {
        return
    }

    CreateMiniMap();

    let sprite = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
    if (!sprite) {
        //sprite = CreateEvacuation(jsonData.path_unit.x, jsonData.path_unit.y, jsonData.base_id, jsonData.transport_id);
        //QuickUpEvacuation(sprite, jsonData.path_unit.rotate)
        return
    }

    game.add.tween(sprite).to({
            x: jsonData.path_unit.x,
            y: jsonData.path_unit.y
        }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
    );

    game.add.tween(sprite.shadow).to({
            x: jsonData.path_unit.x + game.shadowXOffset * 10,
            y: jsonData.path_unit.y + game.shadowYOffset * 10
        }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
    );

    ShortDirectionRotateTween(sprite, Phaser.Math.degToRad(jsonData.path_unit.rotate), jsonData.path_unit.millisecond);
    ShortDirectionRotateTween(sprite.shadow, Phaser.Math.degToRad(jsonData.path_unit.rotate), jsonData.path_unit.millisecond);
}

function QuickUpEvacuation(sprite, rotate) {
    sprite.shadow.x = sprite.x + game.shadowXOffset * 10;
    sprite.shadow.y = sprite.y + game.shadowXOffset * 10;

    sprite.shadow.scale.x = 0.15;
    sprite.shadow.scale.y = 0.15;

    sprite.scale.x = 0.15;
    sprite.scale.y = 0.15;

    sprite.angle = rotate;
    sprite.shadow.angle = rotate;
}