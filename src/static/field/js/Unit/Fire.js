function Fire(unit) {
    let connectPoints = PositionAttachSprite(unit.rotate, unit.sprite.unitBody.width / 2);

    let fireMuzzle = game.make.sprite(connectPoints.x, connectPoints.y, 'fireMuzzle_1');
    if (unit.rotate > 180) {
        fireMuzzle.angle = unit.rotate - 360;
    } else {
        fireMuzzle.angle = unit.rotate;
    }

    fireMuzzle.anchor.setTo(0, 0.5);

    fireMuzzle.animations.add('fireMuzzle_1', [2,1,0]);
    fireMuzzle.animations.play('fireMuzzle_1', 10, false, true);

    unit.sprite.addChild(fireMuzzle);
}

function Explosion(coordinate) {
    let explosion = game.effectsLayer.create(coordinate.sprite.x, coordinate.sprite.y, 'explosion_1');

    explosion.animations.add('explosion_1');
    explosion.animations.play('explosion_1', 10, false, true);
}