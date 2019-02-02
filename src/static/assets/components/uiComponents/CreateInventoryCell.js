function CreateInventoryCell(cell, slotData, slotNumber, parent) {
    cell.className = "InventoryCell active";
    cell.slotData = JSON.stringify(slotData);
    cell.number = slotNumber;

    if (slotData.type === "resource" || slotData.type === "recycle") {
        cell.style.backgroundImage = "url(/assets/resource/" + slotData.item.name + ".png)";
    } else if (slotData.type === "boxes") {
        cell.style.backgroundImage = "url(/assets/" + slotData.type + "/" + slotData.item.name + ".png)";
    } else if (slotData.type === "detail") {
        cell.style.backgroundImage = "url(/assets/resource/detail/" + slotData.item.name + ".png)";
    } else if (slotData.type === "blueprints") {
        cell.style.backgroundImage = "url(/assets/blueprints/" + slotData.item.name + ".png)";
    } else {
        cell.style.backgroundImage = "url(/assets/units/" + slotData.type + "/" + slotData.item.name + ".png)";
    }

    cell.innerHTML = "<span class='QuantityItems'>" + slotData.quantity + "</span>";

    CreateHealBar(cell, "inventory", true);

    $(cell).data("slotData", {parent: parent, data: slotData, number: slotNumber});
    $(cell).draggable({
        disabled: false,
        start: function () {
            let selectItems = $('.InventoryCell.ui-selected');
            if (selectItems.length > 1) {
                // если выделено много элементов то отправляем их все
                let slotsNumbers = [];
                slotsNumbers.push(Number($(cell).data("slotData").number));
                selectItems.each(function (index) {
                    if ($(this).data("slotData") !== undefined && $(this).data("slotData").number !== $(cell).data("slotData").number) {
                        slotsNumbers.push(Number($(this).data("slotData").number));
                    }
                });

                $(cell).data("selectedItems", {parent: parent, slotsNumbers: slotsNumbers});
            } else {
                $(cell).removeData("selectedItems");
                MarkConstructorEquip(cell);
            }
        },
        stop: function () {
            unMarkConstructorEquip();
        },
        drag: function (event, ui) {
            // .ui-draggable-dragging это иконка которая улетает с мышкой
            if ($('.InventoryCell.ui-selected').length > 1) {
                $('.ui-draggable-dragging')
                    .empty()
                    .css('background-image', 'url(/assets/components/inventory/img/dragDetail.png');
            }
        },
        revert: "invalid",
        zIndex: 999,
        helper: 'clone',
        appendTo: "body",
    });
}

function unMarkConstructorEquip() {
    if (document.getElementById("ConstructorMS")) {
        DestroyInventoryClickEvent();
        DestroyInventoryTip();
    }
}

function MarkConstructorEquip(cell) {
    if (document.getElementById("ConstructorMS")) {
        let slotData = $(cell).data("slotData");
        if (slotData.data.item) {
            if (slotData.data.type === "weapon") {
                WeaponSlotMark("inventoryEquip", "inventoryEquipping", 5, null);
                WeaponSlotMark("UnitEquip", "UnitEquip", 3, null);
            } else if (slotData.data.type === "ammo") {
                let ammoCells = document.getElementsByClassName("inventoryAmmoCell");
                for (let i = 0; ammoCells && i < ammoCells.length; i++) {
                    ammoCells[i].style.boxShadow = "0 0 5px 3px rgb(255, 149, 32)";
                    ammoCells[i].style.cursor = "pointer";
                    ammoCells[i].onmouseout = null;
                }
            } else if (slotData.data.type === "equip") {
                EquipSlotMark("inventoryEquip", "inventoryEquipping", slotData.data.item.type_slot, 5, null);
                EquipSlotMark("UnitEquip", "UnitEquip", slotData.data.item.type_slot, 3, null);
            } else if (slotData.data.type === "body") {
                if (slotData.data.item.mother_ship) {
                    document.getElementById("MSIcon").className = "UnitIconSelect";
                } else {
                    if (document.getElementById("UnitIcon")) {
                        document.getElementById("UnitIcon").className = "UnitIconSelect";
                    }
                }
            } else if (slotData.data.type === "recycle" && slotData.data.item.name === "enriched_thorium") {
                SetThorium(null, slotData.number, slotData.parent)
            }
        }
    }
}

function CreateHealBar(cell, type, append) {
    let cellData = JSON.parse(cell.slotData);

    if (cellData.type !== "ammo" && cellData.type !== "resource" && cellData.type !== "recycle"
        && cellData.type !== "boxes" && cellData.type !== "detail" && cellData.type !== "blueprint" &&
        cellData.type !== "blueprints") {

        let backHealBar = document.createElement("div");

        let percentHP = 0;

        if (type === "inventory") {
            backHealBar.className = "backInventoryHealBar";
            percentHP = 100 / (cellData.item.max_hp / cellData.hp);
        } else if (type === "equip") {
            backHealBar.className = "backEquipHealBar";
            percentHP = 100 / (cellData.equip.max_hp / cellData.hp);
        } else if (type === "weapon") {
            backHealBar.className = "backWeaponHealBar";
            percentHP = 100 / (cellData.weapon.max_hp / cellData.hp);
        } else if (type === "body") {
            backHealBar.className = "backBodyHealBar";
            percentHP = 100 / (cellData.body.max_hp / cellData.hp);
        }

        let healBar = document.createElement("div");
        healBar.className = "healBar";

        healBar.style.width = percentHP + "%";

        if (percentHP === 100) {
            backHealBar.style.opacity = "0"
        } else if (percentHP < 90 && percentHP > 75) {
            healBar.style.backgroundColor = "#fff326"
        } else if (percentHP < 75 && percentHP > 50) {
            healBar.style.backgroundColor = "#fac227"
        } else if (percentHP < 50 && percentHP > 25) {
            healBar.style.backgroundColor = "#fa7b31"
        } else if (percentHP < 25 && cellData.hp > 1) {
            healBar.style.backgroundColor = "#ff2615"
        } else if (cellData.hp === 0) {
            backHealBar.style.opacity = "0";
            // todo показывать что предмет сломан например box-shadow insert red
        }

        if (append) {
            backHealBar.appendChild(healBar);
            cell.appendChild(backHealBar);
        }

        return percentHP;
    }
}