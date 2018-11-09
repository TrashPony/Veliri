function LaunchSmallBallistics(xStart, yStart, angle, target) {
    let fireMuzzle = game.weaponEffectsLayer.create(xStart, yStart, 'fireMuzzle_1');
    fireMuzzle.angle = angle - 90;
    fireMuzzle.anchor.setTo(0.5);
    fireMuzzle.animations.add('fireMuzzle_1', [2, 1, 0]);
    fireMuzzle.animations.play('fireMuzzle_1', 10, false, true);

    let bulletShadow = game.artilleryBulletLayer.create(game.shadowXOffset * 2, game.shadowYOffset * 2, "ballistics_small_bullet");
    bulletShadow.anchor.set(0.5);
    bulletShadow.tint = 0x000000;
    bulletShadow.alpha = 0.4;

    let bullet = game.artilleryBulletLayer.create(xStart, yStart, "ballistics_small_bullet");
    bullet.angle = angle;
    bullet.anchor.setTo(0.5);
    bullet.scale.setTo(0.3);
    bullet.alpha = 1;

    let distToTarget = game.physics.arcade.distanceToXY(bullet, target.sprite.x + 50, target.sprite.y + 40); // время полета 2мс на пиксель

    let fire = game.add.tween(bullet).to({x: target.sprite.x + 50, y: target.sprite.y + 40}, distToTarget*2, Phaser.Easing.Linear.None, true, 0);
    fire.onComplete.add(function(){
        Explosion(bullet.x, bullet.y);
        bullet.destroy();
    });
    bullet.addChild(bulletShadow);
}