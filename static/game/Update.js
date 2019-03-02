function update() {

    // if (game && game.mapPoints) {
    //     // todo идея хорошая реализация нет dynamicMap(game.floorLayer, game.mapPoints);
    //     // todo идея хорошая реализация нет DynamicShadowMap();
    // }

    if (game && game.typeService === "battle") {
        UpdateRotateUnit(); // функция для повора юнитовский спрайтов
        MoveUnit();
        FlightBullet(); // ослеживает все летящие спрайты пуль
    }

    if (game && game.typeService === "global") {
        // if (game.squad && game.squad.toBase && game.squad.toBase.into) {
        //     let dist = game.physics.arcade.distanceToXY(game.squad.sprite, game.squad.toBase.x, game.squad.toBase.y);
        //         if (dist < 150) {
        //         global.send(JSON.stringify({
        //             event: "IntoToBase",
        //             base_id: game.squad.toBase.baseID
        //         }));
        //     }
        // }

        if (game.squad && game.squad.toBox && game.squad.toBox.to) {
            let dist = game.physics.arcade.distanceToXY(game.squad.sprite, game.squad.toBox.x, game.squad.toBox.y);
            if (dist < 150) {
                global.send(JSON.stringify({
                    event: "openBox",
                    box_id: game.squad.toBox.boxID
                }));
            }
        }

        AnimateMiningLaser();
        AnimateDigger();

        if (game.squad.colision) {
            CreateCollision(game.squad.colision, game.squad.mather_ship.body, game.squad.mather_ship.rotate, game.squad)
        }

        for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
            if (game.otherUsers[i].colision) {
                CreateCollision(game.otherUsers[i].colision, game.otherUsers[i].body, game.otherUsers[i].rotate, game.otherUsers[i])
            }
        }
    }

    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    game.floorObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
    game.unitLayer.sort('y', Phaser.Group.SORT_DESCENDING);
    game.floorOverObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
}