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
    UpdateCells(1, 5, "inventoryEquip", shipBody.equippingI);
    UpdateCells(2, 5, "inventoryEquip", shipBody.equippingII);
    UpdateCells(3, 5, "inventoryEquip", shipBody.equippingIII);
    UpdateCells(5, 2, "inventoryEquip", shipBody.equippingV);

    UpdateCells(3, 5, "inventoryEquip", shipBody.weapons); /* вепоны надо делать отдельно т.к. храняться отдельно*/

    UpdateShipIcon(shipBody)
}

function UpdateShipIcon(shipBody) {
    let unitIcon = document.getElementById("UnitIcon");
    unitIcon.style.backgroundImage = "url(/assets/" + shipBody.name + ".png)";
}

function UpdateCells(typeSlot, count, idPrefix, shipSlots) {
    for (let i = 1; i <= count; i++) {

        let cell = document.getElementById(idPrefix + Number(i) + typeSlot);

        if (shipSlots.hasOwnProperty(i)) {

            cell.slot = shipSlots[i];

            if (cell.slot.hasOwnProperty("weapon")) {
                cell.className = "inventoryEquipping active weapon";

                if (cell.slot.weapon !== null){
                    cell.style.backgroundImage = "url(/assets/" + cell.slot.weapon.name + ".png)";
                }

                if (cell.slot.ammo != null) {
                    // todo ячейка и бекграунд для аммо
                }

            } else {
                cell.className = "inventoryEquipping active";
            }

            if (shipSlots.hasOwnProperty(i).equip) {
                //todo загрузка бекграунда и фукнций
            }

            cell.onclick = function () {
                console.log("NumberSlot " + this.Number + " typeSlot " + this.type)
            };
        } else {
            cell.slot = null;

            cell.style.backgroundImage = null;
            cell.innerHTML = "";

            cell.onclick = null;
        }
    }
}