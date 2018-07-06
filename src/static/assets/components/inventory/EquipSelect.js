function AmmoRemove(event) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

    if (this.slot.ammo !== null) {

        let slot = this.slot.number_slot;

        let removeFunction = function () {
            inventorySocket.send(JSON.stringify({
                event: "RemoveMotherShipAmmo",
                equip_slot: slot
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        RemoveTip(event, removeFunction);
    }
}

function WeaponRemove(event) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

    if (this.slot.weapon !== null) {

        let slot = this.slot.number_slot;

        let removeFunction = function () {
            inventorySocket.send(JSON.stringify({
                event: "RemoveMotherShipWeapon",
                equip_slot: slot
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        RemoveTip(event, removeFunction);
    }
}

function EquipRemove(event) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

    if (this.slot.equip !== null) {

        let slot = this.slot.number_slot;
        let type = this.slot.type_slot;

        let removeFunction = function () {
            inventorySocket.send(JSON.stringify({
                event: "RemoveMotherShipEquip",
                equip_slot: slot,
                equip_slot_type: type
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        RemoveTip(event, removeFunction);
    }
}

function BodyRemove(event) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

    let removeFunction = function () {
        inventorySocket.send(JSON.stringify({
            event: "RemoveMotherShipBody"
        }));

        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };

    RemoveTip(event, removeFunction);
}