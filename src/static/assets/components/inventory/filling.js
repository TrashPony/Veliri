function FillingInventory(jsonData) {
    let event = JSON.parse(jsonData).event;

    if (event === "openInventory" || event === "UpdateSquad") {
        console.log(jsonData);
        let squad = JSON.parse(jsonData).squad;
        FillingInventoryTable(squad.inventory);

        if (squad.mather_ship.body != null) {
            FillingConstructorTable(squad.mather_ship.body)
        }
    }
}

function FillingInventoryTable(inventoryItems) {
    for (let i = 1; i <= 40; i++) {
        let cell = document.getElementById("inventory " + i + 6);

        if (inventoryItems.hasOwnProperty(i) && inventoryItems[i].item !== null) {

            cell.slot = inventoryItems[i];
            cell.number = i;

            cell.style.backgroundImage = "url(/assets/" + cell.slot.item.name + ".png)";
            cell.innerHTML = "<span class='QuantityItems'>" + cell.slot.quantity + "</span>";

            cell.onclick = SelectInventoryItem
        } else {
            cell.slot = null;

            cell.style.backgroundImage = null;
            cell.innerHTML = "";

            cell.onclick = null;
        }
    }
}

function FillingSquadTable() {

}

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

    if (cell.slot.equip !== null) {
        cell.style.backgroundImage = "url(/assets/" + cell.slot.equip.name + ".png)";
    }
}

function UpdateWeapon(cell) {
    cell.className = "inventoryEquipping active weapon";

    if (cell.slot.weapon !== null) {
        cell.style.backgroundImage = "url(/assets/" + cell.slot.weapon.name + ".png)";
    }

    if (cell.slot.ammo != null) {
        // todo ячейка и бекграунд для аммо
    }
}