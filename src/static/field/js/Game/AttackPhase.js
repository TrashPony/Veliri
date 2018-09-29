function AttackPhase(jsonMessage) {
    console.log(jsonMessage);

    let resultBattle = JSON.parse(jsonMessage).result_battle;
    let resultEquip = JSON.parse(jsonMessage).result_equip;
    let watchNode = JSON.parse(jsonMessage).watch_node;

    for (let i = 0; i < resultBattle.length; i++) {

        let unit = GetGameUnitID(resultBattle[i].attack_unit.id);

        if (unit) {
            Fire(unit);
            UpdateUnit(resultBattle[i].attack_unit);
        }

        for (let j = 0; j < resultBattle[i].targets_units.length; j++) {

            let targetUnit = GetGameUnitID(resultBattle[i].targets_units[j].unit.id);
            
            if (targetUnit) {
                setTimeout(function () {
                    Explosion(targetUnit);
                    DamageText(targetUnit, resultBattle[i].targets_units[j].damage);
                    UpdateUnit(resultBattle[i].targets_units[j].unit);
                }, 500);
            }
        }
    }

    for (let i = 0; i < resultEquip.length; i++) {

    }

    UpdateWatchZone(watchNode)
}