function SetAmmo(ammo, slot) {

    let ammoCells = document.getElementsByClassName("inventoryAmmoCell");

    for (let i = 0; i < ammoCells.length; i++) {

        ammoCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        ammoCells[i].style.cursor = "pointer";
        ammoCells[i].onmouseout = null;
        ammoCells[i].onclick = function (event) {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            inventorySocket.send(JSON.stringify({
                event: "SetMotherShipAmmo",
                ammo_id: Number(ammo.id),
                inventory_slot: Number(slot),
                equip_slot: Number(JSON.parse(this.slotData).number_slot),
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        }
    }
}