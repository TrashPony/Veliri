function StopMining(jsonData) {

    let unit = game.units[jsonData.short_unit.id];
    if (unit) {
        for (let i in unit.miningLaser) {
            if (unit.miningLaser[i] && unit.miningLaser[i].id === "reloadEquip" + unit.id + jsonData.type_slot + jsonData.slot) {

                ShortDirectionRotateTween(unit.miningLaser[i].equipSprite, Phaser.Math.degToRad(0), 500);
                ShortDirectionRotateTween(unit.miningLaser[i].attachPoint, Phaser.Math.degToRad(0), 500);

                unit.miningLaser[i].out.destroy();
                unit.miningLaser[i].in.destroy();
                unit.miningLaser[i] = null;
            }
        }
    }
}