function UsedEquip() {

    RemoveSelect();

    field.send(JSON.stringify({
        event: "UseEquip",
        x: Number(this.gameCoordinateX),
        y: Number(this.gameCoordinateY),
        equip_id: Number(this.equipID)
    }));
}