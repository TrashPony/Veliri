function UpdateRotateUnit() {
    for (let q in game.units) {
        if (game.units.hasOwnProperty(q)) {
            for (let r in game.units[q]) {
                if (game.units[q].hasOwnProperty(r)) {
                    let unit = game.units[q][r];

                    if (unit) {
                        rotateUnitSprites(unit.sprite.unitBody.angle, unit.rotate + 90, unit)
                    }
                }
            }
        }
    }
}