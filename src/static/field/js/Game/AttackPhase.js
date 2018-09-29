function AttackPhase(jsonMessage) {
    console.log(jsonMessage);

    let resultBattle = JSON.parse(jsonMessage).result_battle;
    let resultEquip = JSON.parse(jsonMessage).result_equip;
    let watchNode = JSON.parse(jsonMessage).watch_node;

    for (let i = 0; i < resultBattle.length; i++) {

        let unit = GetGameUnitID(resultBattle[i].attack_unit.id);

        Fire(unit);

        for (let j = 0; j < resultBattle[i].targets_units.length; j++) {

            let targetUnit = GetGameUnitID(resultBattle[i].targets_units[j].unit.id);
            setTimeout(function () {
                Explosion(targetUnit);

                let damage;
                let colorText;

                if (resultBattle[i].targets_units[j].damage > 0) {
                    damage = "-" + resultBattle[i].targets_units[j].damage
                    colorText = "#C00";
                } else if (resultBattle[i].targets_units[j].damage === 0) {
                    damage = "0";
                    colorText = "#05C"
                }

                let style = {font: "24px Finger Paint", fill: colorText};
                let damageText = game.add.text(targetUnit.sprite.x + 20, targetUnit.sprite.y - 50, damage, style);
                damageText.setShadow(1, -1, 'rgba(0,0,0,0.5)', 0);

                let tween = game.add.tween(damageText).to({y: targetUnit.sprite.y - 75}, 750, Phaser.Easing.Linear.None, true, 100);
                tween.onComplete.add(function (damageText) {
                    let tween = game.add.tween(damageText).to({alpha: 0}, 500, Phaser.Easing.Linear.None, true, 10);
                    tween.onComplete.add(function (damageText) {
                        damageText.destroy();
                    })
                }, damageText);

            }, 500);
        }
    }
}