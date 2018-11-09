function SelectUnit() {

    RemoveSelect();

    CreateUnitSubMenu(this);

    MarkUnitSelect(this, 1);

    Fire(this, game.map.OneLayerMap[12][2]);

    if (game.Phase === "targeting") {
        field.send(JSON.stringify({
            event: "SelectWeapon",
            q: Number(this.q),
            r: Number(this.r)
        }));
    }

    field.send(JSON.stringify({
        event: "SelectUnit",
        q: Number(this.q),
        r: Number(this.r)
    }));
}