function SquadTable(squad) {
    for (let slot = 1; slot <= 6; slot++) {

        let cell = document.getElementById("squad " + slot + 4); // 4 это тип ячейки

        if (squad && squad.mather_ship != null && squad.mather_ship.body != null && squad.mather_ship.units && squad.mather_ship.units.hasOwnProperty(slot)) {

            let unitSlot = squad.mather_ship.units[slot];
            unitSlot.standardSize = squad.mather_ship.body.equippingIV[slot].standard_size;

            if (cell.className !== "inventoryUnit select") {
                cell.className = "inventoryUnit active";
            }

            if (squad.mather_ship.units[slot].unit !== null && squad.mather_ship.units[slot].unit !== undefined) {
                let unit = squad.mather_ship.units[slot].unit;
                cell.innerHTML = "";
                cell.style.background = "url(/assets/units/body/" + unit.body.name + ".png) center center / 100% no-repeat," +
                    "url(/assets/units/body/" + unit.body.name + "_bottom.png) center center / 100% no-repeat, #4c4c4c";

                let mask1 = document.createElement('div');
                mask1.className = 'mask unit';
                mask1.style.background = "#" + unit.body_color_1.split('x')[1];
                $(mask1).css("-webkit-mask-image", "url(/assets/units/body/" + unit.body.name + "_mask.png)");

                let mask2 = document.createElement('div');
                mask2.style.opacity = '0.3';
                mask2.className = 'mask unit';
                mask2.style.background = "#" + unit.body_color_2.split('x')[1];
                $(mask2).css("-webkit-mask-image", "url(/assets/units/body/" + unit.body.name + "_mask2.png)");

                cell.appendChild(mask2);
                cell.appendChild(mask1);
                UpdateWeaponIcon(cell, "weaponUnitIcon", squad.mather_ship.units[slot], true);

                let constructorUnit = document.getElementById("ConstructorUnit");
                if (constructorUnit && JSON.parse(constructorUnit.slotData).number_slot === slot) {
                    FillingSquadConstructor(unitSlot);
                }
            } else {
                NoActiveUnitCell(unitSlot);
                cell.style.background = null;
                cell.innerHTML = "<span> Ангар </span>"
            }

            let standardSizeBlock = document.createElement("div");
            standardSizeBlock.className = "standardSizeBlock";
            standardSizeBlock.id = 'standardSizeUnitBlock' + slot;

            if (squad.mather_ship.body.equippingIV[slot].standard_size === 1) {
                standardSizeBlock.innerHTML = "S";
                standardSizeBlock.style.color = "#3bff19";
            } else if (squad.mather_ship.body.equippingIV[slot].standard_size === 2) {
                standardSizeBlock.innerHTML = "M";
                standardSizeBlock.style.color = "#ffe418";
            } else if (squad.mather_ship.body.equippingIV[slot].standard_size === 3) {
                standardSizeBlock.innerHTML = "L";
            }

            cell.appendChild(standardSizeBlock);

            cell.slotData = JSON.stringify(unitSlot);
            cell.onclick = OpenUnitEditor
        } else {

            if (document.getElementById('inventoryBox')) {
                cell.slotData = null;
                cell.onclick = null;
                cell.style.backgroundImage = null;
                cell.className = "inventoryUnit noActive";
                cell.innerHTML = "";
                let constructorUnit = document.getElementById("ConstructorUnit");
                if (constructorUnit && JSON.parse(constructorUnit.slotData).number_slot === slot) {
                    constructorUnit.remove();
                }
            }
        }
    }
}

function OpenUnitEditor() {
    let constructorUnit = document.getElementById("ConstructorUnit");
    let unitData = JSON.parse(this.slotData);

    let inventoryUnits = document.getElementsByClassName("inventoryUnit select");
    for (let slot in inventoryUnits) {
        inventoryUnits[slot].className = "inventoryUnit active";
    }

    if (constructorUnit) {
        if (JSON.parse(constructorUnit.slotData).number_slot === unitData.number_slot) {

            constructorUnit.remove();
            document.getElementById('ConstructorBackGround').style.filter = 'unset';
            FillParams(JSON.parse(document.getElementById("MSIcon").slotData));

            return;
        } else {
            constructorUnit.remove();
        }
    }

    this.className = "inventoryUnit select";

    constructorUnit = document.createElement("div");
    constructorUnit.id = "ConstructorUnit";
    constructorUnit.slotData = this.slotData;

    CreateUnitEquipSlots(constructorUnit);
    document.getElementById('ConstructorBackGround').style.filter = 'grayscale(100%) blur(5px)';
    document.getElementById('inventoryBox').appendChild(constructorUnit);

    if (unitData.unit !== null && unitData.unit !== undefined) {
        FillingSquadConstructor(unitData);
    } else {
        let powerPanel = document.getElementById("unitPowerPanel");
        powerPanel.innerHTML = "<span class='Value'>" + 0 + "/" + 0 + "</span>";

        let unitCubePanel = document.getElementById("unitCubePanel");
        unitCubePanel.innerHTML = "<span class='Value'>" + 0 + "/" + 0 + "</span>";

        let unitIcon = document.getElementById("UnitIcon");
        unitIcon.innerHTML = "<span>Место для корпуса</span>";

        FillParams(null);
    }

    let unitIcon = $('#UnitIcon');
    unitIcon.droppable({
        drop: function (event, ui) {
            $('.ui-selected').removeClass('ui-selected');
            let draggable = ui.draggable;
            let slotData = draggable.data("slotData");
            if (slotData.data.type === "body") {
                inventorySocket.send(JSON.stringify({
                    event: "SetUnitBody",
                    id_body: Number(slotData.data.item.id),
                    inventory_slot: Number(slotData.number),
                    unit_slot: Number(unitData.number_slot),
                    source: slotData.parent,
                }));
                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    });

    unitIcon.mouseover(function () {
        let tipFunc = function (id) {
            for (let i = 0; document.getElementById(id) && i < document.getElementById(id).childNodes.length; i++) {
                let inventoryCell = document.getElementById(id).childNodes[i];
                if (!inventoryCell.slotData) continue;
                let slotData = JSON.parse(inventoryCell.slotData);

                if (slotData.type === "body" && !slotData.item.mother_ship && unitData.standardSize >= slotData.item.standard_size) {
                    inventoryCell.className = "InventoryCell hover";
                } else {
                    inventoryCell.className = "InventoryCell notAllow";
                }
            }
        };

        tipFunc('inventoryStorageInventory');
        tipFunc('inventoryStorage');
    });

    unitIcon.mouseout(function () {
        InventoryCellsReset();
    });
}

function CreateUnitEquipSlots(constructorUnit) {
    let powerIcon = document.createElement("div");
    powerIcon.className = "powerIcon";
    constructorUnit.appendChild(powerIcon);

    let weaponTypeIcon = document.createElement("div");
    weaponTypeIcon.id = "weaponTypeUnitIcon";
    constructorUnit.appendChild(weaponTypeIcon);

    let weaponTypePanel = document.createElement("div");
    weaponTypePanel.id = "weaponTypePanel";
    constructorUnit.appendChild(weaponTypePanel);

    let unitPowerPanel = document.createElement("div");
    unitPowerPanel.id = "unitPowerPanel";
    constructorUnit.appendChild(unitPowerPanel);

    let cubeIcon = document.createElement("div");
    cubeIcon.className = "cubeIcon";
    constructorUnit.appendChild(cubeIcon);

    let unitCubePanel = document.createElement("div");
    unitCubePanel.id = "unitCubePanel";
    constructorUnit.appendChild(unitCubePanel);

    let equippingPanelIII = document.createElement("div");
    CreateCells(3, 3, "UnitEquip noActive", "UnitEquip", equippingPanelIII);
    constructorUnit.appendChild(equippingPanelIII);

    let unitIcon = document.createElement("div");
    unitIcon.id = "UnitIcon";
    unitIcon.className = "UnitIconNoSelect";
    constructorUnit.appendChild(unitIcon);

    let equippingPanelII = document.createElement("div");
    equippingPanelII.className = "verticalEquipPanel";
    CreateCells(2, 3, "UnitEquip noActive", "UnitEquip", equippingPanelII, true);
    constructorUnit.appendChild(equippingPanelII);

    let equippingPanelI = document.createElement("div");
    CreateCells(1, 3, "UnitEquip noActive", "UnitEquip", equippingPanelI);
    constructorUnit.appendChild(equippingPanelI)
}

function FillingSquadConstructor(unitData) {
    let unitIcon = document.getElementById("UnitIcon");
    unitIcon.innerHTML = "";

    let mask1 = document.createElement('div');
    mask1.className = 'mask unit inner';
    mask1.id = 'unitMaskBody1';
    mask1.style.background = "#" + unitData.unit.body_color_1.split('x')[1];
    $(mask1).css("-webkit-mask-image", "url(/assets/units/body/" + unitData.unit.body.name + "_mask.png)");

    let mask2 = document.createElement('div');
    mask2.id = 'unitMaskBody2';
    mask2.style.opacity = '0.3';
    mask2.className = 'mask unit inner';
    mask2.style.background = "#" + unitData.unit.body_color_2.split('x')[1];
    $(mask2).css("-webkit-mask-image", "url(/assets/units/body/" + unitData.unit.body.name + "_mask2.png)");

    unitIcon.appendChild(mask2);
    unitIcon.appendChild(mask1);

    unitIcon.style.background = "url(/assets/units/body/" + unitData.unit.body.name + ".png) center center / 80px no-repeat," +
        "url(/assets/units/body/" + unitData.unit.body.name + "_bottom.png) center center / 80px no-repeat, #4c4c4c";

    unitIcon.slotData = JSON.stringify(unitData);
    unitIcon.unitBody = unitData.unit.body;
    unitIcon.onclick = BodyUnitMenu;

    FillParams(unitData.unit);
    CreateColorInputs(unitIcon, unitData.unit, unitData.number_slot, 'unit');

    UpdateWeaponIcon(unitIcon, "weaponUnitInnerIcon", unitData);
    FillPowerPanel(unitData.unit.body, "unitPowerPanel");
    FillCubePanel(unitData.unit.body, "unitCubePanel");
    FillUnitWeaponTypePanel(unitData.unit.body, "weaponTypePanel");

    UpdateCells(1, "UnitEquip", unitData.unit.body.equippingI, "UnitEquip");
    UpdateCells(2, "UnitEquip", unitData.unit.body.equippingII, "UnitEquip");
    UpdateCells(3, "UnitEquip", unitData.unit.body.equippingIII, "UnitEquip");
    UpdateCells(3, "UnitEquip", unitData.unit.body.weapons, "UnitEquip");
}