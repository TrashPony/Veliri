function ReadResponse(jsonData) {
    if (jsonData.event === "InitGame") {
        LoadGame(jsonData);
    }

    if (jsonData.event === "Error") {
        alert(jsonData.error);
    }

    if (jsonData.event === "PreviewPath") {
        for (let i = 0; i < jsonData.path.length; i++) {
            let label = game.SelectRangeLayer.create(jsonData.path[i].x, jsonData.path[i].y, 'pathCell');
            label.anchor.setTo(0.5);
            label.scale.set(0.5);

            let tween = game.add.tween(label).to({
                alpha: 1
            }, 100 * (i+1), Phaser.Easing.Linear.None, true);

            tween.onComplete.add(function () {
                label.destroy();
            })
        }
    }

    if (jsonData.event === "MoveTo") {
        game.add.tween(game.squad.sprite).to(
            {x: jsonData.path_unit.x, y: jsonData.path_unit.y},
            jsonData.path_unit.millisecond,
            Phaser.Easing.Linear.None, true, 0
        );
        SetMSAngle(game.squad, jsonData.path_unit.rotate + 90, jsonData.path_unit.millisecond)
    }
}

function SetMSAngle(unit, angle, time) {
    if (angle > 180) {
        angle -= 360
    }

    ShortDirectionRotateTween(unit.sprite.unitBody, Phaser.Math.degToRad(angle), time);
    ShortDirectionRotateTween(unit.sprite.bodyShadow, Phaser.Math.degToRad(angle), time);
    if (unit.sprite.weapon) {
        ShortDirectionRotateTween(unit.sprite.weaponShadow, Phaser.Math.degToRad(angle), time);
        ShortDirectionRotateTween(unit.sprite.weapon, Phaser.Math.degToRad(angle), time);
    }
}

function ShortDirectionRotateTween(sprite, desiredRotation, time) {
    // эта функция ищет оптимальный угол для поворота
    let shortestAngle = getShortestAngle(Phaser.Math.radToDeg(desiredRotation), sprite.angle);
    let newAngle = sprite.angle + shortestAngle;
    return game.add.tween(sprite).to({
        'angle': newAngle
    }, time, Phaser.Easing.Linear.None, true);
}

function getShortestAngle(angle1, angle2) {
    // библиотечная функция не подходит
    let difference = angle2 - angle1;
    let times = Math.floor((difference - (-180)) / 360);
    return (difference - (times * 360)) * -1;
}