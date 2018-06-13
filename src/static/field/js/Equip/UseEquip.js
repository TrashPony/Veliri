function UsedEquip(unit, equip) {

    RemoveSelect();

    field.send(JSON.stringify({
        event: "UseEquip",
        x: Number(unit.x),
        y: Number(unit.y),
        unit_id: Number(unit.id),
        equip_id: Number(equip.id)
    }));
}