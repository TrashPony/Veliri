function RotateUnit() {
    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {

                    var unit = game.units[x][y];

                    if (unit.spriteAngle === undefined) {
                        unit.spriteAngle = unit.rotate;
                    }

                    if (unit.spriteAngle < unit.rotate) {
                        unit.spriteAngle += 1;
                    } else {
                        if (unit.spriteAngle > unit.rotate) {
                            unit.spriteAngle -= 1;
                        }
                    }

                    unit.shadow.loadTexture('tank360', unit.spriteAngle + 224); //todo смещение для тестового спрайта
                    unit.sprite.loadTexture('tank360', unit.spriteAngle + 224); //todo смещение для тестового спрайта
                }
            }
        }
    }
}