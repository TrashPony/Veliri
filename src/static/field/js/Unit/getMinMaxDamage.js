function getMinMaxDamage(unit, targetUnit) {

    let maxDamage;
    let minDamage;

    for (let weaponSlot in unit.body.weapons) { // оружие может быть только 1 под диз доке, масив это обман
        if (unit.body.weapons.hasOwnProperty(weaponSlot) && unit.body.weapons[weaponSlot].ammo) {
            maxDamage = unit.body.weapons[weaponSlot].ammo.max_damage;
            minDamage = unit.body.weapons[weaponSlot].ammo.min_damage;
        }
    }

    if (minDamage - targetUnit.armor < 0) {
        minDamage = 0
    } else {
        minDamage = minDamage - targetUnit.armor
    }

    if (maxDamage - targetUnit.armor < 0) {
        maxDamage = 0
    } else {
        maxDamage = maxDamage - targetUnit.armor
    }

    return  minDamage + " - " + maxDamage
}