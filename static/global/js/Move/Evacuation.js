function Evacuation() {
    global.send(JSON.stringify({
        event: "evacuation"
    }));
}

function CreateEvacuation(x, y, baseId, transportId) {
    let sprite;

    if (game.bases[baseId].transports[transportId].fraction === 'Replics') sprite = 'evacuation_replics';
    if (game.bases[baseId].transports[transportId].fraction === 'Explores') sprite = 'evacuation_explores';
    if (game.bases[baseId].transports[transportId].fraction === 'Reverses') sprite = 'evacuation_reverses';

    let shadow = game.flyObjectsLayer.create(x + game.shadowXOffset, y + game.shadowYOffset, sprite);
    shadow.anchor.setTo(0.5);
    shadow.scale.set(0.1);
    shadow.alpha = 0;
    shadow.tint = 0x000000;

    let evacuation = game.flyObjectsLayer.create(x, y, sprite);
    evacuation.anchor.setTo(0.5);
    evacuation.scale.set(0.1);
    evacuation.alpha = 0;
    evacuation.shadow = shadow;

    game.bases[baseId].transports[transportId].sprite = evacuation;

    if (game.bases[baseId].transports[transportId].fraction === 'Reverses') {
        game.add.tween(evacuation.shadow).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);
        game.add.tween(evacuation).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);
    }

    game.add.tween(evacuation.shadow).to({alpha: 0.1}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(evacuation).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);

    return evacuation
}

function startMoveEvacuation(jsonData) {

    let x = jsonData.path_unit.x;
    let y = jsonData.path_unit.y;


    let unit = game.units[jsonData.short_unit.id];

    if (unit.owner_id === game.user_id) {
        FocusUnit(unit.id);
    }

    let tween = game.add.tween(unit).to({alpha: 0.6}, 200, "Linear", true, 0, -1);
    tween.yoyo(true, 1000);

    setTimeout(function () {
        CreateMiniMap(game.map);

        let evacuation = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
        if (!evacuation) {
            evacuation = CreateEvacuation(x, y, jsonData.base_id, jsonData.transport_id);
        }

        if (unit.owner_id === game.user_id) {
            game.camera.follow(evacuation);
        }

        setTimeout(function () {
            EvacuationUp(evacuation)
        }, 700);
    }, 500);
}

function EvacuationUp(sprite, unit) {
    game.add.tween(sprite.shadow).to({
        x: sprite.x + game.shadowXOffset * 10,
        y: sprite.y + game.shadowYOffset * 10
    }, 1000, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(sprite.shadow.scale).to({x: 0.15, y: 0.15}, 1000, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(sprite.scale).to({x: 0.15, y: 0.15}, 1000, Phaser.Easing.Linear.None, true, 0);

    if (unit) {
        game.unitLayer.remove(unit);
        game.flyObjectsLayer.add(unit);
        game.flyObjectsLayer.swap(sprite, unit);

        game.add.tween(unit.bodyShadow).to({
            x: game.shadowXOffset * 10,
            y: game.shadowYOffset * 10
        }, 1000, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(unit.scale).to({x: 1.1, y: 1.1}, 1000, Phaser.Easing.Linear.None, true, 0);
        unit.x = sprite.x;
        unit.y = sprite.y;
    }
}

function EvacuationDown(sprite, unit, destroy) {
    let tween = game.add.tween(sprite.shadow).to({
        x: sprite.x + game.shadowXOffset,
        y: sprite.y + game.shadowYOffset
    }, 1000, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(sprite.shadow.scale).to({x: 0.1, y: 0.1}, 1000, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(sprite.scale).to({x: 0.1, y: 0.1}, 1000, Phaser.Easing.Linear.None, true, 0);

    if (unit) {
        game.add.tween(unit.bodyShadow).to({
            x: game.shadowXOffset,
            y: game.shadowYOffset
        }, 1000, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(unit.scale).to({x: 1, y: 1}, 1000, Phaser.Easing.Linear.None, true, 0);
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

    let path = jsonData.path_unit;
    let unit;
    if (jsonData.short_unit) unit = game.units[jsonData.short_unit.id];

    let sprite = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
    if (!sprite) {
        sprite = CreateEvacuation(path.x, path.y, jsonData.base_id, jsonData.transport_id);
        sprite.scale.set(0.15);
        sprite.shadow.scale.set(0.15);
    }

    game.add.tween(sprite).to({
            x: path.x,
            y: path.y
        }, path.millisecond, Phaser.Easing.Linear.None, true, 0
    );
    ShortDirectionRotateTween(sprite, Phaser.Math.degToRad(path.rotate), path.millisecond);

    game.add.tween(sprite.shadow).to({
            x: path.x + game.shadowXOffset * 10,
            y: path.y + game.shadowYOffset * 10
        }, path.millisecond, Phaser.Easing.Linear.None, true, 0
    );
    ShortDirectionRotateTween(sprite.shadow, Phaser.Math.degToRad(path.rotate), path.millisecond);

    if (unit && squadMove) {
        game.add.tween(unit.sprite).to({
                x: path.x,
                y: path.y
            }, path.millisecond, Phaser.Easing.Linear.None, true, 0
        );
        SetBodyAngle(unit, path.rotate, path.millisecond, true);
    }
}

function stopEvacuation(jsonData) {

    let sprite = game.bases[jsonData.base_id].transports[jsonData.transport_id].sprite;
    if (!sprite) {
        sprite = CreateEvacuation(jsonData.path_unit.x, jsonData.path_unit.y, jsonData.base_id, jsonData.transport_id)
    }

    let unit = game.units[jsonData.short_unit.id];
    EvacuationDown(sprite, unit.sprite, true);
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
        let unit = game.units[jsonData.short_unit.id];
        SetBodyAngle(unit, sprite.angle, 0, true);
        EvacuationUp(sprite, unit.sprite);
    }, 1200)
}