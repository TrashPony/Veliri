function SelectUnit() {

    RemoveSelect();

    CreateUnitSubMenu(this);

    MarkUnitSelect(this, 1);

    field.send(JSON.stringify({
        event: "SelectUnit",
        x: Number(this.x),
        y: Number(this.y)
    }));
}