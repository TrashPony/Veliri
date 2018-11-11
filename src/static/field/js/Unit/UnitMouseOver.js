function UnitMouseOver() {
    VisibleUnitStatus(this)
}

function UnitMouseOut() {
    HideUnitStatus(this)
}

function VisibleUnitStatus(unit, type) {

    unitTip(unit);

    CalculateHealBar(unit);
    game.add.tween(unit.sprite.healBar).to({alpha: 1}, 100, Phaser.Easing.Linear.None, true);

    if (type !== "inTarget") {
        RemoveTargetLine();

        field.send(JSON.stringify({
            event: "GetTargetZone",
            q: Number(unit.q),
            r: Number(unit.r),
            to_q: Number(unit.q),
            to_r: Number(unit.r)
        }));
    }

    /*if (unit.target) {
        MarkTarget(unit.target)
    }*/
}

function HideUnitStatus(unit, type) {
    TipOff();
    //DeleteMarkTarget(unit.target);
    game.add.tween(unit.sprite.healBar).to({alpha: 0}, 100, Phaser.Easing.Linear.None, true);

    if (type !== "inTarget") {
        RemoveTargetLine();
    }
}

