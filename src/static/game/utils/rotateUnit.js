function rotateUnitSprites(spriteRotate, needRotate, unit) {

    spriteRotate = Math.round(spriteRotate);
    needRotate = Math.round(needRotate);

    if (spriteRotate < 0) {
        spriteRotate += 360;
    }

    if (needRotate < 0) {
        needRotate += 360;
    }

    if (needRotate > 360) {
        needRotate -= 360;
    }

    if (spriteRotate !== needRotate) {
        if (directionRotate(spriteRotate, needRotate)) {
            SetAngle(unit, spriteRotate + 1)
        } else {
            SetAngle(unit, spriteRotate - 1)
        }
    }
}
// метод вычесления в какую сторону меньше поворачивать обьект
function directionRotate(spriteAngle, rotate) {
    // true ++
    // false --
    let count = 0;
    let direction;

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

function SetAngle(unit, angle) {
    unit.sprite.unitBody.angle = angle;
    unit.sprite.bodyShadow.angle = angle;

    if (unit.sprite.weapon) {
        unit.sprite.weaponShadow.angle = angle;
        unit.sprite.weapon.angle = angle;
    }
}