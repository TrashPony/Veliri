function SelectUnit() {

    RemoveSelect();

    CreateUnitSubMenu(this);

    MarkUnitSelect(this);

    field.send(JSON.stringify({
        event: "SelectUnit",
        x: Number(this.x),
        y: Number(this.y)
    }));
}