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