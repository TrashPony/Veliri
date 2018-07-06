function FillingConstructorTable(shipBody) {

    UpdateCells(1, "inventoryEquip", shipBody.equippingI);
    UpdateCells(2, "inventoryEquip", shipBody.equippingII);
    UpdateCells(3, "inventoryEquip", shipBody.equippingIII);
    UpdateCells(5, "inventoryEquip", shipBody.equippingV);

    UpdateCells(3, "inventoryEquip", shipBody.weapons);
    /* вепоны надо делать отдельно т.к. храняться отдельно*/

    UpdateShipIcon(shipBody)
}

function UpdateShipIcon(shipBody) {
    let unitIcon = document.getElementById("UnitIcon");
    unitIcon.style.backgroundImage = "url(/assets/" + shipBody.name + ".png)";
}

function UpdateCells(typeSlot, idPrefix, shipSlots) {
    for (let slot in shipSlots) {
        if (shipSlots.hasOwnProperty(slot)) {

            let cell = document.getElementById(idPrefix + slot + shipSlots[slot].type_slot);

            if (cell) {

                cell.slot = shipSlots[slot];

                if (cell.slot.hasOwnProperty("weapon")) {
                    UpdateWeapon(cell);
                } else {
                    UpdateEquips(cell);
                }

                cell.onclick = function () {
                    console.log("NumberSlot " + this.Number + " typeSlot " + this.type)
                };

                cell.onmouseover = function () {
                    this.style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
                    this.style.cursor = "pointer";
                };

                cell.onclick = function () {
                    console.log("NumberSlot " + this.Number + " typeSlot " + this.type)
                };
            } else {
                cell.style.backgroundImage = null;
                cell.innerHTML = "";

                cell.onclick = null;
            }
        }
    }
}

function UpdateEquips(cell) {
    cell.className = "inventoryEquipping active";

    cell.onmouseout = function () {
        this.style.boxShadow = "0 0 0px 0px rgb(0, 0, 0)";
        this.style.cursor = "auto";
    };

    if (cell.slot.equip !== null) {
        cell.style.backgroundImage = "url(/assets/" + cell.slot.equip.name + ".png)";
    }
}

function UpdateWeapon(cell) {
    cell.className = "inventoryEquipping active weapon";

    cell.onmouseout = function () {
        this.style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
        this.style.cursor = "auto";
    };

    if (cell.slot.weapon !== null) {
        cell.style.backgroundImage = "url(/assets/" + cell.slot.weapon.name + ".png)";
    }

    if (cell.ammoCell === null || cell.ammoCell === undefined) {
        let ammoCell = document.createElement("div");
        ammoCell.slot = cell.slot;
        ammoCell.className = "inventoryAmmoCell";

        ammoCell.onclick = function (event) {
            console.log("ammo");
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true)
        };
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

        if (cell.slot.ammo !== null) {
            ammoCell.style.backgroundImage = "url(/assets/" + cell.slot.ammo.name + ".png)";
            ammoCell.innerHTML = "<span class='QuantityItems'>" + cell.slot.ammo_quantity + "</span>";
        }

        cell.appendChild(ammoCell);
        cell.ammoCell = ammoCell;
    } else {
        if (cell.slot.ammo !== null) {
            cell.ammoCell.style.backgroundImage = "url(/assets/" + cell.slot.ammo.name + ".png)";
            cell.ammoCell.innerHTML = "<span class='QuantityItems'>" + cell.slot.ammo_quantity + "</span>";
        } else {
            cell.ammoCell.style.backgroundImage = null;
            cell.ammoCell.innerHTML = "";
        }
    }
}