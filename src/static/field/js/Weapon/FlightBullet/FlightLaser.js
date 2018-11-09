function FlightLaser(bullet) {

    if (!bullet.alive) {
        bullet.destroy();
        setTimeout(function () {
            bullet.laserTrail.destroy();
        }, 3000);
        setTimeout(function () {
            bullet.laser.destroy();
            bullet.laserNeon.destroy();
        }, 100);
    } else {

        let dist = game.physics.arcade.distanceToXY(bullet, bullet.targetX, bullet.targetY);

        if (dist < 60) {
            bullet.alive = false;
        }

        bullet.laser.x = bullet.x;
        bullet.laser.y = bullet.y;

        bullet.laserTrail.x = bullet.x;
        bullet.laserTrail.y = bullet.y;

        bullet.laserNeon.x = bullet.x;
        bullet.laserNeon.y = bullet.y;

        bullet.laser.emitParticle();
        bullet.laserTrail.emitParticle();
        bullet.laserNeon.emitParticle();

        bullet.laser.forEach(function (item) {
            item.scale.setTo(0.05, 0.22);
            if (bullet.angle < 0) {
                item.angle = bullet.angle - 1;
            } else {
                item.angle = bullet.angle + 1;
            }
            item.alpha = 1;
        });

        bullet.laserTrail.forEach(function (item) {
            item.tint = 0x0000FF;
        });

        bullet.laserNeon.forEach(function (item) {
            item.angle = bullet.angle;
            item.scale.setTo(0.4, 0.3);
            item.alpha = 0.05;
        });
    }
}