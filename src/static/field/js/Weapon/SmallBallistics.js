function LaunchSmallBallistics(xStart, yStart, angle, targetX, targetY, targetType) {
    if (targetType !== "outFog") {
        let fireMuzzle = game.weaponEffectsLayer.create(xStart, yStart, 'fireMuzzle_1');
        fireMuzzle.angle = angle - 90;
        fireMuzzle.anchor.setTo(0.5);
        fireMuzzle.animations.add('fireMuzzle_1', [2, 1, 0]);
        fireMuzzle.animations.play('fireMuzzle_1', 10, false, true);
    }

    let bulletShadow = game.artilleryBulletLayer.create(game.shadowXOffset * 2, game.shadowYOffset * 2, "ballistics_small_bullet");
    bulletShadow.anchor.set(0.5);
    bulletShadow.tint = 0x000000;
    bulletShadow.alpha = 0.4;

    let bullet = game.artilleryBulletLayer.create(xStart, yStart, "ballistics_small_bullet");
    bullet.angle = angle;
    bullet.anchor.setTo(0.5);
    bullet.scale.setTo(0.3);
    bullet.alpha = 1;
    bullet.addChild(bulletShadow);

    let distToTarget = game.physics.arcade.distanceToXY(bullet, targetX, targetY); // время полета 2мс на пиксель

    if (targetType !== "outFog" && targetType !== "inFog") {
        let fire = game.add.tween(bullet).to({
            x: targetX,
            y: targetY
        }, distToTarget * 2, Phaser.Easing.Linear.None, true, 0);

        return new Promise((resolve) => {
            fire.onComplete.add(function () {
                Explosion(bullet.x, bullet.y);
                bullet.destroy();
                resolve();
            });
        });
    } else if (targetType === "outFog") {

        bullet.alpha = 0;
        bulletShadow.alpha = 0;

        game.add.tween(bullet).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(bulletShadow).to({alpha: 0.4}, 700, Phaser.Easing.Linear.None, true, 0);

        let fire = game.add.tween(bullet).to({
            x: targetX,
            y: targetY
        }, distToTarget * 2, Phaser.Easing.Linear.None, true, 0);

        return new Promise((resolve) => {
            fire.onComplete.add(function () {
                Explosion(bullet.x, bullet.y);
                bullet.destroy();
                resolve();
            });
        });
    } else if (targetType === "inFog") {
        let fire = game.add.tween(bullet).to({
            x: targetX,
            y: targetY
        }, distToTarget * 2, Phaser.Easing.Linear.None, true, 0);

        return new Promise((resolve) => {
            fire.onComplete.add(function () {
                bullet.destroy();
                resolve();
            });
        });
    }
}