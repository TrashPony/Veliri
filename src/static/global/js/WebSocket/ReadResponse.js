function ReadResponse(jsonData) {
    if (jsonData.event === "InitGame") {
        LoadGame(jsonData);
    }

    if (jsonData.event === "Error") {
        alert(jsonData.error);
    }

    if (jsonData.event === "MoveTo") {
        game.add.tween(game.squad.sprite).to(
            {x: jsonData.path.x, y: jsonData.path.y},
            jsonData.path.millisecond,
            Phaser.Easing.Linear.None, true, 0
        );
        SetMSAngle(game.squad, jsonData.path.rotate + 90, jsonData.path.millisecond)
    }
}

function SetMSAngle(unit, angle, time) {
    if (angle > 180) {
        angle -= 360
    }

    a(unit.sprite.unitBody, Phaser.Math.degToRad(angle), time);
    a(unit.sprite.bodyShadow, Phaser.Math.degToRad(angle), time);
    if (unit.sprite.weapon) {
        a(unit.sprite.weaponShadow, Phaser.Math.degToRad(angle), time);
        a(unit.sprite.weapon, Phaser.Math.degToRad(angle), time);
    }
}

function a(sprite, desiredRotation, time) {
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