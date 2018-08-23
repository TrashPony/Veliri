function CalculateHealBar(unit) {
    let hp = unit.hp;
    let maxHP = unit.body.max_hp;
    let percentageHeal = hp * 100 / maxHP;
    let healSprite = unit.sprite.heal;

    healSprite.scale.x = percentageHeal / 100;

    let ColorOffset = 255 - (255 * percentageHeal / 100);

    let green;
    let blue = "00";
    let red;

    if (percentageHeal < 50 && percentageHeal >= 25) {
        green = Math.round((255 + ColorOffset) / 2 + 30);
        red = Math.round((255 + ColorOffset) / 2 + 30);
    } else {
        if (percentageHeal < 75 && percentageHeal >= 50) {
            if (percentageHeal >= 60) {
                green = Math.round((255 + ColorOffset) / 1.5 + 10);
            } else {
                green = Math.round((255 + ColorOffset) / 1.5);
            }
            red = Math.round((255 + ColorOffset) / 1.5);
        } else {
            if (percentageHeal < 25 && percentageHeal > 15) {
                green = Math.round(255 - ColorOffset);
                red = Math.round(ColorOffset) + 30;
            } else {
                if (percentageHeal >= 75 && percentageHeal < 85) {
                    green = Math.round(255 - ColorOffset);
                    red = Math.round(ColorOffset) + 30;
                } else {
                    green = Math.round(255 - ColorOffset);
                    red = Math.round(ColorOffset);
                }
            }
        }
    }

    if (green < 16) {
        green = "0" + green.toString(16);
    } else {
        green = green.toString(16)
    }

    if (red < 16) {
        red = "0" + red.toString(16);
    } else {
        red = red.toString(16)
    }

    healSprite.tint = "0x" + red + green + blue;
}