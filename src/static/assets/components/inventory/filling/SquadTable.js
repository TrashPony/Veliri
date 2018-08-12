function SquadTable(shipBody) {
    for (let slot in shipBody.equippingIV) {
        if (shipBody.equippingIV.hasOwnProperty(slot)) {

            let unitSlot = shipBody.equippingIV[slot];
            let cell = document.getElementById("squad " + slot + unitSlot.type_slot);
            cell.className = "inventoryUnit active";

            cell.slotData = JSON.stringify(unitSlot);
            cell.onclick = OpenUnitEditor
        }
    }
}

function OpenUnitEditor() {
    let constructorUnit = document.getElementById("ConstructorUnit");

    if (constructorUnit) {
        constructorUnit.remove();
    }

    constructorUnit = document.createElement("div");
    constructorUnit.id = "ConstructorUnit";
    constructorUnit.style.left = Number(this.getBoundingClientRect().left - 75) + "px";
    constructorUnit.style.top = Number(this.getBoundingClientRect().top - 220) + "px";

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