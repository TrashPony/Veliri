function LaunchLaser(xStart, yStart, angle, targetX, targetY) {
    let fakeBullet = game.artilleryBulletLayer.create(xStart, yStart, "missile_bullet", 0);
    fakeBullet.angle = angle;
    fakeBullet.anchor.setTo(0.5);
    fakeBullet.scale.setTo(0.5);
    fakeBullet.alpha = 0;

    let fakeBulletEnd = game.artilleryBulletLayer.create(xStart, yStart, "missile_bullet", 0);
    fakeBulletEnd.angle = angle;
    fakeBulletEnd.anchor.setTo(0.5);
    fakeBulletEnd.scale.setTo(0.5);
    fakeBulletEnd.alpha = 0;

    let laserTrail = game.add.emitter(0, 0, 1000);
    laserTrail.makeParticles('fire1');
    laserTrail.minParticleScale = 0.2;
    laserTrail.maxParticleScale = 1;
    laserTrail.lifespan = 1000;
    laserTrail.setXSpeed(-50, 50);
    laserTrail.setYSpeed(-50, 50);
    laserTrail.gravity = 0;
    laserTrail.setAlpha(0.03, 0, 3000, null, true);

    let laserOut = game.add.graphics(0, 0);
    laserOut.lineStyle(6, 0x10EDFF, 1);

    let laserIn = game.add.graphics(0, 0);
    laserIn.lineStyle(2, 0xFFFFFF, 1);

    let blurX = game.add.filter('BlurX');
    let blurY = game.add.filter('BlurY');
    blurX.blur = 2;
    blurY.blur = 2;
    laserOut.filters = [blurX, blurY];
    blurX.blur = 1;
    blurY.blur = 1;
    laserIn.filters = [blurX, blurY];

    fakeBullet.typeBullet = "laser";
    fakeBullet.laserOut = laserOut;
    fakeBullet.laserIn = laserIn;
    fakeBullet.laserTrail = laserTrail;
    fakeBullet.targetX = targetX;
    fakeBullet.targetY = targetY;

    fakeBullet.fakeBulletEnd = fakeBulletEnd;

    game.physics.enable(fakeBullet, Phaser.Physics.ARCADE);
    game.physics.enable(fakeBulletEnd, Phaser.Physics.ARCADE);

    game.add.tween(fakeBullet).to({x: targetX, y: targetY}, 100, Phaser.Easing.Linear.None, true, 0);

    setTimeout(function () {
        game.add.tween(fakeBulletEnd).to({x: targetX, y: targetY}, 100, Phaser.Easing.Linear.None, true, 0);
    }, 200);
}