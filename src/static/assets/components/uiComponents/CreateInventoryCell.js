function CreateInventoryCell(cell, slotData, slotNumber, parent) {
    cell.className = "InventoryCell active";
    cell.slotData = JSON.stringify(slotData);
    cell.number = slotNumber;

    if (slotData.type === "resource" || slotData.type === "recycle") {
        cell.style.backgroundImage = "url(/assets/resource/" + slotData.item.name + ".png)";
    } else if (slotData.type === "boxes") {
        cell.style.backgroundImage = "url(/assets/" + slotData.type + "/" + slotData.item.name + ".png)";
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
            }
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
        helper: 'clone'
    });
}

function CreateHealBar(cell, type, append) {
    let cellData = JSON.parse(cell.slotData);

    if (cellData.type !== "ammo" && cellData.type !== "resource" && cellData.type !== "recycle" && cellData.type !== "boxes") {
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