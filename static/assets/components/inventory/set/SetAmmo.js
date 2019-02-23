function SetAmmo(ammo, slot, source) {

    let ammoCells = document.getElementsByClassName("inventoryAmmoCell");

    for (let i = 0; ammoCells && i < ammoCells.length; i++) {
        ammoCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        ammoCells[i].style.cursor = "pointer";
        ammoCells[i].onmouseout = null;

        if (ammoCells[i].className === "inventoryAmmoCell inventoryEquipping") {
            ammoCells[i].onclick = function (event) {
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
                inventorySocket.send(JSON.stringify({
                    event: "SetMotherShipAmmo",
                    ammo_id: Number(ammo.id),
                    inventory_slot: Number(slot),
                    equip_slot: Number(JSON.parse(this.slotData).number_slot),
                    source: source,
                }));

                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        } else {
            ammoCells[i].onclick = function (event) {
                let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;
                event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
                inventorySocket.send(JSON.stringify({
                    event: "SetUnitAmmo",
                    ammo_id: Number(ammo.id),
                    inventory_slot: Number(slot),
                    equip_slot: Number(JSON.parse(this.slotData).number_slot),
                    unit_slot: Number(unitSlot),
                    source: source,
                }));

                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    }
}