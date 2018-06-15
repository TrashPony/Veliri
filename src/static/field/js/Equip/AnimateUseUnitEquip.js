function AnimateUseUnitEquip(jsonMessage) {
    console.log(jsonMessage);

    var equipBox = document.getElementById(JSON.parse(jsonMessage).applied_equip.id  + ":equip"); // id:equip
    RemoveEquipCell(equipBox);

    var equip = JSON.parse(jsonMessage).applied_equip; // id:equip
    var unit = GetGameUnitID(JSON.parse(jsonMessage).unit.id);
    unit.effect = JSON.parse(jsonMessage).unit.effect;

    if (equip.type === "repair_kit") {
        repairKitAnimate(unit);
    }

    if (equip.type === "energy_shield") {
        energyShieldAnimate(unit);
    }

}

function repairKitAnimate(unit) {
    var repair = game.make.sprite(0, -50, 'RepairKit', 0);
    repair.animations.add('RepairKit');
    repair.animations.play('RepairKit', 20, false, true);
    repair.anchor.set(0.5);

    unit.sprite.addChild(repair);
}

function energyShieldAnimate(unit) {
    // TODO проверять естли уже этот эфект на анимации или нет если есть то нечего не делать
    var shield = game.make.sprite(0, 0, 'EnergyShield', 0);
    shield.animations.add('EnergyShield_activated');
    shield.animations.play('EnergyShield_activated', 20, true, false);
    shield.anchor.set(0.5);

    unit.sprite.addChild(shield);
}