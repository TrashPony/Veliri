function AnimateUseUnitEquip(jsonMessage) {
    console.log(jsonMessage);

    //let equipBox = document.getElementById(JSON.parse(jsonMessage).applied_equip.id + ":equip"); // id:equip
    //RemoveEquipCell(equipBox);

    let equip = JSON.parse(jsonMessage).applied_equip; // id:equip
    let unit = GetGameUnitID(JSON.parse(jsonMessage).unit.id);
    unit.effects = JSON.parse(jsonMessage).unit.effects;

    if (equip.name === "repair_kit") {
        repairKitAnimate(unit);
    }

    if (equip.name === "energy_shield") {
        energyShieldAnimate(unit);
    }

}

function repairKitAnimate(unit) {
    let repair = game.make.sprite(0, -50, 'RepairKit', 0);
    repair.animations.add('RepairKit');
    repair.animations.play('RepairKit', 20, true, false);
    repair.anchor.set(0.5);

    // таймер для repair таймер делает альфа:0 в течение 500 кадров хз что за линия автостарт через 1000мс //
    let tween = game.add.tween(repair).to({alpha: 0}, 500, Phaser.Easing.Linear.None, true, 1000);
    // функция выполняемая после завершение tween таймера в данном случае удаление спрайта анимации //
    tween.onComplete.add(function (repair) {
        repair.destroy();
    }, repair);

    unit.sprite.addChild(repair);
}

function energyShieldAnimate(unit) {
    if (unit.sprite.shield === undefined) {
        let shield = game.make.sprite(0, 0, 'EnergyShield', 0);
        shield.animations.add('EnergyShield_activated');
        shield.animations.play('EnergyShield_activated', 20, true, false);
        shield.anchor.set(0.5);

        unit.sprite.addChild(shield);

        unit.sprite.shield = shield;
    }
}