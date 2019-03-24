function Fire(unit, target, targetType) {

    let targetX;
    let targetY;

    if (targetType === "coordinate" || targetType === "inFog") {
        targetX = target.sprite.x + 50;
        targetY = target.sprite.y + 40;
    } else {
        targetX = target.sprite.x;
        targetY = target.sprite.y;
    }

    let weapon;

    for (let i in unit.body.weapons) {
        if (unit.body.weapons.hasOwnProperty(i) && unit.body.weapons[i].weapon) {
            weapon = unit.body.weapons[i].weapon
        }
    }

    if (!weapon) return;

    let rotate = Math.atan2(targetY - unit.sprite.weapon.world.y, targetX - unit.sprite.weapon.world.x);
    let angle = (rotate * 180 / 3.14);

    // todo timeRotate не верно считается, и юнит выбирает не самый оптимальный угол поворота
    // TODO стартовая позиция снаряда с зависимостью от скейла карты
    
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

    return new Promise((resolve) => {
        rotateTower.onComplete.add(function () {
            if (weapon.type === "missile") {
                let connectPoints = PositionAttachSprite(unit.sprite.weapon.angle, unit.sprite.weapon.width / 2);

                launchRocket(unit.sprite.weapon.world.x - connectPoints.x / 2, unit.sprite.weapon.world.y - connectPoints.y / 2, unit.sprite.weapon.angle, targetX, targetY, weapon.artillery, targetType);
                setTimeout(function () {
                    launchRocket(unit.sprite.weapon.world.x + connectPoints.x / 2, unit.sprite.weapon.world.y + connectPoints.y / 2, unit.sprite.weapon.angle, targetX, targetY, weapon.artillery, targetType)
                        .then(function () {
                            resolve();
                        });
                }, 500);
            }

            if (weapon.type === "firearms") {
                if (weapon.artillery) {
                    let connectPointsOne = PositionAttachSprite(unit.sprite.weapon.angle - 5, unit.sprite.weapon.width);
                    let connectPointsTwo = PositionAttachSprite(unit.sprite.weapon.angle + 5, unit.sprite.weapon.width);

                    LaunchArtilleryBallistics(unit.sprite.weapon.world.x + connectPointsOne.x, unit.sprite.weapon.world.y + connectPointsOne.y, angle, targetX, targetY, targetType);
                    setTimeout(function () {
                        LaunchArtilleryBallistics(unit.sprite.weapon.world.x + connectPointsTwo.x, unit.sprite.weapon.world.y + connectPointsTwo.y, angle, targetX - 9, targetY - 7, targetType)
                            .then(function () {
                                resolve();
                            });
                    }, 500);
                } else {
                    let connectPoints = PositionAttachSprite(unit.sprite.weapon.angle, unit.sprite.weapon.width / 1.1);

                    LaunchSmallBallistics(unit.sprite.weapon.world.x + connectPoints.x, unit.sprite.weapon.world.y + connectPoints.y, angle, targetX, targetY, targetType)
                        .then(function () {
                            resolve();
                        });
                }
            }

            if (weapon.type === "laser") {
                if (weapon.name === "big_laser") {
                    let connectPointsOne = PositionAttachSprite(unit.sprite.weapon.angle - 5, unit.sprite.weapon.width / 1.5);
                    let connectPointsTwo = PositionAttachSprite(unit.sprite.weapon.angle + 5, unit.sprite.weapon.width / 1.5);

                    LaunchLaser(unit.sprite.weapon.world.x + connectPointsOne.x, unit.sprite.weapon.world.y + connectPointsOne.y, angle, targetX, targetY, targetType);
                    LaunchLaser(unit.sprite.weapon.world.x + connectPointsTwo.x, unit.sprite.weapon.world.y + connectPointsTwo.y, angle, targetX, targetY, targetType)
                        .then(function () {
                            resolve();
                        })
                }
                if (weapon.name === "small_laser") {
                    let connectPoints = PositionAttachSprite(unit.sprite.weapon.angle, unit.sprite.weapon.width / 1.5);
                    LaunchLaser(unit.sprite.weapon.world.x + connectPoints.x, unit.sprite.weapon.world.y + connectPoints.y, angle, targetX, targetY, targetType)
                        .then(function () {
                            resolve();
                        })
                }
            }
        });
    })
}