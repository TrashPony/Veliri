function SquadTable(squad) {
    for (let slot = 1; slot <= 6; slot++) {

        let cell = document.getElementById("squad " + slot + 4); // 4 это тип ячейки

        if (squad.mather_ship != null && squad.mather_ship.units && squad.mather_ship.units.hasOwnProperty(slot)) {

            let unitSlot = squad.mather_ship.units[slot];

            if (cell.className !== "inventoryUnit select") {
                cell.className = "inventoryUnit active";
            }

            if (squad.mather_ship.units[slot].unit !== null && squad.mather_ship.units[slot].unit !== undefined) {
                cell.style.backgroundImage = "url(/assets/" + squad.mather_ship.units[slot].unit.body.name + ".png)";
                let constructorUnit = document.getElementById("ConstructorUnit");
                if (constructorUnit && JSON.parse(constructorUnit.slotData).number_slot === slot) {
                    FillingSquadConstructor(unitSlot);
                    let unitIcon = document.getElementById("UnitIcon");
                    unitIcon.className = null;
                }
            } else {
                NoActiveUnitCell(unitSlot);
                cell.style.backgroundImage = null;
            }

            cell.slotData = JSON.stringify(unitSlot);
            cell.onclick = OpenUnitEditor
        } else {
            cell.slotData = null;
            cell.onclick = null;
            cell.style.backgroundImage = null;
            cell.className = "inventoryUnit noActive";

            let constructorUnit = document.getElementById("ConstructorUnit");
            if (constructorUnit && JSON.parse(constructorUnit.slotData).number_slot === slot) {
                constructorUnit.remove();
            }
        }
    }
}

function OpenUnitEditor() {
    let constructorUnit = document.getElementById("ConstructorUnit");
    let slotData = JSON.parse(this.slotData);

    let inventoryUnits = document.getElementsByClassName("inventoryUnit select");
    for (let slot in inventoryUnits) {
        inventoryUnits[slot].className = "inventoryUnit active";
    }

    if (constructorUnit) {
        if (JSON.parse(constructorUnit.slotData).number_slot === slotData.number_slot) {
            constructorUnit.remove();
            return;
        } else {
            constructorUnit.remove();
        }
    }

    this.className = "inventoryUnit select";

    constructorUnit = document.createElement("div");
    constructorUnit.id = "ConstructorUnit";
    constructorUnit.style.left = Number(this.getBoundingClientRect().left - 75) + "px";
    constructorUnit.style.top = Number(this.getBoundingClientRect().top - 220) + "px";

    constructorUnit.slotData = this.slotData;

    CreateUnitEquipSlots(constructorUnit);
    document.body.appendChild(constructorUnit);

    if (slotData.unit !== null && slotData.unit !== undefined) {
        FillingSquadConstructor(slotData);
    }
}

function CreateUnitEquipSlots(constructorUnit) {
    let unitPowerPanel = document.createElement("div");
    unitPowerPanel.id = "unitPowerPanel";
    constructorUnit.appendChild(unitPowerPanel);

    let equippingPanelIII = document.createElement("div");
    CreateCells(3, 3, "UnitEquip noActive", "UnitEquip", equippingPanelIII);
    constructorUnit.appendChild(equippingPanelIII);

    let unitIcon = document.createElement("div");
    unitIcon.id = "UnitIcon";
    constructorUnit.appendChild(unitIcon);

    let equippingPanelII = document.createElement("div");
    equippingPanelII.className = "verticalEquipPanel";
    CreateCells(2, 3, "UnitEquip noActive", "UnitEquip", equippingPanelII, true);
    constructorUnit.appendChild(equippingPanelII);

    let equippingPanelI = document.createElement("div");
    CreateCells(1, 3, "UnitEquip noActive", "UnitEquip", equippingPanelI);
    constructorUnit.appendChild(equippingPanelI)
}

function FillingSquadConstructor(slotData) {
    let unitIcon = document.getElementById("UnitIcon");
    unitIcon.style.backgroundImage = "url(/assets/" + slotData.unit.body.name + ".png)";
    unitIcon.slotData = JSON.stringify(slotData);
    unitIcon.unitBody = slotData.unit.body;
    unitIcon.onclick = BodyUnitMenu;

    FillPowerPanel(slotData.unit.body, "unitPowerPanel");

    UpdateCells(1, "UnitEquip", slotData.unit.body.equippingI, "UnitEquip");
    UpdateCells(2, "UnitEquip", slotData.unit.body.equippingII, "UnitEquip");
    UpdateCells(3, "UnitEquip", slotData.unit.body.equippingIII, "UnitEquip");
    UpdateCells(3, "UnitEquip", slotData.unit.body.weapons, "UnitEquip");
}