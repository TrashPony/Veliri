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

    if (classPrefix === "inventoryEquipping") {
        cell.onclick = EquipMSMenu;
    } else {
        cell.onclick = EquipUnitMenu;
    }

    cell.onmouseover = function () {

        this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        this.style.cursor = "pointer";

        for (let i = 1; i <= 40; i++) {
            let cell = document.getElementById("inventory " + i + 6);
            if (cell.slotData && JSON.parse(cell.slotData).item.type_slot === typeSlot) {
                cell.className = "InventoryCell hover";
            } else if (cell.slotData && JSON.parse(cell.slotData).item.type_slot !== typeSlot) {
                cell.className = "InventoryCell notAllow";
            }
        }
    };

    cell.onmouseout = function () {
        for (let i = 1; i <= 40; i++) {
            let cell = document.getElementById("inventory " + i + 6);
            cell.className = "InventoryCell";
        }

        this.style.cursor = "auto";
        this.style.boxShadow = "0 0 10px rgba(0,0,0,1)";
        let inventoryTip = document.getElementById("InventoryTipOver");
        if (inventoryTip) {
            inventoryTip.remove()
        }
    };

    if (JSON.parse(cell.slotData).equip !== null) {
        cell.style.backgroundImage = "url(/assets/" + JSON.parse(cell.slotData).equip.name + ".png)";
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

    if (classPrefix === "inventoryEquipping") {
        cell.onclick = WeaponMSMenu;
    } else {
        cell.onclick = WeaponUnitMenu;
    }

    cell.onmouseover = function () {

        this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        this.style.cursor = "pointer";

        for (let i = 1; i <= 40; i++) {
            let cell = document.getElementById("inventory " + i + 6);
            if (cell.slotData && JSON.parse(cell.slotData).type === "weapon") {
                cell.className = "InventoryCell hover";
            } else if (cell.slotData && JSON.parse(cell.slotData).type !== "weapon") {
                cell.className = "InventoryCell notAllow";
            }
        }
    };

    cell.onmouseout = function () {
        for (let i = 1; i <= 40; i++) {
            let cell = document.getElementById("inventory " + i + 6);
            cell.className = "InventoryCell";
        }

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
        cell.style.backgroundImage = "url(/assets/" + JSON.parse(cell.slotData).weapon.name + ".png)";
    } else {
        cell.style.backgroundImage = null;
    }

    if (cell.ammoCell === null || cell.ammoCell === undefined) {

        let ammoCell = CreateAmmoCell(cell, classPrefix);
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

function CreateAmmoCell(cell, classPrefix) {
    let ammoCell = document.createElement("div");
    ammoCell.slotData = cell.slotData;
    ammoCell.className = "inventoryAmmoCell " + classPrefix;

    if (classPrefix === "inventoryEquipping") {
        ammoCell.onclick = AmmoMSMenu;
    } else {
        ammoCell.onclick = AmmoUnitMenu;
    }

    ammoCell.onmouseover = function (event) {
        this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        this.style.cursor = "pointer";
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };

    ammoCell.onmousemove = function (event) {
        // todo подсказка о типе снаряда
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };

    ammoCell.onmouseout = function (event) {
        this.style.boxShadow = "0 0 5px 3px rgb(200, 200, 0)";
        this.style.cursor = "auto";
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
    };

    if (JSON.parse(ammoCell.slotData).ammo !== null) {
        ammoCell.style.backgroundImage = "url(/assets/" + JSON.parse(ammoCell.slotData).ammo.name + ".png)";
        ammoCell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(ammoCell.slotData).ammo_quantity + "</span>";
    }

    return ammoCell
}