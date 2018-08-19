function UpdateCells(typeSlot, idPrefix, shipSlots, classPrefix) {
    for (let slot in shipSlots) {
        if (shipSlots.hasOwnProperty(slot)) {

            let cell = document.getElementById(idPrefix + slot + typeSlot);

            if (cell) {
                cell.slotData = JSON.stringify(shipSlots[slot]);

                if (JSON.parse(cell.slotData).hasOwnProperty("weapon")) {
                    UpdateWeapon(cell, classPrefix);
                } else {
                    UpdateEquips(cell, classPrefix);
                }
                cell.onmouseover = function () {
                    this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
                    this.style.cursor = "pointer";
                };
            } else {
                cell.style.backgroundImage = null;
                cell.innerHTML = "";
                cell.onclick = null;
            }
        }
    }
}

function UpdateEquips(cell, classPrefix) {
    cell.className = classPrefix + " active";

    cell.onclick = EquipMSMenu;

    cell.onmouseout = function () {
        this.style.boxShadow = "0 0 0px 0px rgb(0, 0, 0)";
        this.style.cursor = "auto";
    };

    if (JSON.parse(cell.slotData).equip !== null) {
        cell.style.backgroundImage = "url(/assets/" + JSON.parse(cell.slotData).equip.name + ".png)";
    } else {
        cell.style.backgroundImage = null;
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

    cell.onmouseout = function () {
        this.style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
        this.style.cursor = "auto";
    };

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
        } else {
            cell.ammoCell.style.backgroundImage = null;
            cell.ammoCell.innerHTML = "";
        }
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
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
        this.style.cursor = "pointer";
    };

    ammoCell.onmouseout = function (event) {
        event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
        this.style.boxShadow = "0 0 5px 3px rgb(200, 200, 0)";
        this.style.cursor = "auto";
    };

    if (JSON.parse(ammoCell.slotData).ammo !== null) {
        ammoCell.style.backgroundImage = "url(/assets/" + JSON.parse(ammoCell.slotData).ammo.name + ".png)";
        ammoCell.innerHTML = "<span class='QuantityItems'>" + JSON.parse(ammoCell.slotData).ammo_quantity + "</span>";
    }

    return ammoCell
}