function SelectUnit() {

    RemoveSelect();

    CreateUnitSubMenu(this);

    MarkUnitSelect(this);

    field.send(JSON.stringify({
        event: "SelectUnit",
        x: Number(this.info.x),
        y: Number(this.info.y)
    }));
}