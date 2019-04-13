function AttackPhase(jsonMessage) {

    RemoveTargetsLine();

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

        let unit = GetGameUnitID(resultAction.attack_unit.id);
        let target = null;
        let unitInTargetCoordinate = null;

        if (resultAction.target.type !== "hide") { // смотрим видит игрок цель юнита
            target = game.map.OneLayerMap[resultAction.target.q][resultAction.target.r];
            unitInTargetCoordinate = GetGameUnitXY(resultAction.target.q, resultAction.target.r);
        }

        if (resultAction.weapon_slot.weapon) { // юнит стреляет
            yield fireWeapon(unit, unitInTargetCoordinate, resultAction, target).then(function () {
                return UpdateTargetUnits(resultAction)
            });
        }

        if (resultAction.equip_slot.equip) {  // юнит использует снарягу
            yield fireEquip(unit, unitInTargetCoordinate, resultAction, target).then(function () {
                return UpdateTargetUnits(resultAction)
            });
        }

        if (resultAction.reload) { // юнит перезаряжается
            yield reloadUnit(unit, unitInTargetCoordinate, resultAction).then(function () {
                return UpdateTargetUnits(resultAction)
            });
        }

        // если нечего не использовалось, значит это конец боя и отрабатывают наложеные эффекты
        if (!resultAction.equip_slot.equip && !resultAction.weapon_slot.weapon && !resultAction.reload) {
            yield UpdateTargetUnits(resultAction)
        }
    }

    yield new Promise(resolve => setTimeout(resolve, 300));
}