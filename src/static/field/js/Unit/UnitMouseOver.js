function UnitMouseOver() {
    unitTip(this);

    CalculateHealBar(this);
    game.add.tween(this.sprite.healBar).to({alpha: 1}, 100, Phaser.Easing.Linear.None, true);

    field.send(JSON.stringify({
        event: "GetTargetZone",
        x: Number(this.x),
        y: Number(this.y),
        to_x: Number(this.x),
        to_y: Number(this.y)
    }));

    if (this.target) {
        MarkTarget(this.target)
    }
}

function UnitMouseOut() {
    TipOff();
    DeleteMarkTarget(this.target);
    game.add.tween(this.sprite.healBar).to({alpha: 0}, 100, Phaser.Easing.Linear.None, true);


    if (game.SelectLayer.children.length === 0) {
        RemoveTargetLine();
    }
}

