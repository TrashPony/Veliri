function ConstructorTable(ms) {

    UpdateShipIcon(ms);

    UpdateCells(1, "inventoryEquip", ms.body.equippingI, "inventoryEquipping");
    UpdateCells(2, "inventoryEquip", ms.body.equippingII, "inventoryEquipping");
    UpdateCells(3, "inventoryEquip", ms.body.equippingIII, "inventoryEquipping");
    UpdateCells(5, "inventoryEquip", ms.body.equippingV, "inventoryEquipping");

    UpdateCells(3, "inventoryEquip", ms.body.weapons, "inventoryEquipping");
    /* вепоны надо делать отдельно т.к. храняться отдельно*/

}

function UpdateShipIcon(ms) {
    let unitIcon = document.getElementById("MSIcon");
    unitIcon.innerHTML = '';
    unitIcon.shipBody = unitIcon;
    unitIcon.style.backgroundImage = "url(/assets/units/body/" + ms.body.name + ".png), url(/assets/units/body/" + ms.body.name + "_bottom.png)";

    let mask1 = document.createElement('div');
    mask1.className = 'mask body';
    mask1.id = 'msBodyMask1';
    mask1.style.background = "#" + ms.body_color_1.split('x')[1];
    $(mask1).css("-webkit-mask-image", "url(/assets/units/body/" + ms.body.name + "_mask.png)");

    let mask2 = document.createElement('div');
    mask2.style.opacity = '0.3';
    mask2.className = 'mask body';
    mask2.id = 'msBodyMask2';
    mask2.style.background = "#" + ms.body_color_2.split('x')[1];
    $(mask2).css("-webkit-mask-image", "url(/assets/units/body/" + ms.body.name + "_mask2.png)");

    unitIcon.appendChild(mask2);
    unitIcon.appendChild(mask1);

    unitIcon.slotData = JSON.stringify(ms);
    unitIcon.onclick = BodyMSMenu;

    unitIcon.onmousemove = function (e) {

        let tipFunc = function (id) {
            for (let i = 0; document.getElementById(id) && i < document.getElementById(id).childNodes.length; i++) {
                let inventoryCell = document.getElementById(id).childNodes[i];
                if (!inventoryCell.slotData) continue;
                let slotData = JSON.parse(inventoryCell.slotData);

                if (slotData.type === "body" && slotData.item.mother_ship) {
                    inventoryCell.className = "InventoryCell hover";
                } else if (slotData.type !== "body" || !slotData.item.mother_ship) {
                    inventoryCell.className = "InventoryCell notAllow";
                }
            }
        };

        tipFunc('inventoryStorageInventory');
        tipFunc('inventoryStorage');

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
        InventoryCellsReset();
    };

    let slotData = {};
    slotData.unit = ms;

    if (!document.getElementById('ConstructorUnit'))
        FillParams(ms);

    CreateColorInputs(unitIcon, ms, 0, 'ms');
    CreateThoriumSlots(unitIcon, ms);
    UpdateWeaponIcon(unitIcon, "weaponIcon", slotData);
    CreateHealBar(unitIcon, "body", true);
}

function CreateThoriumSlots(unitIcon, ms) {

    let div = document.createElement("div");
    div.id = "thorium";

    let thoriumWrapper = document.createElement("div");
    thoriumWrapper.id = "thoriumWrapper";
    div.appendChild(thoriumWrapper);

    let statWrapper = document.createElement("div");
    statWrapper.id = "statWrapper";
    thoriumWrapper.appendChild(statWrapper);

    let thoriumSlotsWrapper = document.createElement("div");
    thoriumSlotsWrapper.id = "thoriumSlotsWrapper";
    thoriumWrapper.appendChild(thoriumSlotsWrapper);

    let speedEfficiency = document.createElement("div");
    speedEfficiency.id = "speedEfficiency";
    statWrapper.appendChild(speedEfficiency);

    let thoriumEfficiency = document.createElement("div");
    thoriumEfficiency.id = "thoriumEfficiency";
    statWrapper.appendChild(thoriumEfficiency);

    let countSlot = 0;
    let fullCount = 0;

    for (let i in ms.body.thorium_slots) {

        countSlot++;

        let thoriumSlots = document.createElement("div");
        thoriumSlots.className = "thoriumSlots";

        $(thoriumSlots).droppable({
            drop: function (event, ui) {
                $('.ui-selected').removeClass('ui-selected');

                let draggable = ui.draggable;
                let slotData = draggable.data("slotData");

                if (slotData.data.type === "recycle" && slotData.data.item.name === "enriched_thorium") {
                    inventorySocket.send(JSON.stringify({
                        event: "SetThorium",
                        inventory_slot: Number(slotData.number),
                        thorium_slot: Number(i),
                        source: slotData.parent,
                    }));
                    DestroyInventoryClickEvent();
                    DestroyInventoryTip();
                }
            }
        });

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

        thoriumSlotsWrapper.appendChild(thoriumSlots);
    }

    let efficiencyCalc = 0;
    let thoriumEfficiencyCalc = 0;
    if (fullCount > 0) {
        efficiencyCalc = (fullCount * 100) / countSlot;
        thoriumEfficiencyCalc = (100 - efficiencyCalc);
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

    unitIcon.appendChild(div);
}

function ItemOverTip(e, slot) {
    let inventoryTip = document.getElementById("InventoryTipOver");
    if (inventoryTip) {
        inventoryTip.style.top = e.clientY + "px";
        inventoryTip.style.left = e.clientX + "px";
    } else {
        InventorySelectTip(slot, true, false);
    }
}