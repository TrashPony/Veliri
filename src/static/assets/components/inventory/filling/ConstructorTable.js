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
    unitIcon.style.backgroundImage = "url(/assets/units/body/" + ms.body.name + ".png)";
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

    let slotData = {};
    slotData.unit = ms;

    CreateThoriumSlots(unitIcon, ms);
    UpdateWeaponIcon(unitIcon, "weaponIcon", slotData);
    CreateHealBar(unitIcon, "body", true);
}

function CreateThoriumSlots(unitIcon, ms) {

    let div = document.createElement("div");
    div.id = "thorium";

    let speedEfficiency = document.createElement("div");
    speedEfficiency.id = "speedEfficiency";
    div.appendChild(speedEfficiency);

    let thoriumEfficiency = document.createElement("div");
    thoriumEfficiency.id = "thoriumEfficiency";
    div.appendChild(thoriumEfficiency);

    let countSlot = 0;
    let fullCount = 0;

    for (let i in ms.body.thorium_slots) {

        countSlot++;

        let thoriumSlots = document.createElement("div");
        thoriumSlots.className = "thoriumSlots";

        thoriumSlots.innerHTML = ms.body.thorium_slots[i].count + "/" + ms.body.thorium_slots[i].max_count;
        thoriumSlots.count = ms.body.thorium_slots[i].count;
        thoriumSlots.maxCount = ms.body.thorium_slots[i].max_count;
        thoriumSlots.numberSlot = i;

        if (ms.body.thorium_slots[i].count > 0) {
            fullCount++;
            thoriumSlots.style.backgroundImage = "url(/assets/resource/enriched_thorium.png)";
        }

        thoriumSlots.onmouseover = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

        };

        thoriumSlots.onmousemove = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);

        };

        thoriumSlots.onclick = function () {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            if (ms.body.thorium_slots[i].count > 0) {
                inventorySocket.send(JSON.stringify({
                    event: "RemoveThorium",
                    thorium_slot: Number(i)
                }));
            }
        };

        div.appendChild(thoriumSlots);
    }

    let efficiencyCalc = 0;
    let thoriumEfficiencyCalc = 0;
    if (fullCount > 0) {
        efficiencyCalc = (fullCount * 100) / countSlot;
        thoriumEfficiencyCalc = (100 - efficiencyCalc) + 100;
    }

    if (efficiencyCalc <= 33) {
        speedEfficiency.style.color = "#FF0000";
    } else if (efficiencyCalc <= 66) {
        speedEfficiency.style.color = "#FFF000";
    } else if (efficiencyCalc === 100) {
        speedEfficiency.style.color = "#00FF00";
    }

    thoriumEfficiency.innerHTML = (thoriumEfficiencyCalc).toFixed(0) + "%";
    speedEfficiency.innerHTML = efficiencyCalc.toFixed(0) + "%";

    div.style.left = "calc(50% - " + (countSlot * 34) / 2 + "px)";

    unitIcon.appendChild(div);
}

function ItemOverTip(e, slot) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = e.clientY + "px";
        inventoryTip.style.left = e.clientX + "px";
    } else {
        InventorySelectTip(slot, e.clientX, e.clientY, true, false);
    }
}