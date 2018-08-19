function DestroyInventoryTip() {
    if (document.getElementById("InventoryTip")) {
        document.getElementById("InventoryTip").remove();
    }
}

function DestroyInventoryClickEvent() {
    cellUnitIconDestroySelect();

    cellAmmoDestroySelect();

    cellEquipDestroySelect(1, 5, "inventoryEquip", "inventoryEquipping"); // обнуляем ячейки эквипа мса
    cellEquipDestroySelect(2, 5, "inventoryEquip", "inventoryEquipping");
    cellEquipDestroySelect(3, 5, "inventoryEquip", "inventoryEquipping");
    cellEquipDestroySelect(5, 2, "inventoryEquip", "inventoryEquipping");

    let constructorUnit = document.getElementById("ConstructorUnit");
    if (constructorUnit) {
        cellEquipDestroySelect(1, 3, "UnitEquip", "UnitEquip"); // обнуляем ячейки эквипа юнита
        cellEquipDestroySelect(2, 3, "UnitEquip", "UnitEquip");
        cellEquipDestroySelect(3, 3, "UnitEquip", "UnitEquip");
    }
}

function cellEquipDestroySelect(typeSlot, count, idPrefix, classPrefix) {
    for (let i = 1; i <= count; i++) {
        let equipSlot = document.getElementById(idPrefix + Number(i) + typeSlot);
        if (equipSlot.className === classPrefix + " active select") {
            if (JSON.parse(equipSlot.slotData).hasOwnProperty("weapon")) {

                equipSlot.style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
                equipSlot.style.cursor = "auto";

                equipSlot.onmouseout = function () {
                    this.style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
                    this.style.cursor = "auto";
                };

                equipSlot.className = classPrefix + " active weapon";

                if (JSON.parse(equipSlot.slotData).weapon !== null) {
                    if (classPrefix === "inventoryEquipping") {
                        equipSlot.onclick = WeaponMSMenu;
                    } else {
                        equipSlot.onclick = WeaponUnitMenu;
                    }
                } else {
                    equipSlot.onclick = null;
                }

            } else {

                equipSlot.style.boxShadow = "0 0 0px 0px rgb(0, 0, 0)";
                equipSlot.style.cursor = "auto";

                equipSlot.onmouseout = function () {
                    this.style.boxShadow = "0 0 0px 0px rgb(0, 0, 0)";
                    this.style.cursor = "auto";
                };

                equipSlot.className = classPrefix + " active";

                if (JSON.parse(equipSlot.slotData) !== null) {
                    if (classPrefix === "inventoryEquipping") {
                        equipSlot.onclick = EquipMSMenu;
                    } else {
                        equipSlot.onclick = EquipUnitMenu;
                    }
                } else {
                    equipSlot.onclick = null;
                }
            }
        }
    }
}

function cellAmmoDestroySelect() {
    let ammoCells = document.getElementsByClassName("inventoryAmmoCell"); // обнуляем ячейки боеприпасов
    for (let i = 0; i < ammoCells.length; i++) {
        ammoCells[i].onmouseout = function (event) {
            event.stopPropagation ? event.stopPropagation() : (event.cancelBubble = true);
            this.style.boxShadow = "0 0 5px 3px rgb(200, 200, 0)";
            this.style.cursor = "auto";
        };
        ammoCells[i].style.boxShadow = "0 0 5px 3px rgb(200, 200, 0)";
        ammoCells[i].style.cursor = "auto";

        if (JSON.parse(ammoCells[i].slotData).ammo != null && JSON.parse(ammoCells[i].slotData).ammo !== undefined) {
            if (ammoCells[i].className === "inventoryAmmoCell inventoryEquipping") {
                ammoCells[i].onclick = AmmoMSMenu;
            } else {
                ammoCells[i].onclick = AmmoUnitMenu;
            }
        } else {
            ammoCells[i].onclick = null;
        }
    }
}

function cellUnitIconDestroySelect() {
    let shipIcon = document.getElementById("MSIcon"); // обнуляем икноку мазершипа
    shipIcon.className = "";

    if (shipIcon.shipBody != null && shipIcon.shipBody !== undefined) {
        shipIcon.onclick = BodyMSMenu;
    } else {
        shipIcon.onclick = null;
    }

    let unitIcon = document.getElementById("UnitIcon"); // обнуляем икноку мазершипа
    if (unitIcon) {
        unitIcon.className = "";
        if (unitIcon.unitBody != null && unitIcon.unitBody !== undefined) {
            unitIcon.onclick = BodyUnitMenu;
        } else {
            unitIcon.onclick = null;
        }
    }
}