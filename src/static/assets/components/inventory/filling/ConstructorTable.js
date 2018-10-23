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
    unitIcon.slotData = JSON.stringify(ms.body);

    unitIcon.onclick = BodyMSMenu;

    unitIcon.addEventListener("mousemove", function (e) {
        let slot = {};
        slot.item = JSON.parse(this.slotData);
        slot.type = "body";
        slot.hp = ms.hp;
        ItemOverTip(e, slot)
    });
    unitIcon.addEventListener("mouseout", function () {
        let inventoryTip = document.getElementById("InventoryTipOver");
        if (inventoryTip) {
            inventoryTip.remove()
        }
    });
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