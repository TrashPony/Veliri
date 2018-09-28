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
    fireMuzzle.animations.play('fireMuzzle_1', 10, true, false);

    let tween = game.add.tween(fireMuzzle).to({alpha: 0}, 5, Phaser.Easing.Linear.None, true, 250);
    tween.onComplete.add(function (fireMuzzle) {
        fireMuzzle.destroy();
    }, fireMuzzle);

    unit.sprite.addChild(fireMuzzle);
}