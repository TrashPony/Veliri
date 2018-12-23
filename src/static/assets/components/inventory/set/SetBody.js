function SetBody(body, slot) {
    if (body.mother_ship) {
        let shipIcon = document.getElementById("MSIcon");

        if (shipIcon) {
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
        }
    } else {
        let unitIcon = document.getElementById("UnitIcon");
        if (unitIcon) {
            let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;

            unitIcon.className = "UnitIconSelect";
            unitIcon.onclick = function () {

                inventorySocket.send(JSON.stringify({
                    event: "SetUnitBody",
                    id_body: Number(body.id),
                    inventory_slot: Number(slot),
                    unit_slot: Number(unitSlot)
                }));

                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    }
}