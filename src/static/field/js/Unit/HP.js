function CalculateHealBar(unit) {
    var hp = unit.hp;
    var maxHP = unit.max_hp;
    var percentageHeal = hp * 100 / maxHP;
    var healSprite = unit.sprite.heal;

    healSprite.scale.x = percentageHeal / 100;

    var ColorOffset = 255 - (255 * percentageHeal / 100);

    var blue = Math.round(0xFF - (ColorOffset));
    var green = Math.round(0xFF - (ColorOffset));

    if (blue < 16) {
        blue = "0" + blue.toString(16);
    } else {
        blue = blue.toString(16)
    }

    if (green < 16) {
        green = "0" + green.toString(16);
    } else {
        green = green.toString(16)
    }

    healSprite.tint = "0xFF" + green + blue;
}