function SelectUnit() {

    RemoveSelect();

    if (!this.info.action && game.user.name === this.info.owner) {
        CreateUnitSubMenu(this);
    }

    MarkUnitSelect(this);

    field.send(JSON.stringify({
        event: "SelectUnit",
        x: Number(this.info.x),
        y: Number(this.info.y)
    }));
}