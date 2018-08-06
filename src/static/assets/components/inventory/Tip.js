function InventoryTip(item, x, y) {


    let tip = document.createElement("div");
    tip.style.top = y + "px";
    tip.style.left = x + "px";
    tip.id = "InventoryTip";

    let name = document.createElement("span");
    name.className = "InventoryTipName";
    name.innerHTML = item.name;
    tip.appendChild(name);

    let cancelButton = document.createElement("input");
    cancelButton.type = "button";
    cancelButton.className = "lobbyButton inventoryTip";
    cancelButton.value = "Отменить";
    cancelButton.style.pointerEvents = "auto";

    cancelButton.onclick = function () {
        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    };

    tip.appendChild(cancelButton);

    document.body.appendChild(tip);
}

function DestroyInventoryTip() {
    if (document.getElementById("InventoryTip")) {
        document.getElementById("InventoryTip").remove();
    }
}

function DestroyInventoryClickEvent() {
    cellUnitIconDestroySelect();

    cellAmmoDestroySelect();

    cellEquipDestroySelect(1, 5, "inventoryEquip"); // обнуляем ячейки эквипа
    cellEquipDestroySelect(2, 5, "inventoryEquip");
    cellEquipDestroySelect(3, 5, "inventoryEquip");
    cellEquipDestroySelect(5, 2, "inventoryEquip");
}

function cellEquipDestroySelect(typeSlot, count, idPrefix) {
    for (let i = 1; i <= count; i++) {
        let equipSlot = document.getElementById(idPrefix + Number(i) + typeSlot);
        if (equipSlot.className === "inventoryEquipping active select") {
            if (equipSlot.slot.hasOwnProperty("weapon")) {

                equipSlot.style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
                equipSlot.style.cursor = "auto";

                equipSlot.onmouseout = function () {
                    this.style.boxShadow = "0 0 5px 3px rgb(255, 0, 0)";
                    this.style.cursor = "auto";
                };

                equipSlot.className = "inventoryEquipping active weapon";

                if (equipSlot.slot.weapon !== null) {
                    equipSlot.onclick = WeaponRemove;
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

                equipSlot.className = "inventoryEquipping active";

                if (equipSlot.slot.weapon !== null) {
                    equipSlot.onclick = EquipRemove;
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

        if (ammoCells[i].slot.ammo != null && ammoCells[i].slot.ammo !== undefined) {
            ammoCells[i].onclick = AmmoRemove;
        } else {
            ammoCells[i].onclick = null;
        }
    }
}

function cellUnitIconDestroySelect() {
    let shipIcon = document.getElementById("UnitIcon"); // обнуляем икноку мазершипа
    shipIcon.className = "";

    if (shipIcon.shipBody != null && shipIcon.shipBody !== undefined) {
        shipIcon.onclick = RemoveTip;
    } else {
        shipIcon.onclick = null;
    }
}