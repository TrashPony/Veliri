function Fire(unit, target) {
    let connectPoints = PositionAttachSprite(unit.rotate, unit.sprite.unitBody.width / 2);

    /*let fireMuzzle = game.weaponEffectsLayer.create(unit.sprite.x + connectPoints.x, unit.sprite.y + connectPoints.y, 'fireMuzzle_1');
    if (unit.rotate > 180) {
        fireMuzzle.angle = unit.rotate - 360;
    } else {
        fireMuzzle.angle = unit.rotate;
    }

    fireMuzzle.anchor.setTo(0, 0.5);
    fireMuzzle.animations.add('fireMuzzle_1', [2,1,0]);
    fireMuzzle.animations.play('fireMuzzle_1', 10, false, true);
*/
    let weapon;

    for (let i in unit.body.weapons){
        if (unit.body.weapons.hasOwnProperty(i) && unit.body.weapons[i].weapon) {
            weapon = unit.body.weapons[i].weapon
        }
    }

    console.log(weapon);

    if (weapon.type === "missile") {
        launchRocket(unit.sprite.x + connectPoints.x, unit.sprite.y + connectPoints.y, unit.sprite.weapon.angle, game.map.OneLayerMap[13][2]);
    }
}