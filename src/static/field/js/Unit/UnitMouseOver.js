function UnitMouseOver() {
    unitTip(this);

    if (game.Phase === "targeting") {
        field.send(JSON.stringify({
            event: "GetTargetZone",
            x: Number(this.info.x),
            y: Number(this.info.y),
            to_x: Number(this.info.x),
            to_y: Number(this.info.y)
        }));
    }

    if (this.info.target) {
        MarkTarget(this.info.target)
    }
}

function UnitMouseOut() {
    TipOff();
    DeleteMarkTarget(this.info.target);

    if (game.SelectLayer.children.length === 0 && game.Phase === "targeting") {
        RemoveTargetLine();
    }
}

