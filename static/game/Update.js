function update() {
    GrabCamera(); // функцуия для перетаскивания карты мышкой /* Магия */

    // if (game && game.mapPoints) {
    //     // todo идея хорошая реализация нет dynamicMap(game.floorLayer, game.mapPoints);
    //     // todo идея хорошая реализация нет DynamicShadowMap();
    // }

    if (game && game.typeService === "battle") {
        MoveUnit();
        FlightBullet(); // ослеживает все летящие спрайты пуль
    }

    if (game && game.typeService === "global") {
        //
        // AnimateDigger();
        //
        //
        // DebugCollision();
        //

        for (let i in game.units){
            let unit = game.units[i];
            AnimationMove(unit);

            if (unit && unit.toBox && unit.toBox.to) {
                let dist = game.physics.arcade.distanceToXY(unit.sprite, unit.toBox.x, unit.toBox.y);
                if (dist < 100) {
                    global.send(JSON.stringify({
                        event: "openBox",
                        box_id: unit.toBox.boxID
                    }));
                }
            }
        }
        StartSelectableUnits();
        AnimateMiningLaser();
    }
}

function DebugCollision() {
    if (game.squad && game.squad.colision) {
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
                if (reservoir && reservoir.sprite) {
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
}