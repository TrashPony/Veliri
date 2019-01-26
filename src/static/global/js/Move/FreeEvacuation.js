function FreeMoveEvacuation(jsonData) {

    if (!(game && game.bases && game.bases[jsonData.base_id] && game.bases[jsonData.base_id].transports)) {
        return
    }

    CreateMiniMap();

    let sprite = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
    if (!sprite) {
        sprite = CreateEvacuation(jsonData.path_unit.x, jsonData.path_unit.y, jsonData.base_id, jsonData.transport_id);
        EvacuationUp(sprite)
    }

    game.add.tween(sprite).to({
            x: jsonData.path_unit.x,
            y: jsonData.path_unit.y
        }, 1000, Phaser.Easing.Linear.None, true, 0
    );

    game.add.tween(sprite.shadow).to({
            x: jsonData.path_unit.x + game.shadowXOffset * 10,
            y: jsonData.path_unit.y + game.shadowYOffset * 10
        }, 1000, Phaser.Easing.Linear.None, true, 0
    );
}