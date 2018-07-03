function SelectInventoryItem(e) {
    DestroyInventoryTip();
    InventoryTip(this.item, e.clientX, e.clientY);

    if (this.type === "body") {
        SelectInventoryBody(this.item, this.slot);
    }

    if (this.type === "ammo") {
        console.log(this.item);
        console.log(this.quantity);
    }

    if (this.type === "weapon") {
        console.log(this.item);
        console.log(this.quantity);
    }

    if (this.type === "equip") {
        console.log(this.item);
        console.log(this.quantity);
    }
}

function SelectInventoryBody(body, slot) {
    if (body.mother_ship) {
        let shipIcon = document.getElementById("UnitIcon");
        shipIcon.className = "UnitIconSelect";
        shipIcon.onclick = function () {

            inventorySocket.send(JSON.stringify({
                event: "SetMotherShipBody",
                id_body: body.id,
                inventory_slot: slot
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        }
    }
}