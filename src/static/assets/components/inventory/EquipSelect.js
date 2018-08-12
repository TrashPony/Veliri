function AmmoRemove(event) {

    if (JSON.parse(this.slotData).ammo !== null) {

        let slot = JSON.parse(this.slotData).number_slot;

        let removeFunction = function () {

            inventorySocket.send(JSON.stringify({
                event: "RemoveMotherShipAmmo",
                equip_slot: Number(slot)
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        RemoveTip(event, removeFunction);
    }
}

function WeaponRemove(event) {

    if (JSON.parse(this.slotData).weapon !== null) {

        let slot = JSON.parse(this.slotData).number_slot;

        let removeFunction = function () {

            inventorySocket.send(JSON.stringify({
                event: "RemoveMotherShipWeapon",
                equip_slot: Number(slot)
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        RemoveTip(event, removeFunction);
    }
}

function EquipRemove(event) {

    if (JSON.parse(this.slotData).equip !== null) {

        let slot = JSON.parse(this.slotData).number_slot;
        let type = JSON.parse(this.slotData).type_slot;

        let removeFunction = function () {

            inventorySocket.send(JSON.stringify({
                event: "RemoveMotherShipEquip",
                equip_slot: Number(slot),
                equip_slot_type: Number(type)
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        RemoveTip(event, removeFunction);
    }
}

function BodyRemove(event) {

    let removeFunction = function () {

        inventorySocket.send(JSON.stringify({
            event: "RemoveMotherShipBody"
        }));

        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };

    RemoveTip(event, removeFunction);
}