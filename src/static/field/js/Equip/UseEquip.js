function UsedEquip() {

    RemoveSelect();

    field.send(JSON.stringify({
        event: "UseMapEquip",
        target_q: Number(this.targetQ),
        target_r: Number(this.targetR),
        q: Number(this.unitQ),
        r: Number(this.unitR),
        equip_id: Number(this.equipID),
        equip_type: Number(this.typeSlot),
        number_slot: Number(this.numberSlot)
    }));
}