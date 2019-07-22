function SetAngle(unit, angle, time, ignoreOldTween) {
    if (!unit.rotateTween || ignoreOldTween) {
        SetShadowAngle(unit, time);
        if (angle > 180) {
            angle -= 360
        }

        unit.rotateTween = ShortDirectionRotateTween(unit.sprite, Phaser.Math.degToRad(angle), time);
        unit.rotateTween.onComplete.add(function () {
            unit.rotateTween = null;
        })
    }
}

function SetShadowAngle(unit, time) {

    let shadowTime = function (sprite, newX, newY, rotateTime = 10) {
        game.add.tween(sprite).to({
            'x': newX,
            'y': newY,
        }, rotateTime, Phaser.Easing.Linear.None, true);
    };

    // поворачиваем тени относительно поворота главного спрайта unit.sprite.angle
    // выставляем положение каждые 100мс
    let rotateShadow = setInterval(function () {
        let shadowAngle = 45 - unit.sprite.angle;
        let connectPoints = PositionAttachSprite(shadowAngle, game.shadowXOffset);

        shadowTime(unit.sprite.bodyShadow, connectPoints.x, connectPoints.y);
        shadowTime(unit.sprite.bodyBottomShadow, connectPoints.x, connectPoints.y);

        if (unit.sprite.weapon) {
            let connectWeapons = PositionAttachSprite(shadowAngle, game.shadowXOffset / 2);
            shadowTime(unit.sprite.weaponShadow, connectWeapons.x + unit.sprite.weapon.xAttach, connectWeapons.y + unit.sprite.weapon.yAttach);
        }

        for (let i = 0; i < unit.sprite.equipSprites.length; i++) {
            let slot = unit.sprite.equipSprites[i];
            let connectWeapons = PositionAttachSprite(shadowAngle, game.shadowXOffset / 4);
            shadowTime(slot.shadow, connectWeapons.x + slot.xAttach, connectWeapons.y + slot.yAttach);
            slot.shadow.angle = slot.sprite.angle;
        }
    }, 10);

    // когда кончается общее время данное на поворот time, останавливаем проверку положения тени
    setTimeout(function () {
        clearInterval(rotateShadow);
    }, time)
}