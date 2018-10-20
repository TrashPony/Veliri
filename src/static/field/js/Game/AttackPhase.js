function AttackPhase(jsonMessage) {

    let resultBattle = JSON.parse(jsonMessage).result_battle;
    let resultEquip = JSON.parse(jsonMessage).result_equip;
    let watchNode = JSON.parse(jsonMessage).watch_node;

    if (resultBattle != null && resultBattle.length > 0) {
        playAttacks(PlayAttack(resultBattle))
    }

    if (resultEquip != null) {
        for (let i = 0; i < resultEquip.length; i++) {

        }
    }

    UpdateWatchZone(watchNode)
}

function playAttacks(resultBattle, yieldValue) {

    let next = resultBattle.next(yieldValue);

    if (!next.done) {
        next.value.then(
            result => playAttacks(resultBattle, result),
            err => resultBattle.throw(err)
        );
    }
}

function* PlayAttack(resultBattle) {
    for (let i = 0; i < resultBattle.length; i++) {
        let unit = GetGameUnitID(resultBattle[i].attack_unit.id);

        if (unit) {
            game.camera.x = unit.sprite.x - game.camera.width / 2; // наводим камеру на место событий
            game.camera.y = unit.sprite.y - game.camera.height / 2;
            setTimeout(function () {
                Fire(unit);
            }, 200);
        }

        for (let j = 0; j < resultBattle[i].targets_units.length; j++) {

            let targetUnit = GetGameUnitID(resultBattle[i].targets_units[j].unit.id);

            if (targetUnit) {

                setTimeout(function () {

                    game.add.tween(targetUnit.sprite.healBar).to({alpha: 1}, 100, Phaser.Easing.Linear.None, true); // показываем хил бар юнита
                    game.camera.x = targetUnit.sprite.x - game.camera.width / 2; // наводим камеру на место событий
                    game.camera.y = targetUnit.sprite.y - game.camera.height / 2;
                    setTimeout(function () {
                        // взрыв происходит там куда упал снаряд
                        Explosion(game.map.OneLayerMap[resultBattle[i].attack_unit.target.q][resultBattle[i].attack_unit.target.r]);
                        DamageText(targetUnit, resultBattle[i].targets_units[j].damage);
                        UpdateUnit(resultBattle[i].targets_units[j].unit);
                        CalculateHealBar(targetUnit);

                        setTimeout(function () {
                            game.add.tween(targetUnit.sprite.healBar).to({alpha: 0}, 100, Phaser.Easing.Linear.None, true);
                        }, 2000);
                    }, 200);
                }, 500);
            }
        }
        UpdateUnit(resultBattle[i].attack_unit);
        yield new Promise(resolve => setTimeout(resolve, 3000));
    }
}