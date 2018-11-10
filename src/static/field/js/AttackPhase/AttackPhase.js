function AttackPhase(jsonMessage) {

    // todo удалять таргет лайны

    let resultBattle = JSON.parse(jsonMessage).result_battle;

    if (resultBattle != null && resultBattle.length > 0) {
        playAttacks(PlayAttack(resultBattle))
    }
}

function playAttacks(resultBattle, yieldValue) {
    let next = resultBattle.next(yieldValue);
    if (!next.done) {
        next.value.then(
            result => playAttacks(resultBattle, result)
        );
    }
}

function* PlayAttack(resultBattle) {
    for (let i = 0; i < resultBattle.length; i++) {

        let resultAction = resultBattle[i];

        console.log(resultAction);

        let unit = GetGameUnitID(resultAction.attack_unit.id);
        let target = null;
        let unitInTargetCoordinate = null;

        if (resultAction.target.type !== "hide") { // смотрим видит игрок цель юнита
            target = game.map.OneLayerMap[resultAction.target.q][resultAction.target.r];
            unitInTargetCoordinate = GetGameUnitXY(resultAction.target.q, resultAction.target.r);
        }

        if (resultAction.weapon_slot.weapon) {
            if (unit && unitInTargetCoordinate) { // юнит стреляет и юзер видит все
                yield Fire(unit, unitInTargetCoordinate).then(function () {
                    return UpdateTargetUnits(resultAction)
                });
            } else if (unit && target) {
                yield Fire(unit, target, "coordinate").then(function () {
                    return UpdateTargetUnits(resultAction)
                });
            }

            if (unit && !target) { // видит кто стреляет, не видит куда стреляет
                // endWatchAttack последняя координата видимости полета снаряда
                let endWatchAttack = game.map.OneLayerMap[resultAction.end_watch_attack.q][resultAction.end_watch_attack.r];
                yield Fire(unit, endWatchAttack, "inFog").then(function () {
                    return UpdateTargetUnits(resultAction)
                });
            }

            if (!unit && target) { // не видит кто стреляет, видит куда

                // startWatchAttack первая координата откуда видно стрельбу
                let startWatchAttack = game.map.OneLayerMap[resultAction.start_watch_attack.q][resultAction.start_watch_attack.r];
                let weapon = resultAction.weapon_slot.weapon;

                if (unitInTargetCoordinate) {
                    yield OutFogFire(startWatchAttack, unitInTargetCoordinate, weapon).then(function () {
                        return UpdateTargetUnits(resultAction)
                    });
                } else {
                    yield OutFogFire(startWatchAttack, target, weapon, "coordinate").then(function () {
                        return UpdateTargetUnits(resultAction)
                    });
                }
            }
        }

        if (resultAction.equip_slot.equip) {  // юнит использует снарягу и юзер видит все
            if (unit && unitInTargetCoordinate) {

            } else if (unit && target) {

            }
        }
    }

    yield new Promise(resolve => setTimeout(resolve, 300));
}

function UpdateTargetUnits(resultAction) {
    for (let j = 0; j < resultAction.targets_units.length; j++) {
        let targetUnit = GetGameUnitID(resultAction.targets_units[j].unit.id);
        if (targetUnit) {

            DamageText(targetUnit, resultAction.targets_units[j].damage);
            UpdateUnit(resultAction.targets_units[j].unit);
            CalculateHealBar(targetUnit);

            if (targetUnit.hp <= 0) {

                let explosion = game.effectsLayer.create(targetUnit.sprite.x, targetUnit.sprite.y, 'explosion_2');
                explosion.anchor.setTo(0.5);

                if (targetUnit.body.mother_ship) {
                    explosion.scale.set(3);
                } else {
                    explosion.scale.set(1.5);
                }

                explosion.animations.add('explosion_2');
                explosion.animations.play('explosion_2', 10, false, true);

                game.add.tween(targetUnit.sprite).to({alpha: 0}, 40, Phaser.Easing.Linear.None, true, 0);

            } else {
                game.add.tween(targetUnit.sprite.healBar).to({alpha: 1}, 100, Phaser.Easing.Linear.None, true);
                setTimeout(function () {
                    game.add.tween(targetUnit.sprite.healBar).to({alpha: 0}, 100, Phaser.Easing.Linear.None, true);
                }, 2000);
            }
        }
    }

    for (let i in resultAction.watch_node) { // обновляем видимость
        if (resultAction.watch_node.hasOwnProperty(i)) {
            UpdateWatchZone(resultAction.watch_node[i]);
        }
    }

    return new Promise((resolve) => {
        setTimeout(function () {
            resolve();
        }, 2000)
    })
}