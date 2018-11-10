function LaunchArtilleryBallistics(xStart, yStart, angle, targetX, targetY, targetType) {

    if (targetType !== "outFog") {
        let fireMuzzle = game.weaponEffectsLayer.create(xStart, yStart, 'fireMuzzle_2');
        fireMuzzle.angle = angle - 90;
        fireMuzzle.anchor.setTo(0.5);
        fireMuzzle.animations.add('fireMuzzle_2', [2, 1, 0]);
        fireMuzzle.animations.play('fireMuzzle_2', 10, false, true);
    }

    let bulletShadow = game.artilleryBulletLayer.create(game.shadowXOffset * 2, game.shadowYOffset * 2, "ballistics_artillery_bullet");
    bulletShadow.anchor.set(0.5);
    bulletShadow.tint = 0x000000;
    bulletShadow.alpha = 0.4;

    let bullet = game.artilleryBulletLayer.create(xStart, yStart, "ballistics_artillery_bullet");
    bullet.angle = angle;
    bullet.anchor.setTo(0.5);
    bullet.scale.setTo(0.3);
    bullet.alpha = 1;
    bullet.addChild(bulletShadow);

    let distToTarget = game.physics.arcade.distanceToXY(bullet, targetX, targetY); // время полета 2мс на пиксель

    let xCenter = (xStart + targetX) / 2;
    let yCenter = (yStart + targetY) / 2;

    xCenter = (xCenter + targetX) / 2;
    yCenter = (yCenter + targetY) / 2;


    if (targetType !== "outFog" && targetType !== "inFog") {

        let startFire = game.add.tween(bullet).to({
            x: xCenter,
            y: yCenter
        }, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);
        game.add.tween(bullet.scale).to({x: 0.5, y: 0.5}, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);
        game.add.tween(bulletShadow).to({
            x: game.shadowXOffset * 15,
            y: game.shadowYOffset * 15
        }, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);

        return new Promise((resolve) => {
            startFire.onComplete.add(function () {
                let endFire = game.add.tween(bullet).to({
                    x: targetX,
                    y: targetY
                }, distToTarget * 2, Phaser.Easing.Sinusoidal.In, true, 0);
                game.add.tween(bullet.scale).to({
                    x: 0.3,
                    y: 0.3
                }, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);
                game.add.tween(bulletShadow).to({
                    x: game.shadowXOffset * 2,
                    y: game.shadowYOffset * 2
                }, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);

                endFire.onComplete.add(function () {
                    Explosion(bullet.x, bullet.y);
                    bullet.destroy();
                    resolve();
                });
            });
        });
    } else if (targetType === "outFog") {
        bullet.alpha = 0;
        bulletShadow.alpha = 0;

        game.add.tween(bullet).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(bulletShadow).to({alpha: 0.4}, 700, Phaser.Easing.Linear.None, true, 0);

        bullet.scale.set(0.5);
        bulletShadow.x = game.shadowXOffset * 15;
        bulletShadow.y = game.shadowYOffset * 15;

        let endFire = game.add.tween(bullet).to({
            x: targetX,
            y: targetY
        }, distToTarget * 2, Phaser.Easing.Sinusoidal.In, true, 0);

        game.add.tween(bullet.scale).to({
            x: 0.3,
            y: 0.3
        }, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);
        game.add.tween(bulletShadow).to({
            x: game.shadowXOffset * 2,
            y: game.shadowYOffset * 2
        }, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);

        endFire.onComplete.add(function () {
            Explosion(bullet.x, bullet.y);
            bullet.destroy();
        });

        return new Promise((resolve) => {
            endFire.onComplete.add(function () {
                resolve();
            });
        });
    } else if (targetType === "inFog") {
        let startFire = game.add.tween(bullet).to({
            x: targetX,
            y: targetY
        }, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);
        game.add.tween(bullet.scale).to({x: 0.5, y: 0.5}, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);
        game.add.tween(bulletShadow).to({
            x: game.shadowXOffset * 15,
            y: game.shadowYOffset * 15
        }, distToTarget * 2, Phaser.Easing.Sinusoidal.Out, true, 0);

        return new Promise((resolve) => {
            startFire.onComplete.add(function () {
                bullet.destroy();
                resolve();
            })
        });
    }
}