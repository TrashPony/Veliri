function Evacuation() {
    global.send(JSON.stringify({
        event: "evacuation"
    }));
}

function startMoveEvacuation(jsonData) {
    FocusMS();

    let x = jsonData.path_unit.x;
    let y = jsonData.path_unit.y;

    let tween = game.add.tween(game.squad.sprite.unitBody).to({alpha: 0.6}, 200, "Linear", true, 0, -1);
    tween.yoyo(true, 1000);

    setTimeout(function () {
        game.camera.focusOnXY(x * game.camera.scale.x, y * game.camera.scale.y);
        CreateMiniMap(game.map);

        let shadow = game.flyObjectsLayer.create(x + game.shadowXOffset, y + game.shadowYOffset, 'evacuation');
        shadow.anchor.setTo(0.5);
        shadow.scale.set(0.2);
        shadow.alpha = 0;
        shadow.tint = 0x000000;

        let evacuation = game.flyObjectsLayer.create(x, y, 'evacuation');
        evacuation.anchor.setTo(0.5);
        evacuation.scale.set(0.2);
        evacuation.alpha = 0;

        evacuation.shadow = shadow;
        evacuation.id = jsonData.base_id;

        game.evacuations.push(evacuation);

        game.add.tween(shadow).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);
        game.add.tween(evacuation).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);
        game.add.tween(shadow).to({alpha: 0.3}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(evacuation).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);

        setTimeout(function () {
            game.add.tween(shadow).to({
                x: x + game.shadowXOffset * 10,
                y: y + game.shadowYOffset * 10
            }, 1000, Phaser.Easing.Linear.None, true, 0);
            game.add.tween(shadow.scale).to({x: 0.3, y: 0.3}, 1000, Phaser.Easing.Linear.None, true, 0);
            game.add.tween(evacuation.scale).to({x: 0.3, y: 0.3}, 1000, Phaser.Easing.Linear.None, true, 0);
        }, 700);

    }, 500);
}

function evacuationMove(jsonData, squad) {
    for (let i = 0; i < game.evacuations.length; i++) {
        if (game.evacuations[i].id === jsonData.base_id) {
            game.add.tween(game.evacuations[i]).to({
                    x: jsonData.path_unit.x,
                    y: jsonData.path_unit.y
                }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
            );

            game.add.tween(game.evacuations[i].shadow).to({
                    x: jsonData.path_unit.x + game.shadowXOffset * 10,
                    y: jsonData.path_unit.y + game.shadowYOffset * 10
                }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
            );

            if (squad) {
                game.add.tween(squad).to({
                        x: jsonData.path_unit.x,
                        y: jsonData.path_unit.y
                    }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
                );
            }
        }
    }
}

function stopEvacuation(jsonData) {
    for (let i = 0; i < game.evacuations.length; i++) {
        if (game.evacuations[i].id === jsonData.base_id) {
            game.add.tween(game.evacuations[i].shadow).to({
                x: game.evacuations[i].x + game.shadowXOffset,
                y: game.evacuations[i].y + game.shadowYOffset
            }, 1000, Phaser.Easing.Linear.None, true, 0);

            game.add.tween(game.evacuations[i].shadow.scale).to({
                x: 0.2,
                y: 0.2
            }, 1000, Phaser.Easing.Linear.None, true, 0);
            game.add.tween(game.evacuations[i].scale).to({x: 0.2, y: 0.2}, 1000, Phaser.Easing.Linear.None, true, 0);

            game.add.tween(game.squad.sprite.bodyShadow).to({
                x: game.shadowXOffset,
                y: game.shadowYOffset
            }, 1000, Phaser.Easing.Linear.None, true, 0);
            game.add.tween(game.squad.sprite.scale).to({x: 1, y: 1}, 1000, Phaser.Easing.Linear.None, true, 0);
        }
    }
}

function placeEvacuation(jsonData) {
    for (let i = 0; i < game.evacuations.length; i++) {
        if (game.evacuations[i].id === jsonData.base_id) {

            game.add.tween(game.evacuations[i].shadow).to({
                x: game.evacuations[i].x + game.shadowXOffset,
                y: game.evacuations[i].y + game.shadowYOffset
            }, 1000, Phaser.Easing.Linear.None, true, 0);

            game.add.tween(game.evacuations[i].shadow.scale).to({
                x: 0.2,
                y: 0.2
            }, 1000, Phaser.Easing.Linear.None, true, 0);
            game.add.tween(game.evacuations[i].scale).to({x: 0.2, y: 0.2}, 1000, Phaser.Easing.Linear.None, true, 0);

            setTimeout(function () {
                game.add.tween(game.evacuations[i].shadow).to({
                    x: game.evacuations[i].x + game.shadowXOffset * 10,
                    y: game.evacuations[i].y + game.shadowYOffset * 10
                }, 1000, Phaser.Easing.Linear.None, true, 0);
                game.add.tween(game.evacuations[i].shadow.scale).to({
                    x: 0.3,
                    y: 0.3
                }, 1000, Phaser.Easing.Linear.None, true, 0);
                game.add.tween(game.evacuations[i].scale).to({
                    x: 0.3,
                    y: 0.3
                }, 1000, Phaser.Easing.Linear.None, true, 0);

                game.add.tween(game.squad.sprite.bodyShadow).to({
                    x: game.shadowXOffset * 10,
                    y: game.shadowYOffset * 10
                }, 1000, Phaser.Easing.Linear.None, true, 0);
                game.add.tween(game.squad.sprite.scale).to({x: 1.1, y: 1.1}, 1000, Phaser.Easing.Linear.None, true, 0);

                game.squad.sprite.x = game.evacuations[i].x;
                game.squad.sprite.y = game.evacuations[i].y;
            }, 1200)
        }
    }
}