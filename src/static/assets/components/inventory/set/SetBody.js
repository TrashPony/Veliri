function SetBody(body, slot) {
    if (body.mother_ship) {
        let shipIcon = document.getElementById("MSIcon");

        shipIcon.className = "UnitIconSelect";
        shipIcon.onclick = function () {

            inventorySocket.send(JSON.stringify({
                event: "SetMotherShipBody",
                id_body: Number(body.id),
                inventory_slot: Number(slot)
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        }
    } else {
        let unitIcon = document.getElementById("UnitIcon");
        if (unitIcon) {
            let unitData = JSON.parse(document.getElementById("ConstructorUnit").slotData);

            unitIcon.className = "UnitIconSelect";
            unitIcon.onclick = function () {

                inventorySocket.send(JSON.stringify({
                    event: "SetUnitBody",
                    id_body: Number(body.id),
                    inventory_slot: Number(slot),
                    unit_slot: Number(unitData.number_slot)
                }));

                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    }
}