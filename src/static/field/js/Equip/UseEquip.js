function UsedEquip() {

    RemoveSelect();

    field.send(JSON.stringify({
        event: "UseEquip",
        target_x: Number(this.gameCoordinateX),
        target_y: Number(this.gameCoordinateY),
        x: Number(this.unitX),
        y: Number(this.unitY),
        equip_id: Number(this.equipID),
        equip_type: Number(this.typeSlot),
        number_slot: Number(this.numberSlot)
    }));
}