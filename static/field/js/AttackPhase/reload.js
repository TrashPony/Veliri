// TODO очередная порция говнокода, но работает
function reloadUnit(unit, unitInTargetCoordinate, resultAction) {
    // игрок видит мс которые перезаряжает и кого перезаряжают
    if (unit && unitInTargetCoordinate) { // +
        let reloadDroid = LaunchReloadDrone({x: unit.sprite.x, y: unit.sprite.y});
        return flyReloadDroid(reloadDroid, {
            x: unitInTargetCoordinate.sprite.x,
            y: unitInTargetCoordinate.sprite.y
        }, {x: unit.sprite.x, y: unit.sprite.y}, unitInTargetCoordinate, unit);
    }

    // игрок видит мс которые перезаряжает но не видет кого
    // end_watch_attack последняя координата видимости полета снаряда
    if (unit && !unitInTargetCoordinate) {
        let reloadDroid = LaunchReloadDrone({x: unit.sprite.x, y: unit.sprite.y});
        return flyReloadDroid(
            reloadDroid,
            GetXYCenterHex(resultAction.end_watch_attack.q, resultAction.end_watch_attack.r),
            {
                x: unit.sprite.x,
                y: unit.sprite.y
            },
            unitInTargetCoordinate, unit
        );
    }

    // не видит кто заправлет, видит кого
    if (!unit && unitInTargetCoordinate) {
        let reloadDroid = LaunchReloadDrone(GetXYCenterHex(resultAction.start_watch_attack.q, resultAction.start_watch_attack.r), true);

        //start_watch_attack
        return flyReloadDroid(
            reloadDroid,
            {
                x: unitInTargetCoordinate.sprite.x,
                y: unitInTargetCoordinate.sprite.y
            },
            GetXYCenterHex(resultAction.start_watch_attack.q, resultAction.start_watch_attack.r),
            unitInTargetCoordinate, unit
        );
    }
}

function flyReloadDroid(droid, toPoint, returnPoint, toUnit, ammoUnit) {

    let dist = game.physics.arcade.distanceToXY(droid, toPoint.x, toPoint.y);

    return new Promise(function (resolve) {
        // движемся к точке куда надо
        droidToPoint(toPoint.x, toPoint.y, droid, dist * 25).then(function () {

            if (toUnit && ammoUnit) {
                // долетаем до юнита проигрываем анимации возвращаемся обратно
                downDroid(toPoint.x, toPoint.y, droid, false).then(function () {
                    toUnit.sprite.addChild(ReloadAnimation(false, true));
                    upDroid(toPoint.x, toPoint.y, droid).then(function () {
                        droidToPoint(returnPoint.x, returnPoint.y, droid, dist * 25).then(function () {
                            downDroid(returnPoint.x, returnPoint.y, droid, true).then(function () {
                                resolve()
                            })
                        });
                    });
                });

            } else if (ammoUnit) {
                // долетаем до тумана войны, пропадаем, появляемся летим обратно
                hideDroid(droid, false).then(function () {
                    unhideDroid(droid).then(function () {
                        droidToPoint(returnPoint.x, returnPoint.y, droid, dist * 25).then(function () {
                            downDroid(returnPoint.x, returnPoint.y, droid, true).then(function () {
                                resolve()
                            })
                        });
                    });
                });
            } else if (toUnit) {
                // заправляем юнита , возвращаемся за туман войны
                downDroid(toPoint.x, toPoint.y, droid, false).then(function () {
                    toUnit.sprite.addChild(ReloadAnimation(false, true));
                    upDroid(toPoint.x, toPoint.y, droid).then(function () {
                        droidToPoint(returnPoint.x, returnPoint.y, droid, dist * 25).then(function () {
                            hideDroid(droid, true).then(function () {
                                resolve()
                            })
                        });
                    });
                });
            }
        });
    })
}

function upDroid(x, y, droid) {
    return new Promise(function (resolve) {
        setTimeout(function () {
            resolve();
        }, 700);

        game.add.tween(droid.shadow.scale).to({x: 0.1, y: 0.1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droid.scale).to({x: 0.1, y: 0.1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droid.shadow).to({
            x: x + game.shadowXOffset * 5,
            y: y + game.shadowYOffset * 5
        }, 700, Phaser.Easing.Linear.None, true, 0);
    });
}

function downDroid(x, y, droid, destroy) {
    return new Promise(function (resolve) {
        if (destroy) {
            hideDroid(droid, destroy).then(function () {
                resolve();
            })
        } else {
            setTimeout(function () {
                resolve();
            }, 700)
        }

        game.add.tween(droid.shadow.scale).to({x: 0.05, y: 0.05}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droid.scale).to({x: 0.05, y: 0.05}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droid.shadow).to({
            x: x,
            y: y
        }, 700, Phaser.Easing.Linear.None, true, 0);
    });
}

function hideDroid(droid, destroy) {
    return new Promise(function (resolve) {
        let tween = game.add.tween(droid.shadow).to({alpha: 0}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droid).to({alpha: 0}, 700, Phaser.Easing.Linear.None, true, 0);
        tween.onComplete.add(function () {
            if (destroy) {
                droid.shadow.destroy();
                droid.destroy();
            }
            resolve();
        })
    });
}

function unhideDroid(droid) {
    return new Promise(function (resolve) {
        let tween = game.add.tween(droid.shadow).to({alpha: 0.3}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(droid).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);
        tween.onComplete.add(function () {
            resolve();
        })
    });
}

function droidToPoint(x, y, droid, time) {
    return new Promise(function (resolve) {
        let tween = game.add.tween(droid).to({
            x: x,
            y: y
        }, time, Phaser.Easing.Linear.None, true, 0);

        game.add.tween(droid.shadow).to({
            x: x + game.shadowXOffset * 5,
            y: y + game.shadowXOffset * 5
        }, time, Phaser.Easing.Linear.None, true, 0);

        tween.onComplete.add(function () {
            resolve();
        })
    })
}

function LaunchReloadDrone(startPoint, up) {
    // todo повторяющий код, этот метод есть в глобальной игре для дигера
    let shadowDrone = game.flyObjectsLayer.create(startPoint.x + game.shadowXOffset * 3, startPoint.y + game.shadowYOffset * 3, 'digger');
    shadowDrone.scale.set(0.05);
    shadowDrone.anchor.setTo(0.5);
    shadowDrone.alpha = 0;
    shadowDrone.tint = 0x000000;
    game.physics.enable(shadowDrone, Phaser.Physics.ARCADE);

    let equipDrone = game.flyObjectsLayer.create(startPoint.x, startPoint.y, 'digger');
    equipDrone.scale.set(0.05);
    equipDrone.anchor.setTo(0.5);
    equipDrone.alpha = 0;
    game.physics.enable(equipDrone, Phaser.Physics.ARCADE);

    equipDrone.shadow = shadowDrone;

    game.add.tween(equipDrone.shadow).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);
    game.add.tween(equipDrone).to({angle: 360}, 3000, Phaser.Easing.Linear.None, true, 0, 0, false).loop(true);

    game.add.tween(equipDrone.shadow).to({alpha: 0.3}, 700, Phaser.Easing.Linear.None, true, 0);
    game.add.tween(equipDrone).to({alpha: 1}, 700, Phaser.Easing.Linear.None, true, 0);

    if (up) {
        equipDrone.scale.set(0.1);
        equipDrone.shadow.scale.set(0.1);
        equipDrone.shadow.x = startPoint.x + game.shadowXOffset * 5;
        equipDrone.shadow.y = startPoint.y + game.shadowYOffset * 5;
    } else {
        game.add.tween(equipDrone.shadow.scale).to({x: 0.1, y: 0.1}, 700, Phaser.Easing.Linear.None, true, 0);
        game.add.tween(equipDrone.scale).to({x: 0.1, y: 0.1}, 700, Phaser.Easing.Linear.None, true, 0);

        game.add.tween(equipDrone.shadow).to({
            x: startPoint.x + game.shadowXOffset * 5,
            y: startPoint.y + game.shadowYOffset * 5
        }, 700, Phaser.Easing.Linear.None, true, 0);
    }

    return equipDrone
}