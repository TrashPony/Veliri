function AnimateUseUnitEquip(jsonMessage) {
    console.log(jsonMessage);

    var equipBox = document.getElementById(JSON.parse(jsonMessage).applied_equip.id + ":equip"); // id:equip
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
    repair.animations.play('RepairKit', 20, true, false);
    repair.anchor.set(0.5);

    // таймер для repair таймер делает альфа:0 в течение 500 кадров хз что за линия автостарт через 1000мс //
    var tween = game.add.tween(repair).to({alpha: 0}, 500, Phaser.Easing.Linear.None, true, 1000);
    // функция выполняемая после завершение tween таймера в данном случае удаление спрайта анимации //
    tween.onComplete.add(function (repair) {
        repair.destroy();
    }, repair);

    unit.sprite.addChild(repair);
}

function energyShieldAnimate(unit) {
    if (unit.sprite.shield === undefined) {
        var shield = game.make.sprite(0, 0, 'EnergyShield', 0);
        shield.animations.add('EnergyShield_activated');
        shield.animations.play('EnergyShield_activated', 20, true, false);
        shield.anchor.set(0.5);

        unit.sprite.addChild(shield);

        unit.sprite.shield = shield;
    }
}