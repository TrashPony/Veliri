function update() {

    dynamicMap(game.floorLayer, game.mapPoints);

    if (game && game.typeService === "battle") {
        UpdateRotateUnit(); // функция для повора юнитовский спрайтов
        MoveUnit();
        FlightBullet(); // ослеживает все летящие спрайты пуль
    }

    if (game && game.typeService === "global") {
        if (game.squad && game.squad.toMove) {
            // TODO анимация движения, следы от гусей, выхлоп и тд
            // TODO отрисовка линии маршрута
            let rotate = Math.atan2(game.squad.toMove.y - game.squad.sprite.world.y, game.squad.toMove.x - game.squad.sprite.world.x);

            let spriteAngle = game.squad.sprite.unitBody.angle;
            let needAngle = (rotate * 180 / 3.14) + 90;

            let angleDiff;
            if (needAngle > 180) {
                needAngle -= 360;
            }

            if (spriteAngle >= needAngle) {
                angleDiff = spriteAngle - needAngle
            } else {
                angleDiff = needAngle - spriteAngle
            }

            let dist = game.physics.arcade.distanceToXY(game.squad.sprite, game.squad.toMove.x, game.squad.toMove.y);

            if (angleDiff > 1 && !game.squad.tweenTo) {
                rotateUnitSprites(spriteAngle, needAngle, game.squad);

                if (dist < 250 && angleDiff > 15) {
                    game.squad.sprite.body.angularVelocity = 0;
                    game.squad.sprite.body.velocity.x = 0;
                    game.squad.sprite.body.velocity.y = 0;
                } else {
                    game.squad.sprite.body.velocity.copyFrom(
                        game.physics.arcade.velocityFromAngle(spriteAngle - 90, game.squad.mather_ship.speed * 30)
                    );
                }

                if (Math.round(dist) >= -10 && Math.round(dist) <= 10) {
                    game.squad.sprite.body.angularVelocity = 0;
                    game.squad.sprite.body.velocity.x = 0;
                    game.squad.sprite.body.velocity.y = 0;
                    game.squad.toMove = null;
                }
            } else {
                game.squad.sprite.body.angularVelocity = 0;
                game.squad.sprite.body.velocity.x = 0;
                game.squad.sprite.body.velocity.y = 0;

                if (!game.squad.tweenTo) {
                    game.squad.tweenTo = game.add.tween(game.squad.sprite).to(
                        {x: game.squad.toMove.x, y: game.squad.toMove.y},
                        dist / game.squad.mather_ship.speed * 30,
                        Phaser.Easing.Linear.None, true, 0
                    );

                    game.squad.tweenTo.onComplete.add(function () {
                        game.squad.tweenTo = null; //todo возможно утечка памяти т.к. твис сам по себе не удаляется
                        game.squad.toMove = null;
                    })
                }
            }
        }
    }

    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */
    game.floorObjectLayer.sort('y', Phaser.Group.SORT_DESCENDING);
}