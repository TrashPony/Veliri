function CreateTargetLine(unitStat) {
    let unit = GetGameUnitID(unitStat.id);

    if (!unit) {
        unit = GetStorageUnit(unitStat.id)
    }

    if (unit.targetLine) {
        unit.targetLine.destroy();
    }

    if (unit.targetsEquipLine && unit.targetsEquipLine.length > 0) {
        for (let i = 0; i < unit.targetsEquipLine.length; i++) {
            unit.targetsEquipLine[i].destroy();
        }
    }

    unit.targetsEquipLine = [];

    if (unit.target) {
        createWeaponLine(unit)
    }

    // TODO эфекты наведения на линиях что бы было понятно какая линия какому эквпу, рефакторинг и красивость линий)
    for (let i in unit.body.equippingII) {
        if (unit.body.equippingII.hasOwnProperty(i) && unit.body.equippingII[i].equip && unit.body.equippingII[i].target) {
            createEquipLine(unit, unit.body.equippingII[i].target)
        }
    }

    for (let i in unit.body.equippingIII) {
        if (unit.body.equippingIII.hasOwnProperty(i) && unit.body.equippingIII[i].equip && unit.body.equippingIII[i].target) {
            createEquipLine(unit, unit.body.equippingIII[i].target)
        }
    }
}

function createWeaponLine(unit) {

    let xy = GetXYCenterHex(unit.target.q, unit.target.r);
    let targetLine = game.add.graphics(0, 0);
    targetLine.lineStyle(3, 0xff0000, 0.3);
    targetLine.moveTo(unit.sprite.x, unit.sprite.y);
    targetLine.lineTo(xy.x, xy.y);
    targetLine.endFill();

    unit.targetLine = targetLine;
}

function createEquipLine(unit, target) {
    let xy = GetXYCenterHex(target.q, target.r);
    let targetLine = game.add.graphics(0, 0);
    targetLine.lineStyle(3, 0x0000FF, 0.3);
    targetLine.moveTo(unit.sprite.x, unit.sprite.y);
    targetLine.lineTo(xy.x, xy.y);
    targetLine.endFill();

    unit.targetsEquipLine.push(targetLine);
}

function RemoveTargetsLine() {
    for (let x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (let y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {
                    let unit = game.units[x][y];
                    if (unit.targetLine) {
                        unit.targetLine.destroy();
                    }
                    if (unit.targetsEquipLine && unit.targetsEquipLine.length > 0) {
                        for (let i = 0; i < unit.targetsEquipLine.length; i++) {
                            unit.targetsEquipLine[i].destroy();
                        }
                    }
                }
            }
        }
    }
}