function CreateConstructorMenu() {
    let constructorBackGround = document.getElementById("ConstructorBackGround");

    let powerPanel = document.createElement("div");
    powerPanel.id = "powerPanel";

    let weaponType = document.createElement("div");
    weaponType.id = "MSWeaponPanel";
    constructorBackGround.appendChild(weaponType);

    let spanPower = document.createElement("span");
    spanPower.className = "Value";
    spanPower.innerHTML = "Энергия: Max/Use";
    powerPanel.appendChild(spanPower);
    constructorBackGround.appendChild(powerPanel);

    let constructorMS = document.createElement("div");
    constructorMS.id = "ConstructorMS";
    constructorBackGround.appendChild(constructorMS);

    /* 3 type slots */
    let equippingPanelIII = document.createElement("div");
    CreateCells(3, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelIII);
    constructorMS.appendChild(equippingPanelIII);

    /* 5 type slots */
    let equippingPanelV = document.createElement("div");
    equippingPanelV.className = "verticalEquipPanel";
    CreateCells(5, 2, "inventoryEquipping noActive", "inventoryEquip", equippingPanelV, true);
    constructorMS.appendChild(equippingPanelV);

    /* shipIcon */
    let unitIcon = document.createElement("div");
    unitIcon.id = "MSIcon";
    unitIcon.className = "UnitIconNoSelect";
    $(unitIcon).droppable({
        drop: function (event, ui) {
            $('.ui-selected').removeClass('ui-selected');
            let draggable = ui.draggable;
            let slotData = draggable.data("slotData");

            if (slotData.parent === "squadInventory" && slotData.data.type === "body") {
                inventorySocket.send(JSON.stringify({
                    event: "SetMotherShipBody",
                    id_body: Number(slotData.data.item.id),
                    inventory_slot: Number(slotData.number),
                    unit_slot: Number(slotData.number_slot)
                }));
                DestroyInventoryClickEvent();
                DestroyInventoryTip();
            }
        }
    });
    constructorMS.appendChild(unitIcon);

    /* 2 type slots */
    let equippingPanelII = document.createElement("div");
    equippingPanelII.className = "verticalEquipPanel";
    CreateCells(2, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelII, true);
    constructorMS.appendChild(equippingPanelII);

    /* 1 type slots */
    let equippingPanelI = document.createElement("div");
    CreateCells(1, 5, "inventoryEquipping noActive", "inventoryEquip", equippingPanelI);
    constructorMS.appendChild(equippingPanelI);

    let repairButton = document.createElement("div");
    repairButton.className = "repairButton";
    repairButton.id = "repairButton";
    repairButton.onclick = CreateRepairMenu;
    constructorMS.appendChild(repairButton);
}