function update() {

    DynamicShadowMap(game.floorLayer, game.mapPoints);

    if (game && game.typeService === "battle") {
        UpdateRotateUnit(); // функция для повора юнитовский спрайтов
        MoveUnit();
        FlightBullet(); // ослеживает все летящие спрайты пуль
    }

    if (game && game.typeService === "global") {
        if (game.squad && game.squad.toBase && game.squad.toBase.into) {
            let dist = game.physics.arcade.distanceToXY(game.squad.sprite, game.squad.toBase.x, game.squad.toBase.y);
            if (dist < 300) {
                global.send(JSON.stringify({
                    event: "IntoToBase",
                    base_id: game.squad.toBase.baseID
                }));
            }
        }

        if (game.squad && game.squad.toBox && game.squad.toBox.to) {
            let dist = game.physics.arcade.distanceToXY(game.squad.sprite, game.squad.toBox.x, game.squad.toBox.y);
            if (dist < 300) {
                global.send(JSON.stringify({
                    event: "openBox",
                    box_id: game.squad.toBox.boxID
                }));
            }
        }

        AnimateMiningLaser();
        AnimateDigger();
    }

    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    game.floorObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
    game.unitLayer.sort('y', Phaser.Group.SORT_DESCENDING);
    game.floorOverObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
}