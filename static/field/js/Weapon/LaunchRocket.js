function launchRocket(xStart, yStart, angle, targetX, targetY, artillery, targetType) {

    let missileBulletShadow = game.artilleryBulletLayer.create(xStart + game.shadowXOffset * 2, yStart + game.shadowYOffset * 2, "missile_bullet", 0);
    missileBulletShadow.angle = angle;
    missileBulletShadow.anchor.set(0.5);
    missileBulletShadow.scale.setTo(0.25);
    missileBulletShadow.tint = 0x000000;
    missileBulletShadow.alpha = 0.4;

    let missileBullet = game.artilleryBulletLayer.create(xStart, yStart, "missile_bullet", 0);
    missileBullet.angle = angle;
    missileBullet.anchor.setTo(0.5, 0.7);
    missileBullet.scale.setTo(0.25);
    missileBullet.alpha = 1;
    missileBullet.targetX = targetX;
    missileBullet.targetY = targetY;

    missileBullet.shadow = missileBulletShadow;

    let smokeTrail = game.add.emitter(0, 0, 100);
    smokeTrail.makeParticles('smoke_puff');
    smokeTrail.minParticleScale = 0.1;
    smokeTrail.maxParticleScale = 0.2;
    smokeTrail.lifespan = 1400;
    smokeTrail.setXSpeed(0, 0);
    smokeTrail.setYSpeed(0, 0);
    smokeTrail.gravity = -20;
    smokeTrail.setAlpha(0, 0.5, 700, null, true);

    let smokeTrailShadow = game.add.emitter(0, 0, 100);
    smokeTrailShadow.makeParticles('smoke_puff');
    smokeTrailShadow.minParticleScale = 0.1;
    smokeTrailShadow.maxParticleScale = 0.2;
    smokeTrailShadow.lifespan = 1400;
    smokeTrailShadow.setXSpeed(0, 0);
    smokeTrailShadow.setYSpeed(0, 0);
    smokeTrailShadow.gravity = -20;
    smokeTrailShadow.setAlpha(0, 0.1, 700, null, true);

    let fireTrail = game.add.emitter(0, 5, 100);
    fireTrail.makeParticles(['fire1', 'fire2', 'fire3']);
    fireTrail.lifespan = 200;
    fireTrail.setScale(0.15, 0, 0.15, 0, 500);
    fireTrail.setXSpeed(0, 0);
    fireTrail.setYSpeed(0, 0);

    missileBullet.smokeTrail = smokeTrail;
    missileBullet.smokeTrailShadow = smokeTrailShadow;
    missileBullet.fireTrail = fireTrail;
    missileBullet.addChild(fireTrail);

    missileBullet.targetType = targetType;
    missileBullet.typeBullet = "rocket";
    missileBullet.artillery = artillery;

    game.physics.enable(missileBullet, Phaser.Physics.ARCADE);
    game.physics.enable(missileBulletShadow, Phaser.Physics.ARCADE);

    if (artillery) {
        if (targetType === "outFog") {

            missileBulletShadow.angle = angle;
            missileBulletShadow.frame = 20;
            missileBulletShadow.x = xStart + game.shadowXOffset * 10;
            missileBulletShadow.y = yStart + game.shadowYOffset * 10;
            missileBulletShadow.scale.set(0.30);

            missileBullet.scale.set(0.30);
            missileBullet.frame = 20;

        } else {
            game.add.tween(fireTrail).to({y: 10}, 700, Phaser.Easing.Linear.None, true, 0);

            game.add.tween(missileBullet.scale).to({x: 0.30, y: 0.30}, 700, Phaser.Easing.Linear.None, true, 0);
            game.add.tween(missileBulletShadow.scale).to({x: 0.30, y: 0.30}, 700, Phaser.Easing.Linear.None, true, 0);

            missileBullet.shadow.animations.add('launch', [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]);
            missileBullet.shadow.animations.play('launch', 30, false, false);

            missileBullet.animations.add('launch', [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]);
            missileBullet.animations.play('launch', 30, false, false);

            game.physics.arcade.moveToXY(missileBullet, targetX, targetY, 50);

            setTimeout(function () {
                game.physics.arcade.moveToXY(missileBullet, targetX, targetY, 100);
            }, 250);

            setTimeout(function () {
                game.physics.arcade.moveToXY(missileBullet, targetX, targetY, 150);
            }, 500);

            setTimeout(function () {
                game.physics.arcade.moveToXY(missileBullet, targetX, targetY, 200);
            }, 750);
        }
    } else {
        missileBulletShadow.angle = angle;
        missileBullet.frame = 20;
        game.physics.arcade.moveToXY(missileBullet, targetX, targetY, 200);
        game.physics.arcade.moveToXY(missileBulletShadow, targetX + game.shadowXOffset * 2, targetY + game.shadowYOffset * 2, 200);
    }

    if (targetType === "outFog") {
        missileBullet.alpha = 0;
        missileBulletShadow.alpha = 0;
        smokeTrail.alpha = 0;
        smokeTrailShadow.alpha = 0;
        fireTrail.alpha = 0;

        game.add.tween(missileBullet).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(missileBulletShadow).to({alpha: 0.4}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(smokeTrail).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(smokeTrailShadow).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(fireTrail).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);

        game.physics.arcade.moveToXY(missileBullet, targetX, targetY, 200);
    }

    return new Promise((resolve) => {
        missileBullet.events.onDestroy.add(function () {
            resolve();
        }, this);
    })
}