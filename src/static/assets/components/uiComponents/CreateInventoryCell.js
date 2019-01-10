function CreateInventoryCell(cell, slotData, slotNumber, parent) {
    cell.className = "InventoryCell active";
    cell.slotData = JSON.stringify(slotData);
    cell.number = slotNumber;

    if (slotData.type === "resource" || slotData.type === "recycle") {
        cell.style.backgroundImage = "url(/assets/resource/" + slotData.item.name + ".png)";
    } else if (slotData.type === "boxes") {
        cell.style.backgroundImage = "url(/assets/" + slotData.type + "/" +slotData.item.name + ".png)";
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
            if (selectItems.length > 0) {
                // если выделено много элементов то отправляем их все
                // helper это иконка которая улетает с мышкой
                let helper = $('.InventoryCell.ui-draggable.ui-draggable-handle.ui-draggable-dragging');
                helper.empty();
                helper.css('background-image', 'url(/assets/components/inventory/img/dragDetail.png');

                let slotsNumbers = [];
                slotsNumbers.push($(cell).data("slotData").number);
                selectItems.each(function (index) {
                    if ($(this).data("slotData") !== undefined && $(this).data("slotData").number !== $(cell).data("slotData").number) {
                        slotsNumbers.push($(this).data("slotData").number);
                    }
                });

                $(cell).data("selectedItems", {parent: parent, slotsNumbers: slotsNumbers});
            } else {
                $(cell).removeData("selectedItems");
            }
        },
        revert: "invalid",
        zIndex: 999,
        helper: 'clone'
    });
}