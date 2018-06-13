function UpdateRotateUnit() {
    for (var x in game.units) {
        if (game.units.hasOwnProperty(x)) {
            for (var y in game.units[x]) {
                if (game.units[x].hasOwnProperty(y)) {

                    var unit = game.units[x][y];

                    if (unit.spriteAngle === undefined) {
                        unit.spriteAngle = unit.rotate;
                    }

                    if (unit.spriteAngle !== unit.rotate) {
                        if (directionRotate(unit.spriteAngle, unit.rotate)) {
                            if (unit.spriteAngle >= 360) {
                                unit.spriteAngle = 0;
                            } else {
                                unit.spriteAngle++;
                            }
                        } else {
                            if (unit.spriteAngle <= 0) {
                                unit.spriteAngle = 360;
                            } else {
                                unit.spriteAngle--;
                            }
                        }
                        unit.RotateUnit(unit.spriteAngle);
                    }
                }
            }
        }
    }
}

function directionRotate(spriteAngle, rotate) {
    // true ++
    // false --
    var count = 0;
    var direction;

    if (spriteAngle < rotate) {
        for (; spriteAngle < rotate; spriteAngle++) {
            count++;
            direction = true;
        }
    } else {
        for (; spriteAngle > rotate; rotate++) {
            count++;
            direction = false;
        }
    }

    if (direction) {
        return count <= 180
    } else {
        return !(count <= 180)
    }
}

function RotateUnit(unit, angle) {
    for (var sprite in unit) {
        if (unit.hasOwnProperty(sprite) && unit[sprite] !== null && unit[sprite].hasOwnProperty('_frame')) {
            unit[sprite].frame = angle;
        }
    }
}