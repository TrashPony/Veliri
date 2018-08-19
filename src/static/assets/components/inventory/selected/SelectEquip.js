// todo вероятно можно обьеденить эти методы в имя рефакторинга
function AmmoMSMenu(event) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

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

function AmmoUnitMenu(event) {
    event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    console.log("dfdf");
    let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;

    if (JSON.parse(this.slotData).ammo !== null) {
        let slot = JSON.parse(this.slotData).number_slot;
        let removeFunction = function () {

            inventorySocket.send(JSON.stringify({
                event: "RemoveUnitAmmo",
                equip_slot: Number(slot),
                unit_slot: unitSlot
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };
        RemoveTip(event, removeFunction);
    }
}

function WeaponMSMenu(event) {

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

function WeaponUnitMenu(event) {
    let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;

    if (JSON.parse(this.slotData).weapon !== null) {

        let slot = JSON.parse(this.slotData).number_slot;

        let removeFunction = function () {

            inventorySocket.send(JSON.stringify({
                event: "RemoveUnitWeapon",
                equip_slot: Number(slot),
                unit_slot: unitSlot
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        RemoveTip(event, removeFunction);
    }
}

function EquipMSMenu(event) {

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

function EquipUnitMenu(event) {
    let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;

    if (JSON.parse(this.slotData).equip !== null) {

        let slot = JSON.parse(this.slotData).number_slot;
        let type = JSON.parse(this.slotData).type_slot;

        let removeFunction = function () {

            inventorySocket.send(JSON.stringify({
                event: "RemoveUnitEquip",
                equip_slot: Number(slot),
                equip_slot_type: Number(type),
                unit_slot: unitSlot
            }));

            DestroyInventoryClickEvent();
            DestroyInventoryTip();
        };

        RemoveTip(event, removeFunction);
    }
}

function BodyMSMenu(event) {

    let removeFunction = function () {

        inventorySocket.send(JSON.stringify({
            event: "RemoveMotherShipBody"
        }));

        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };

    RemoveTip(event, removeFunction);
}

function BodyUnitMenu(event) {

    let numberUnitSlot = JSON.parse(this.slotData).number_slot;

    let removeFunction = function () {
        inventorySocket.send(JSON.stringify({
            event: "RemoveUnitBody",
            unit_slot: Number(numberUnitSlot)
        }));

        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };

    RemoveTip(event, removeFunction);
}