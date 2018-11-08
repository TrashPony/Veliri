function launchRocket(xStart,yStart, angle, target) {

    let targetX = target.sprite.x + 50;
    let targetY = target.sprite.y + 40;

    let missileBulletShadow = game.bulletLayer.create(game.shadowXOffset * 2, game.shadowYOffset * 2, "missile_bullet", 0);
    missileBulletShadow.anchor.set(0.5);
    missileBulletShadow.tint = 0x000000;
    missileBulletShadow.alpha = 0.4;
    missileBulletShadow.angle = 15;

    let missileBullet = game.bulletLayer.create(xStart, yStart, "missile_bullet", 0);
    missileBullet.angle = angle;
    missileBullet.anchor.setTo(0.5, 0.7);
    missileBullet.scale.setTo(0.5);
    missileBullet.alpha = 1;
    missileBullet.targetX = targetX;
    missileBullet.targetY = targetY;

    missileBullet.addChild(missileBulletShadow);
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
    fireTrail.makeParticles( [ 'fire1', 'fire2', 'fire3'] );
    fireTrail.lifespan = 200;
    fireTrail.setScale(0.15, 0, 0.15, 0, 500);
    fireTrail.setXSpeed(0, 0);
    fireTrail.setYSpeed(0, 0);

    missileBullet.smokeTrail = smokeTrail;
    missileBullet.smokeTrailShadow = smokeTrailShadow;
    missileBullet.fireTrail = fireTrail;
    missileBullet.addChild(fireTrail);
    missileBullet.typeBullet = "rocket";

    game.add.tween(fireTrail).to({y: 10}, 700, Phaser.Easing.Linear.None, true, 0);

    game.add.tween(missileBulletShadow).to({x: game.shadowXOffset * 10}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(missileBulletShadow).to({y: game.shadowYOffset * 10}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(missileBulletShadow).to({angle: 0}, 700, Phaser.Easing.Linear.None, true, 0);

    game.add.tween(missileBullet.scale).to({x: 0.55, y: 0.55}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(missileBullet).to({alpha: 1}, 500, Phaser.Easing.Linear.None, true, 0);


    game.physics.enable(missileBullet, Phaser.Physics.ARCADE);

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