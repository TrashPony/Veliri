function SquadTable(squad) {
    for (let slot = 1; slot <= 6; slot++) {

        let cell = document.getElementById("squad " + slot + 4); // 4 это тип ячейки

        if (squad.mather_ship.units.hasOwnProperty(slot)) {

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

function NoActiveUnitCell(slotData) {
    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) {
        if (JSON.parse(constructorUnit.slotData).number_slot === slotData.number_slot) {
            let cells = document.getElementsByClassName("UnitEquip");
            for (let i = 0; i < cells.length; i++) {
                cells[i].ammoCell = null;

                cells[i].className = "UnitEquip noActive";
                cells[i].style.backgroundImage = "";
                cells[i].style.boxShadow = "0 0 0 0 rgb(0, 0, 0)";

                cells[i].onmouseout = null;
                cells[i].onmouseover = null;
                cells[i].onclick = null;

                for (let child in cells[i].childNodes) {
                    if (cells[i].childNodes.hasOwnProperty(child)) {
                        cells[i].childNodes[child].remove();
                    }
                }
            }

            let unitIcon = document.getElementById("UnitIcon");
            unitIcon.style.backgroundImage = null;
            unitIcon.onclick = null;
            unitIcon.shipBody = null;
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
    unitIcon.onclick = BodyUnitRemove;

    UpdateCells(1, "UnitEquip", slotData.unit.body.equippingI, "UnitEquip");
    UpdateCells(2, "UnitEquip", slotData.unit.body.equippingII, "UnitEquip");
    UpdateCells(3, "UnitEquip", slotData.unit.body.equippingIII, "UnitEquip");
    UpdateCells(3, "UnitEquip", slotData.unit.body.weapons, "UnitEquip");
}