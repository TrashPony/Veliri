function UnitMouseOver() {
    unitTip(this);
    this.sprite.healBar.visible = true;

    if (game.Phase === "targeting") {
        field.send(JSON.stringify({
            event: "GetTargetZone",
            x: Number(this.x),
            y: Number(this.y),
            to_x: Number(this.x),
            to_y: Number(this.y)
        }));
    }

    if (this.target) {
        MarkTarget(this.target)
    }
}

function UnitMouseOut() {
    TipOff();
    DeleteMarkTarget(this.target);
    this.sprite.healBar.visible = false;


    if (game.SelectLayer.children.length === 0 && game.Phase === "targeting") {
        RemoveTargetLine();
    }
}

