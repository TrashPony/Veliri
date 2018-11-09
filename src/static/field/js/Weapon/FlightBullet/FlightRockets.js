function FlightRockets(bullet) {
    let dist = game.physics.arcade.distanceToXY(bullet, bullet.targetX, bullet.targetY);

    if (dist < 120) {
        if (!bullet.slow && bullet.artillery) {
            bullet.slow = true;

            game.add.tween(bullet.shadow).to({x: game.shadowXOffset}, 900, Phaser.Easing.Linear.None, true, 0);
            game.add.tween(bullet.shadow).to({y: game.shadowYOffset}, 900, Phaser.Easing.Linear.None, true, 0);

            game.add.tween(bullet.shadow).to({angle: -10}, 900, Phaser.Easing.Linear.None, true, 0);

            game.add.tween(bullet.scale).to({x: 0.45, y: 0.45}, 900, Phaser.Easing.Linear.None, true, 0);
            game.add.tween(bullet.fireTrail).to({y: 5}, 900, Phaser.Easing.Linear.None, true, 0);

            bullet.animations.add('end', [21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40]);
            bullet.animations.play('end', 30, false, true);

            bullet.shadow.animations.add('end', [21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40]);
            bullet.shadow.animations.play('end', 30, false, true);

            game.physics.arcade.moveToXY(bullet, bullet.targetX, bullet.targetY, 150);
            setTimeout(function () {
                game.physics.arcade.moveToXY(bullet, bullet.targetX, bullet.targetY, 100);
            }, 200);
        } else if (!bullet.artillery && dist < 10) {
            bullet.alive = false;
        }
    }

    bullet.smokeTrail.x = bullet.x;
    bullet.smokeTrail.y = bullet.y;
    bullet.smokeTrail.emitParticle();

    bullet.smokeTrailShadow.x = bullet.shadow.world.x;
    bullet.smokeTrailShadow.y = bullet.shadow.world.y;
    bullet.smokeTrailShadow.emitParticle();

    bullet.fireTrail.emitParticle();

    if (!bullet.alive) {
        Explosion(bullet.x, bullet.y);
        bullet.shadow.destroy();
        bullet.destroy();
    }
}