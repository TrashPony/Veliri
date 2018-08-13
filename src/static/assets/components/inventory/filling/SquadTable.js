function SquadTable(squad) {

    console.log(squad);

    for (let slot in squad.mather_ship.units) {
        if (squad.mather_ship.units.hasOwnProperty(slot)) {

            let unitSlot = squad.mather_ship.units[slot];

            let cell = document.getElementById("squad " + slot + 4); // 4 это тип ячейки

            if (cell.className !== "inventoryUnit select") {
                cell.className = "inventoryUnit active";
            }

            cell.slotData = JSON.stringify(unitSlot);
            cell.onclick = OpenUnitEditor
        }
    }
}

function OpenUnitEditor() {
    let constructorUnit = document.getElementById("ConstructorUnit");

    let inventoryUnits = document.getElementsByClassName("inventoryUnit select");
    for (let slot in inventoryUnits) {
        inventoryUnits[slot].className = "inventoryUnit active";
    }

    if (constructorUnit) {
        if (JSON.parse(constructorUnit.slotData).number_slot === JSON.parse(this.slotData).number_slot) {
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
}

function CreateUnitEquipSlots(constructorUnit) {
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