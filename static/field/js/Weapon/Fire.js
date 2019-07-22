function Fire(unit, target, targetType) {

    let targetX;
    let targetY;

    if (targetType === "coordinate" || targetType === "inFog") {
        let xy = GetXYCenterHex(target.q, target.r);
        targetX = xy.x;
        targetY = xy.y;
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

    let rotate = Math.atan2(targetY - (unit.sprite.weapon.world.y / game.camera.scale.y), targetX - (unit.sprite.weapon.world.x / game.camera.scale.x));
    let angle = (rotate * 180 / 3.14);
    let ammoAngle = angle + 90;

    // учитываем положение корпуса на карте
    if (unit.sprite.angle > 0) {
        angle -= unit.sprite.angle;
    } else {
        angle += -1 * unit.sprite.angle;
    }

    let timeRotate;
    if (angle >= unit.sprite.weapon.angle) {
        timeRotate = (angle - unit.sprite.weapon.angle) * 15;
    } else {
        timeRotate = (unit.sprite.weapon.angle - angle) * 15;
    }

    // todo перемещать спайт цели на верх если снаряд не артилерийский
    // game.world.bringToTop(target.sprite); // позволяет поднять группу спрайтов поверх всего мира
    // var oldParent = sprite.parent;oldParent.remove(sprite);newParent.add(sprite); переместить спрайт из 1 в другую группу

    ShortDirectionRotateTween(unit.sprite.weaponShadow, Phaser.Math.degToRad(angle), timeRotate);
    let rotateTower = ShortDirectionRotateTween(unit.sprite.weapon, Phaser.Math.degToRad(angle), timeRotate);

    return new Promise((resolve) => {
        rotateTower.onComplete.add(function () {
            if (weapon.type === "missile") {
                let connectPoints = PositionAttachSprite(ammoAngle, unit.sprite.weapon.width / 2);

                launchRocket(
                    (unit.sprite.weapon.world.x / game.camera.scale.x) - (connectPoints.x / 2),
                    (unit.sprite.weapon.world.y / game.camera.scale.y) - (connectPoints.y / 2),
                    ammoAngle, targetX, targetY, weapon.artillery, targetType);
                setTimeout(function () {
                    launchRocket(
                        (unit.sprite.weapon.world.x / game.camera.scale.x) + connectPoints.x / 2,
                        (unit.sprite.weapon.world.y / game.camera.scale.y) + connectPoints.y / 2,
                        ammoAngle, targetX, targetY, weapon.artillery, targetType)
                        .then(function () {
                            resolve();
                        });
                }, 500);
            }

            if (weapon.type === "firearms") {
                if (weapon.artillery) {
                    let connectPointsOne = PositionAttachSprite(ammoAngle - 95, unit.sprite.weapon.width);
                    let connectPointsTwo = PositionAttachSprite(ammoAngle - 85, unit.sprite.weapon.width);

                    LaunchArtilleryBallistics(
                        (unit.sprite.weapon.world.x / game.camera.scale.x) + connectPointsOne.x,
                        (unit.sprite.weapon.world.y / game.camera.scale.y) + connectPointsOne.y,
                        ammoAngle, targetX, targetY, targetType);

                    setTimeout(function () {
                        LaunchArtilleryBallistics(
                            (unit.sprite.weapon.world.x / game.camera.scale.x) + connectPointsTwo.x,
                            (unit.sprite.weapon.world.y / game.camera.scale.y) + connectPointsTwo.y,
                            ammoAngle, targetX - 9, targetY - 7, targetType)
                            .then(function () {
                                resolve();
                            });
                    }, 500);
                } else {
                    let connectPoints = PositionAttachSprite(ammoAngle - 90, unit.sprite.weapon.width / 1.1);

                    LaunchSmallBallistics(
                        (unit.sprite.weapon.world.x / game.camera.scale.x) + connectPoints.x,
                        (unit.sprite.weapon.world.y / game.camera.scale.y) + connectPoints.y,
                        ammoAngle, targetX, targetY, targetType)
                        .then(function () {
                            resolve();
                        });
                }
            }

            if (weapon.type === "laser") {
                if (weapon.name === "big_laser") {
                    let connectPointsOne = PositionAttachSprite(ammoAngle - 95, unit.sprite.weapon.width / 1.5);
                    let connectPointsTwo = PositionAttachSprite(ammoAngle - 85, unit.sprite.weapon.width / 1.5);

                    LaunchLaser(
                        (unit.sprite.weapon.world.x / game.camera.scale.x) + connectPointsOne.x,
                        (unit.sprite.weapon.world.y / game.camera.scale.y) + connectPointsOne.y,
                        ammoAngle, targetX, targetY, targetType);
                    LaunchLaser(
                        (unit.sprite.weapon.world.x / game.camera.scale.x) + connectPointsTwo.x,
                        (unit.sprite.weapon.world.y / game.camera.scale.y) + connectPointsTwo.y,
                        ammoAngle, targetX, targetY, targetType)
                        .then(function () {
                            resolve();
                        })
                }
                if (weapon.name === "small_laser") {
                    let connectPoints = PositionAttachSprite(ammoAngle - 90, (unit.sprite.weapon.width / 1.5));
                    LaunchLaser(
                        (unit.sprite.weapon.world.x / game.camera.scale.x) + connectPoints.x,
                        (unit.sprite.weapon.world.y / game.camera.scale.y) + connectPoints.y,
                        angle, targetX, targetY, targetType)
                        .then(function () {
                            resolve();
                        })
                }
            }
        });
    })
}