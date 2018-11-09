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

    // todo перемещать спайт цели на верх если снаряд не артилерийский
    // game.world.bringToTop(target.sprite); // позволяет поднять группу спрайтов поверх всего мира
    // var oldParent = sprite.parent;oldParent.remove(sprite);newParent.add(sprite); переместить спрайт из 1 в другую группу

    game.add.tween(unit.sprite.weaponShadow).to({angle: angle}, timeRotate, Phaser.Easing.Linear.None, true, 0);
    let rotateTower = game.add.tween(unit.sprite.weapon).to({angle: angle}, timeRotate, Phaser.Easing.Linear.None, true, 0);

    rotateTower.onComplete.add(function () {
        if (weapon.type === "missile") {
            let connectPoints = PositionAttachSprite(unit.sprite.weapon.angle, unit.sprite.weapon.width / 2);

            launchRocket(unit.sprite.weapon.world.x - connectPoints.x / 2, unit.sprite.weapon.world.y - connectPoints.y / 2, unit.sprite.weapon.angle, target, weapon.artillery);
            setTimeout(function () {
                launchRocket(unit.sprite.weapon.world.x + connectPoints.x / 2, unit.sprite.weapon.world.y + connectPoints.y / 2, unit.sprite.weapon.angle, target, weapon.artillery);
            }, 500);
        }

        if (weapon.type === "firearms") {
            if (weapon.artillery) {
                let connectPointsOne = PositionAttachSprite(unit.sprite.weapon.angle - 85, unit.sprite.weapon.width);
                let connectPointsTwo = PositionAttachSprite(unit.sprite.weapon.angle - 95, unit.sprite.weapon.width);
                LaunchArtilleryBallistics(unit.sprite.weapon.world.x + connectPointsOne.x, unit.sprite.weapon.world.y + connectPointsOne.y, angle, target.sprite.x + 50, target.sprite.y + 40);
                setTimeout(function () {
                    LaunchArtilleryBallistics(unit.sprite.weapon.world.x + connectPointsTwo.x, unit.sprite.weapon.world.y + connectPointsTwo.y, angle, target.sprite.x + 50 / 1.2, target.sprite.y + 40 / 1.2);
                }, 500);
            } else {
                let connectPoints = PositionAttachSprite(unit.sprite.weapon.angle - 90, unit.sprite.weapon.width / 1.1);
                LaunchSmallBallistics(unit.sprite.weapon.world.x + connectPoints.x, unit.sprite.weapon.world.y + connectPoints.y, angle, target)
            }
        }

        if (weapon.type === "laser") {
            if (weapon.name === "small_laser") {
                let connectPoints = PositionAttachSprite(unit.sprite.weapon.angle - 90, unit.sprite.weapon.width / 1.5);
                LaunchLaser(unit.sprite.weapon.world.x + connectPoints.x, unit.sprite.weapon.world.y + connectPoints.y, angle, target.sprite.x + 50, target.sprite.y + 40)
            }
            if (weapon.name === "big_laser") {
                let connectPointsOne = PositionAttachSprite(unit.sprite.weapon.angle - 85, unit.sprite.weapon.width / 1.5);
                let connectPointsTwo = PositionAttachSprite(unit.sprite.weapon.angle - 95, unit.sprite.weapon.width / 1.5);

                LaunchLaser(unit.sprite.weapon.world.x + connectPointsOne.x, unit.sprite.weapon.world.y + connectPointsOne.y, angle, target.sprite.x + 50, target.sprite.y + 40);
                LaunchLaser(unit.sprite.weapon.world.x + connectPointsTwo.x, unit.sprite.weapon.world.y + connectPointsTwo.y, angle, target.sprite.x + 50 / 1.2, target.sprite.y + 40 / 1.2);
            }
        }
    });
}