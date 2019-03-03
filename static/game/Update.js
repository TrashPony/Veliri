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


        /* DEBAG COLLISION */
        if (game.squad.colision) {
            CreateCollision(game.squad.colision, game.squad.mather_ship.body, game.squad.mather_ship.rotate, game.squad);
            for (let i = 0; i < game.boxes.length; i++) {
                game.squad.colision.beginFill(0xFF0000, 0.5);
                if (game.boxes[i] && game.boxes[i].sprite) {
                    game.squad.colision.drawCircle(game.boxes[i].sprite.x, game.boxes[i].sprite.y, 10);
                }
            }

            for (let q in game.map.reservoir) {
                for (let r in game.map.reservoir[q]) {
                    let reservoir = game.map.reservoir[q][r];
                    if(reservoir && reservoir.sprite){
                        game.squad.colision.drawCircle(reservoir.sprite.x, reservoir.sprite.y, 30);
                    }
                }
            }
        }

        for (let i = 0; game.otherUsers && i < game.otherUsers.length; i++) {
            if (game.otherUsers[i].colision) {
                CreateCollision(game.otherUsers[i].colision, game.otherUsers[i].body, game.otherUsers[i].rotate, game.otherUsers[i])
            }
        }
        /* DEBAG COLLISION */
    }

    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    game.floorObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
    game.unitLayer.sort('y', Phaser.Group.SORT_DESCENDING);
    game.floorOverObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
}