function UpdateCells(typeSlot, idPrefix, shipSlots, classPrefix) {
    for (let slot in shipSlots) {
        if (shipSlots.hasOwnProperty(slot)) {

            let cell = document.getElementById(idPrefix + slot + typeSlot);

            if (cell) {
                cell.slotData = JSON.stringify(shipSlots[slot]);

                if (shipSlots[slot].hasOwnProperty("weapon")) {
                    UpdateWeapon(cell, classPrefix, typeSlot);
                } else {
                    UpdateEquips(cell, classPrefix, typeSlot);
                }

                cell.onmousemove = function (e) {
                    if (JSON.parse(this.slotData).equip) {
                        let equipSlot = JSON.parse(this.slotData);
                        equipSlot.item = JSON.parse(this.slotData).equip;
                        equipSlot.type = "equip";
                        ItemOverTip(e, equipSlot)
                    } else if (JSON.parse(this.slotData).weapon) {
                        let weaponSlot = JSON.parse(this.slotData);
                        weaponSlot.item = JSON.parse(this.slotData).weapon;
                        weaponSlot.type = "weapon";
                        ItemOverTip(e, weaponSlot)
                    }
                }
            } else {
                cell.style.backgroundImage = null;
                cell.innerHTML = "";
                cell.onclick = null;
            }
        }
    }
}

function UpdateEquips(cell, classPrefix, typeSlot) {
    cell.className = classPrefix + " active";
    cell.style.boxShadow = "0 0 10px rgba(0,0,0,1)";

    CreateEquipInBody(JSON.parse(cell.slotData));
    if (JSON.parse(cell.slotData).mining) {
        cell.style.boxShadow = "rgb(173, 177, 26) 0px 0px 4px 3px";
    }

    $(cell).draggable({
        disabled: false,
        start: function () {
            if ($(cell).hasClass('inventoryEquipping')) {
                $(cell).data("slotData", {
                    event: "RemoveMotherShipEquip",
                    parent: "Constructor",
                    equipSlot: Number(JSON.parse(cell.slotData).number_slot),
                    equipType: Number(typeSlot),
                });
            } else if ($(cell).hasClass('UnitEquip')) {
                let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;
                $(cell).data("slotData", {
                    event: "RemoveUnitEquip",
                    parent: "Constructor",
                    equipSlot: Number(JSON.parse(cell.slotData).number_slot),
                    equipType: Number(typeSlot),
                    unitSlot: unitSlot,
                });
            }
        },
        revert: "invalid",
        zIndex: 999,
        helper: 'clone',
        appendTo: "body",
    });

    $(cell).droppable({
        drop: function (event, ui) {
            $('.ui-selected').removeClass('ui-selected');
            let draggable = ui.draggable;
            let slotData = draggable.data("slotData");

            if (slotData.data && slotData.data.type === "equip") {
                if ($(cell).hasClass('inventoryEquipping')) {
                    inventorySocket.send(JSON.stringify({
                        event: "SetMotherShipEquip",
                        equip_id: Number(slotData.data.item.id),
                        inventory_slot: Number(slotData.number),
                        equip_slot: Number(JSON.parse(cell.slotData).number_slot),
                        equip_slot_type: Number(typeSlot),
                        source: slotData.parent,
                    }));
                } else if ($(cell).hasClass('UnitEquip')) {
                    let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;
                    inventorySocket.send(JSON.stringify({
                        event: "SetUnitEquip",
                        equip_id: Number(slotData.data.item.id),
                        inventory_slot: Number(slotData.number),
                        equip_slot: Number(JSON.parse(cell.slotData).number_slot),
                        equip_slot_type: Number(typeSlot),
                        unit_slot: Number(unitSlot),
                        source: slotData.parent,
                    }));
                }
                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    });

    if (classPrefix === "inventoryEquipping") {
        cell.onclick = EquipMSMenu;
    } else {
        cell.onclick = EquipUnitMenu;
    }

    cell.onmouseover = function () {

        this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        this.style.cursor = "pointer";
        
        // TODO РЕФАКТОРИНГ
        for (let i = 0; i < document.getElementById('inventoryStorageInventory').childNodes.length; i++) {
            let cell = document.getElementById('inventoryStorageInventory').childNodes[i];

            if (cell.slotData && JSON.parse(cell.slotData).item.type_slot === typeSlot) {
                cell.className = "InventoryCell hover";
            } else if (cell.slotData && JSON.parse(cell.slotData).item.type_slot !== typeSlot) {
                cell.className = "InventoryCell notAllow";
            }
        }

        for (let i = 0; i < document.getElementById('inventoryStorage').childNodes.length; i++) {
            let cell = document.getElementById('inventoryStorage').childNodes[i];

            if (cell.slotData && JSON.parse(cell.slotData).item.type_slot === typeSlot) {
                cell.className = "InventoryCell hover";
            } else if (cell.slotData && JSON.parse(cell.slotData).item.type_slot !== typeSlot) {
                cell.className = "InventoryCell notAllow";
            }
        }
    };

    cell.onmouseout = function () {
        InventoryCellsReset();

        this.style.cursor = "auto";
        this.style.boxShadow = "0 0 10px rgba(0,0,0,1)";

        if (JSON.parse(cell.slotData).mining) {
            cell.style.boxShadow = "rgb(173, 177, 26) 0px 0px 4px 3px";
        }

        let inventoryTip = document.getElementById("InventoryTipOver");
        if (inventoryTip) {
            inventoryTip.remove()
        }
    };

    if (JSON.parse(cell.slotData).equip !== null) {
        cell.style.backgroundImage = "url(/assets/units/equip/icon/" + JSON.parse(cell.slotData).equip.name + ".png)";
        cell.innerText = "";
        CreateHealBar(cell, "equip", true);
    } else {
        cell.style.backgroundImage = null;

        if (typeSlot === 1) {
            cell.innerText = "I";
        } else if (typeSlot === 2) {
            cell.innerText = "II";
        } else if (typeSlot === 3) {
            cell.innerText = "III";
        } else if (typeSlot === 4) {
            cell.innerText = "IV";
        } else if (typeSlot === 5) {
            cell.innerText = "V";
        }
    }
}

function UpdateWeapon(cell, classPrefix) {
    cell.className = classPrefix + " active weapon";
    cell.style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
    cell.innerHTML = '';

    $(cell).draggable({
        disabled: false,
        start: function () {
            if ($(cell).hasClass('inventoryEquipping')) {
                $(cell).data("slotData", {
                    event: "RemoveMotherShipWeapon",
                    parent: "Constructor",
                    equipSlot: Number(JSON.parse(cell.slotData).number_slot),
                });
            } else if ($(cell).hasClass('UnitEquip')) {
                let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;
                $(cell).data("slotData", {
                    event: "RemoveUnitWeapon",
                    parent: "Constructor",
                    equipSlot: Number(JSON.parse(cell.slotData).number_slot),
                    unitSlot: unitSlot,
                });
            }
        },
        revert: "invalid",
        zIndex: 999,
        helper: 'clone',
        appendTo: "body",
    });

    $(cell).droppable({
        drop: function (event, ui) {
            $('.ui-selected').removeClass('ui-selected');
            let draggable = ui.draggable;
            let slotData = draggable.data("slotData");

            if (slotData.data && slotData.data.type === "weapon") {
                if ($(cell).hasClass('inventoryEquipping')) {
                    inventorySocket.send(JSON.stringify({
                        event: "SetMotherShipWeapon",
                        weapon_id: Number(slotData.data.item.id),
                        inventory_slot: Number(slotData.number),
                        equip_slot: Number(JSON.parse(cell.slotData).number_slot),
                        source: slotData.parent,
                    }));
                } else if ($(cell).hasClass('UnitEquip')) {
                    let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;
                    inventorySocket.send(JSON.stringify({
                        event: "SetUnitWeapon",
                        weapon_id: Number(slotData.data.item.id),
                        inventory_slot: Number(slotData.number),
                        equip_slot: Number(JSON.parse(cell.slotData).number_slot),
                        unit_slot: Number(unitSlot),
                        source: slotData.parent,
                    }));
                }
                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    });

    if (classPrefix === "inventoryEquipping") {
        cell.onclick = WeaponMSMenu;
    } else {
        cell.onclick = WeaponUnitMenu;
    }

    cell.onmouseover = function () {

        this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        this.style.cursor = "pointer";

        // TODO РЕФАКТОРИНГ
        for (let i = 0; i < document.getElementById('inventoryStorageInventory').childNodes.length; i++) {
            let cell = document.getElementById('inventoryStorageInventory').childNodes[i];
            let slotData = JSON.parse(cell.slotData);

            if (slotData && slotData.type === "weapon") {

                let bodyData;
                if (classPrefix === "inventoryEquipping") {
                    bodyData = JSON.parse(document.getElementById("MSIcon").slotData).body;
                } else {
                    bodyData = document.getElementById("UnitIcon").unitBody;
                }

                if (bodyData) {
                    if (bodyData.standard_size_big && slotData.item.standard_size === 3) {
                        cell.className = "InventoryCell hover";
                        continue
                    } else {
                        cell.className = "InventoryCell notAllow";
                    }

                    if (bodyData.standard_size_medium && slotData.item.standard_size === 2) {
                        cell.className = "InventoryCell hover";
                        continue
                    } else {
                        cell.className = "InventoryCell notAllow";
                    }

                    if (bodyData.standard_size_small && slotData.item.standard_size === 1) {
                        cell.className = "InventoryCell hover";
                    } else {
                        cell.className = "InventoryCell notAllow";
                    }
                }
            } else if (cell.slotData && JSON.parse(cell.slotData).type !== "weapon") {
                cell.className = "InventoryCell notAllow";
            }
        }

        for (let i = 0; i < document.getElementById('inventoryStorage').childNodes.length; i++) {
            let cell = document.getElementById('inventoryStorage').childNodes[i];
            let slotData = JSON.parse(cell.slotData);

            if (slotData && slotData.type === "weapon") {

                let bodyData;
                if (classPrefix === "inventoryEquipping") {
                    bodyData = JSON.parse(document.getElementById("MSIcon").slotData).body;
                } else {
                    bodyData = document.getElementById("UnitIcon").unitBody;
                }

                if (bodyData) {
                    if (bodyData.standard_size_big && slotData.item.standard_size === 3) {
                        cell.className = "InventoryCell hover";
                        continue
                    } else {
                        cell.className = "InventoryCell notAllow";
                    }

                    if (bodyData.standard_size_medium && slotData.item.standard_size === 2) {
                        cell.className = "InventoryCell hover";
                        continue
                    } else {
                        cell.className = "InventoryCell notAllow";
                    }

                    if (bodyData.standard_size_small && slotData.item.standard_size === 1) {
                        cell.className = "InventoryCell hover";
                    } else {
                        cell.className = "InventoryCell notAllow";
                    }
                }
            } else if (cell.slotData && JSON.parse(cell.slotData).type !== "weapon") {
                cell.className = "InventoryCell notAllow";
            }
        }
    };

    cell.onmouseout = function () {
        InventoryCellsReset();

        this.style.cursor = "auto";
        this.style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
        let inventoryTip = document.getElementById("InventoryTipOver");
        if (inventoryTip) {
            inventoryTip.remove()
        }
    };

    if (cell.ammoCell) {
        cell.ammoCell.remove();
        cell.ammoCell = null;
        cell.innerHTML = "";
    }

    if (JSON.parse(cell.slotData).weapon !== null) {
        cell.style.backgroundImage = "url(/assets/units/weapon/" + JSON.parse(cell.slotData).weapon.name + ".png)";
    } else {
        cell.style.backgroundImage = null;
    }

    if (cell.ammoCell === null || cell.ammoCell === undefined) {

        let ammoCell = CreateAmmoCell(cell, classPrefix, JSON.parse(cell.slotData).weapon);
        cell.appendChild(ammoCell);
        cell.ammoCell = ammoCell;

    } else {
        if (JSON.parse(cell.slotData).ammo !== null) {
            cell.ammoCell.style.backgroundImage = "url(/assets/" + JSON.parse(cell.slotData).ammo.name + ".png)";
            cell.ammoCell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(cell.slotData).ammo_quantity + "</span>";
            cell.ammoCell.slotData = cell.slotData;

            if (classPrefix === "inventoryEquipping") {
                cell.ammoCell.onclick = AmmoMSMenu;
            } else {
                cell.ammoCell.onclick = AmmoUnitMenu;
            }

        } else {
            cell.ammoCell.style.backgroundImage = null;
            cell.ammoCell.innerHTML = "";
        }
    }

    if (JSON.parse(cell.slotData).weapon !== null) {
        CreateHealBar(cell, "weapon", true);
    }
}

function CreateAmmoCell(cell, classPrefix, weapon) {
    let ammoCell = document.createElement("div");
    ammoCell.slotData = cell.slotData;
    ammoCell.className = "inventoryAmmoCell " + classPrefix;

    $(ammoCell).draggable({
        disabled: false,
        start: function () {
            if ($(ammoCell).hasClass('inventoryEquipping')) {
                $(ammoCell).data("slotData", {
                    event: "RemoveMotherShipAmmo",
                    parent: "Constructor",
                    equipSlot: Number(JSON.parse(ammoCell.slotData).number_slot),
                });
            } else if ($(ammoCell).hasClass('UnitEquip')) {
                let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;
                $(ammoCell).data("slotData", {
                    event: "RemoveUnitAmmo",
                    parent: "Constructor",
                    equipSlot: Number(JSON.parse(ammoCell.slotData).number_slot),
                    unitSlot: unitSlot,
                });
            }
        },
        revert: "invalid",
        zIndex: 999,
        helper: 'clone',
        appendTo: "body",
    });

    $(ammoCell).droppable({
        drop: function (event, ui) {
            $('.ui-selected').removeClass('ui-selected');
            let draggable = ui.draggable;
            let slotData = draggable.data("slotData");

            if (slotData.data && slotData.data.type === "ammo") {
                if ($(ammoCell).hasClass('inventoryAmmoCell') && $(ammoCell).hasClass('inventoryEquipping')) {
                    inventorySocket.send(JSON.stringify({
                        event: "SetMotherShipAmmo",
                        ammo_id: Number(slotData.data.item.id),
                        inventory_slot: Number(slotData.number),
                        equip_slot: Number(JSON.parse(ammoCell.slotData).number_slot),
                        source: slotData.parent,
                    }));
                } else if ($(ammoCell).hasClass('inventoryAmmoCell') && $(ammoCell).hasClass('UnitEquip')) {
                    let unitSlot = JSON.parse(document.getElementById("ConstructorUnit").slotData).number_slot;
                    inventorySocket.send(JSON.stringify({
                        event: "SetUnitAmmo",
                        ammo_id: Number(slotData.data.item.id),
                        inventory_slot: Number(slotData.number),
                        equip_slot: Number(JSON.parse(ammoCell.slotData).number_slot),
                        unit_slot: Number(unitSlot),
                        source: slotData.parent,
                    }));
                }
                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    });

    if (classPrefix === "inventoryEquipping") {
        ammoCell.onclick = AmmoMSMenu;
    } else {
        ammoCell.onclick = AmmoUnitMenu;
    }

    ammoCell.onmouseover = function (event) {
        this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        this.style.cursor = "pointer";
        if (weapon) {
            // TODO РЕФАКТОРИНГ
            for (let i = 0; i < document.getElementById('inventoryStorageInventory').childNodes.length; i++) {
                let cell = document.getElementById('inventoryStorageInventory').childNodes[i];
                let slotData = JSON.parse(cell.slotData);
                if (slotData && slotData.type === "ammo" && weapon.type === slotData.item.type && weapon.standard_size === slotData.item.standard_size) {
                    cell.className = "InventoryCell hover";
                } else if (slotData) {
                    cell.className = "InventoryCell notAllow";
                }
            }

            for (let i = 0; i < document.getElementById('inventoryStorage').childNodes.length; i++) {
                let cell = document.getElementById('inventoryStorage').childNodes[i];
                let slotData = JSON.parse(cell.slotData);
                if (slotData && slotData.type === "ammo" && weapon.type === slotData.item.type && weapon.standard_size === slotData.item.standard_size) {
                    cell.className = "InventoryCell hover";
                } else if (slotData) {
                    cell.className = "InventoryCell notAllow";
                }
            }
        }
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };

    ammoCell.onmousemove = function (event) {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };

    ammoCell.onmouseout = function (event) {
        this.style.boxShadow = "0 0 5px 3px rgb(200, 200, 0)";
        this.style.cursor = "auto";
        InventoryCellsReset();

        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };

    if (JSON.parse(ammoCell.slotData).ammo !== null) {
        ammoCell.style.backgroundImage = "url(/assets/units/ammo/" + JSON.parse(ammoCell.slotData).ammo.name + ".png)";
        ammoCell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(ammoCell.slotData).ammo_quantity + "</span>";
    }

    return ammoCell
}