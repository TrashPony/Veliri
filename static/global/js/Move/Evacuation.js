function Evacuation() {
    global.send(JSON.stringify({
        event: "evacuation"
    }));
}

function CreateEvacuation(x, y, baseId, transportId) {
    let shadow = game.flyObjectsLayer.create(x + game.shadowXOffset, y + game.shadowYOffset, 'evacuation');
    shadow.anchor.setTo(0.5);
    shadow.scale.set(0.1);
    shadow.alpha = 0;
    shadow.tint = 0x000000;

    let evacuation = game.flyObjectsLayer.create(x, y, 'evacuation');
    evacuation.anchor.setTo(0.5);
    evacuation.scale.set(0.1);
    evacuation.alpha = 0;
    evacuation.shadow = shadow;

    game.bases[baseId].transports[transportId].sprite = evacuation;

    game.add.tween(evacuation.shadow).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);
    game.add.tween(evacuation).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);

    game.add.tween(evacuation.shadow).to({alpha: 0.1}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(evacuation).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);

    return evacuation
}

function startMoveEvacuation(jsonData) {

    let x = jsonData.path_unit.x;
    let y = jsonData.path_unit.y;

    if (game.squad.id === jsonData.other_user.squad_id) {
        FocusMS();
        let tween = game.add.tween(game.squad.sprite.unitBody).to({alpha: 0.6}, 200, "Linear", true, 0, -1);
        tween.yoyo(true, 1000);
    } else {
        for (let j = 0; j < game.otherUsers.length; j++) {
            if (game.otherUsers[j].squad_id === jsonData.other_user.squad_id) {
                let tween = game.add.tween(game.otherUsers[j].sprite.unitBody).to({alpha: 0.6}, 200, "Linear", true, 0, -1);
                tween.yoyo(true, 1000);
            }
        }
    }

    setTimeout(function () {
        CreateMiniMap(game.map);

        let evacuation = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
        if (!evacuation) {
            console.log(jsonData);
            evacuation = CreateEvacuation(x, y, jsonData.base_id, jsonData.transport_id);
        }

        if (game.squad.id === jsonData.other_user.squad_id) {
            game.camera.follow(evacuation);
        }

        setTimeout(function () {
            EvacuationUp(evacuation)
        }, 700);
    }, 500);
}

function EvacuationUp(sprite, squad) {
    game.add.tween(sprite.shadow).to({
        x: sprite.x + game.shadowXOffset * 10,
        y: sprite.y + game.shadowYOffset * 10
    }, 1000, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(sprite.shadow.scale).to({x: 0.15, y: 0.15}, 1000, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(sprite.scale).to({x: 0.15, y: 0.15}, 1000, Phaser.Easing.Linear.None, true, 0);

    if (squad) {
        game.unitLayer.remove(squad);
        game.flyObjectsLayer.add(squad);
        game.flyObjectsLayer.swap(sprite, squad);

        game.add.tween(squad.bodyShadow).to({
            x: game.shadowXOffset * 10,
            y: game.shadowYOffset * 10
        }, 1000, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(squad.scale).to({x: 1.1, y: 1.1}, 1000, Phaser.Easing.Linear.None, true, 0);
        squad.x = sprite.x;
        squad.y = sprite.y;
    }
}

function EvacuationDown(sprite, squad, destroy) {
    let tween = game.add.tween(sprite.shadow).to({
        x: sprite.x + game.shadowXOffset,
        y: sprite.y + game.shadowYOffset
    }, 1000, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(sprite.shadow.scale).to({x: 0.1, y: 0.1}, 1000, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(sprite.scale).to({x: 0.1, y: 0.1}, 1000, Phaser.Easing.Linear.None, true, 0);

    if (squad) {
        game.add.tween(squad.bodyShadow).to({
            x: game.shadowXOffset,
            y: game.shadowYOffset
        }, 1000, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(squad.scale).to({x: 1, y: 1}, 1000, Phaser.Easing.Linear.None, true, 0);
    }

    if (destroy) {
        tween.onComplete.add(function () {
            EvacuationUp(sprite)
        })
    }
}

function evacuationMove(jsonData, squadMove) {
    if (game.map) {
        CreateMiniMap();
    }

    let squad;

    if (jsonData.other_user) {
        if (game.squad.id === jsonData.other_user.squad_id) {
            squad = game.squad.sprite
        } else {
            for (let j = 0; j < game.otherUsers.length; j++) {
                if (game.otherUsers[j].squad_id === jsonData.other_user.squad_id) {
                    squad = game.otherUsers[j].sprite;
                }
            }
        }
    }

    let sprite = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
    if (!sprite) {
        sprite = CreateEvacuation(jsonData.path_unit.x, jsonData.path_unit.y, jsonData.base_id, jsonData.transport_id);
        sprite.scale.set(0.15);
        sprite.shadow.scale.set(0.15);
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

    if (squad && squadMove) {
        game.add.tween(squad).to({
                x: jsonData.path_unit.x,
                y: jsonData.path_unit.y
            }, jsonData.path_unit.millisecond, Phaser.Easing.Linear.None, true, 0
        );
    }
}

function stopEvacuation(jsonData) {

    let sprite = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
    if (!sprite) {
        sprite = CreateEvacuation(jsonData.path_unit.x, jsonData.path_unit.y, jsonData.base_id, jsonData.transport_id)
    }

    if (game.squad.id === jsonData.other_user.squad_id) {
        EvacuationDown(sprite, game.squad.sprite, true);
    } else {
        for (let j = 0; j < game.otherUsers.length; j++) {
            if (game.otherUsers[j].squad_id === jsonData.other_user.squad_id) {
                EvacuationDown(sprite, game.otherUsers[j].sprite, true);
            }
        }
    }

}

function placeEvacuation(jsonData) {
    let sprite = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
    if (!sprite) {
        sprite = CreateEvacuation(jsonData.path_unit.x, jsonData.path_unit.y, jsonData.base_id, jsonData.transport_id)
    }
    // TODO создать еще 1 мелкую тень которая падает на корпус + тень которая на земле
    // TODO сделать так что бы мелкая тень при пересечение корпуса пропадала (и осталась ток на корпусе)
    // TODO большая тень прорисовывалась только если находиться за ТЕНЬЮ корпуса
    // https://codepen.io/BeFiveINFO/pen/bdJvad

    EvacuationDown(sprite);

    setTimeout(function () {
        if (game.squad.id === jsonData.other_user.squad_id) {
            EvacuationUp(sprite, game.squad.sprite);
        } else {
            for (let j = 0; j < game.otherUsers.length; j++) {
                if (game.otherUsers[j].squad_id === jsonData.other_user.squad_id) {
                    EvacuationUp(sprite, game.otherUsers[j].sprite);
                }
            }
        }
    }, 1200)
}