function AttackPhase(jsonMessage) {
    console.log(jsonMessage);

    let resultBattle = JSON.parse(jsonMessage).result_battle;
    let resultEquip = JSON.parse(jsonMessage).result_equip;
    let watchNode = JSON.parse(jsonMessage).watch_node;

    for (let i = 0; i < resultBattle.length; i++) {

        let unit = GetGameUnitID(resultBattle[i].attack_unit.id);

        Fire(unit);
        UpdateUnit(unit);
        for (let j = 0; j < resultBattle[i].targets_units.length; j++) {
            let targetUnit = GetGameUnitID(resultBattle[i].targets_units[j].unit.id);
            setTimeout(function () {
                Explosion(targetUnit);
                DamageText(targetUnit, resultBattle[i].targets_units[j].damage);
                UpdateUnit(targetUnit);
            }, 500);
        }
    }
}