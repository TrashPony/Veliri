function DamageText(targetUnit, damage) {
    let colorText;

    if (damage > 0) {
        damage = "-" + damage;
        colorText = "#C00";
    } else if (damage === 0) {
        damage = "0";
        colorText = "#05C"
    }

    let style = {font: "24px Finger Paint", fill: colorText};
    let damageText = game.add.text(targetUnit.sprite.x + 20, targetUnit.sprite.y - 50, damage, style);
    damageText.setShadow(1, -1, 'rgba(0,0,0,0.5)', 0);

    let tween = game.add.tween(damageText).to({y: targetUnit.sprite.y - 75}, 750, Phaser.Easing.Linear.None, true, 100);
    tween.onComplete.add(function (damageText) {
        let tween = game.add.tween(damageText).to({alpha: 0}, 500, Phaser.Easing.Linear.None, true, 10);
        tween.onComplete.add(function (damageText) {
            damageText.destroy();
        })
    }, damageText);
}