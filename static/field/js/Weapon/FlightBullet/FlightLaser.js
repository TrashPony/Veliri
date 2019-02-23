function FlightLaser(bullet) {

    if (!bullet.alive) {
        bullet.laserTrail.destroy();
        bullet.fakeBulletEnd.destroy();
        bullet.destroy();
    }

    let dist = game.physics.arcade.distanceToXY(bullet, bullet.targetX, bullet.targetY);
    let distEnd = game.physics.arcade.distanceToXY(bullet.fakeBulletEnd, bullet.targetX, bullet.targetY);

    if (dist < 50 && bullet.alive) {
        bullet.body.angularVelocity = 0;
        bullet.body.velocity.x = 0;
        bullet.body.velocity.y = 0;
        setTimeout(function () {
            bullet.alive = false;
        }, 3000);
        setTimeout(function () {
            bullet.laserOut.destroy();
            bullet.laserIn.destroy();
        }, 200);
    }

    if (distEnd < 50 && bullet.fakeBulletEnd.alive) {
        bullet.fakeBulletEnd.body.angularVelocity = 0;
        bullet.fakeBulletEnd.body.velocity.x = 0;
        bullet.fakeBulletEnd.body.velocity.y = 0;
    }

    bullet.laserOut.clear();
    bullet.laserOut.lineStyle(6, 0x10EDFF, 1);
    bullet.laserOut.moveTo(bullet.x, bullet.y);
    bullet.laserOut.lineTo(bullet.fakeBulletEnd.x, bullet.fakeBulletEnd.y);

    bullet.laserIn.clear();
    bullet.laserIn.lineStyle(2, 0xFFFFFF, 1);
    bullet.laserIn.moveTo(bullet.x, bullet.y);
    bullet.laserIn.lineTo(bullet.fakeBulletEnd.x, bullet.fakeBulletEnd.y);

    if (bullet.alive) {
        bullet.laserTrail.x = bullet.x;
        bullet.laserTrail.y = bullet.y;
        bullet.laserTrail.emitParticle();
        bullet.laserTrail.forEach(function (item) {
            item.tint = 0x0000FF;
        });
    }
}