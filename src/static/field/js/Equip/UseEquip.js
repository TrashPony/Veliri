function UsedEquip(cell) {
    RemoveSelect();

    // todo придумать окно подтверждения типо "Применить equip_type к выбраному юниту?"

    var equip = cell.equip;
    var unit = cell.unit.info;

    field.send(JSON.stringify({
        event: "UseEquip",
        x: Number(unit.x),
        y: Number(unit.y),
        unit_id: Number(unit.id),
        equip_id: Number(equip.id)
    }));
}