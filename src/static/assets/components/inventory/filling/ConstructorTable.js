function ConstructorTable(ms) {

    UpdateCells(1, "inventoryEquip", ms.body.equippingI, "inventoryEquipping");
    UpdateCells(2, "inventoryEquip", ms.body.equippingII, "inventoryEquipping");
    UpdateCells(3, "inventoryEquip", ms.body.equippingIII, "inventoryEquipping");
    UpdateCells(5, "inventoryEquip", ms.body.equippingV, "inventoryEquipping");

    UpdateCells(3, "inventoryEquip", ms.body.weapons, "inventoryEquipping");
    /* вепоны надо делать отдельно т.к. храняться отдельно*/

    UpdateShipIcon(ms)
}

function UpdateShipIcon(ms) {
    let unitIcon = document.getElementById("MSIcon");
    unitIcon.innerHTML = "";
    unitIcon.shipBody = unitIcon;
    unitIcon.style.backgroundImage = "url(/assets/" + ms.body.name + ".png)";
    unitIcon.slotData = JSON.stringify(ms);

    unitIcon.onclick = BodyMSMenu;

    unitIcon.onmousemove = function (e) {
        for (let i = 1; i <= 40; i++) {
            let cell = document.getElementById("inventory " + i + 6);
            if (cell) {
                if (cell.slotData && JSON.parse(cell.slotData).type === "body" && JSON.parse(cell.slotData).item.mother_ship) {
                    cell.className = "InventoryCell hover";
                } else if (cell.slotData && (JSON.parse(cell.slotData).type !== "body" || !JSON.parse(cell.slotData).item.mother_ship)) {
                    cell.className = "InventoryCell notAllow";
                }
            }
        }

        if (unitIcon.shipBody) {
            let slot = {};
            slot.item = ms.body;
            slot.type = "body";
            slot.hp = ms.hp;
            slot.item.name = ms.body.name;
            ItemOverTip(e, slot)
        }
    };

    unitIcon.onmouseout = function () {
        OffTip();
        for (let i = 1; i <= 40; i++) {
            let cell = document.getElementById("inventory " + i + 6);
            if (cell) {
                cell.className = "InventoryCell";
            }
        }
    };

    CreateHealBar(unitIcon, "body", true);
}

function ItemOverTip(e, slot) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = e.clientY + "px";
        inventoryTip.style.left = e.clientX + "px";
    } else {
        InventorySelectTip(slot, e.clientX, e.clientY, true);
    }
}