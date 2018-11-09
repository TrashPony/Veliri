function LaunchLaser(xStart, yStart, angle, targetX, targetY) {
    let fakeBullet = game.artilleryBulletLayer.create(xStart, yStart, "missile_bullet", 0);
    fakeBullet.angle = angle;
    fakeBullet.anchor.setTo(0.5);
    fakeBullet.scale.setTo(0.5);
    fakeBullet.alpha = 0;

    let laserNeon = game.add.emitter(0, 0, 1000);
    laserNeon.makeParticles('laserBall');
    laserNeon.lifespan = 300;
    laserNeon.setXSpeed(0, 0);
    laserNeon.setYSpeed(0, 0);
    laserNeon.minRotation = 0;
    laserNeon.maxRotation = 0;
    laserNeon.gravity = 0;

    let laserTrail = game.add.emitter(0, 0, 1000);
    laserTrail.makeParticles('fire1');
    laserTrail.minParticleScale = 1;
    laserTrail.maxParticleScale = 2;
    laserTrail.lifespan = 3000;
    laserTrail.setXSpeed(-50, 50);
    laserTrail.setYSpeed(-50, 50);
    laserTrail.gravity = 0;
    laserTrail.setAlpha(0.03, 0, 3000, null, true);

    let laser = game.add.emitter(0, 0, 1000);
    laser.makeParticles('laserBall');
    laser.lifespan = 300;
    laser.setXSpeed(0, 0);
    laser.setYSpeed(0, 0);
    laser.minRotation = 0;
    laser.maxRotation = 0;
    laser.gravity = 0;

    fakeBullet.typeBullet = "laser";
    fakeBullet.laser = laser;
    fakeBullet.laserTrail = laserTrail;
    fakeBullet.laserNeon = laserNeon;
    fakeBullet.targetX = targetX;
    fakeBullet.targetY = targetY;

    game.physics.enable(fakeBullet, Phaser.Physics.ARCADE);
    game.physics.arcade.moveToXY(fakeBullet, targetX, targetY, 2000);
}