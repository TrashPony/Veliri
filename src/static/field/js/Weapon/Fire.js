function Fire(unit, target) {

    let weapon;

    for (let i in unit.body.weapons) {
        if (unit.body.weapons.hasOwnProperty(i) && unit.body.weapons[i].weapon) {
            weapon = unit.body.weapons[i].weapon
        }
    }

    let rotate = Math.atan2(target.sprite.y + 40 - unit.sprite.weapon.world.y, target.sprite.x + 50 - unit.sprite.weapon.world.x);
    let angle = (rotate * 180 / 3.14) + 90;

    let timeRotate;
    if (angle >= unit.sprite.weapon.angle) {
        timeRotate = (angle - unit.sprite.weapon.angle) * 15;
    } else {
        timeRotate = (unit.sprite.weapon.angle - angle) * 15;
    }

    game.add.tween(unit.sprite.weaponShadow).to({angle: angle}, timeRotate, Phaser.Easing.Linear.None, true, 0);
    let rotateTower = game.add.tween(unit.sprite.weapon).to({angle: angle}, timeRotate, Phaser.Easing.Linear.None, true, 0);

    rotateTower.onComplete.add(function () {

        let connectPoints;

        if (weapon.type === "missile") {
            connectPoints = PositionAttachSprite(unit.sprite.weapon.angle, unit.sprite.weapon.width / 2);

            launchRocket(unit.sprite.weapon.world.x - connectPoints.x / 2, unit.sprite.weapon.world.y - connectPoints.y / 2, unit.sprite.weapon.angle, target, weapon.artillery);
            setTimeout(function () {
                launchRocket(unit.sprite.weapon.world.x + connectPoints.x / 2, unit.sprite.weapon.world.y + connectPoints.y / 2, unit.sprite.weapon.angle, target, weapon.artillery);
            }, 500);
        }

        if (weapon.type === "firearms") {
            connectPoints = PositionAttachSprite(unit.sprite.weapon.angle - 90, unit.sprite.weapon.width / 1.5);

            let fireMuzzle = game.weaponEffectsLayer.create(unit.sprite.weapon.world.x + connectPoints.x, unit.sprite.weapon.world.y + connectPoints.y, 'fireMuzzle_1');

            fireMuzzle.angle = unit.sprite.weapon.angle - 90;
            fireMuzzle.anchor.setTo(0, 0.45);
            fireMuzzle.animations.add('fireMuzzle_1', [2,1,0]);
            fireMuzzle.animations.play('fireMuzzle_1', 10, false, true);
            // todo
        }

        if (weapon.type === "laser") {
            // todo
        }
    });
}