function FillRecycler(jsonData) {
    let itemsPool = document.getElementById("itemsPool");
    let previewPool = document.getElementById("previewPool");
    if (!itemsPool) return;

    $("#itemsPool .InventoryCell").remove();
    $("#itemsPool .RecycleSection").remove();
    $("#previewPool .InventoryCell").remove();
    $("#previewPool .RecycleSection").remove();

    for (let i in jsonData.recycle_slots) {

        let cell = document.createElement("div");
        CreateInventoryCell(cell, jsonData.recycle_slots[i].slot, i, "recycler", onclick);

        $(cell).draggable({
            revert: false,
            stop: function (event, ui) {
                let elem = document.elementFromPoint(ui.position.left, ui.position.top);
                if (!$(elem).hasClass("itemsPools")) {
                    if ($(this).data("slotData")) {
                        if ($(this).data("selectedItems") !== undefined) {
                            lobby.send(JSON.stringify({
                                event: "RemoveItemsFromProcessor",
                                storage_slots: $(this).data("selectedItems").slotsNumbers,
                            }));
                        } else {
                            lobby.send(JSON.stringify({
                                event: "RemoveItemFromProcessor",
                                recycler_slot: Number($(this).data("slotData").number),
                            }));
                        }
                    }
                }
            },
        });

        if (!jsonData.recycle_slots[i].recycled) {
            cell.style.border = "1px solid red";
        }

        cell.onclick = function () {
            lobby.send(JSON.stringify({
                event: "RemoveItemFromProcessor",
                recycler_slot: Number(i),
            }));
        };

        let section = CheckRecycleSection(jsonData.recycle_slots[i].slot, itemsPool);
        section.appendChild(cell);
    }

    for (let i in jsonData.preview_recycle_slots) {

        let cell = document.createElement("div");
        CreateInventoryCell(cell, jsonData.preview_recycle_slots[i], i, "", onclick);

        let section = CheckRecycleSection(jsonData.preview_recycle_slots[i], previewPool);
        section.appendChild(cell);
    }
}

function CheckRecycleSection(slotData, parent) {
    if (document.getElementById("RecycleSection" + slotData.type + parent.id)) {
        return document.getElementById("RecycleSection" + slotData.type + parent.id);
    } else {
        let newSection = document.createElement("div");
        newSection.className = "RecycleSection";
        newSection.id = "RecycleSection" + slotData.type + parent.id;

        let nameSection = document.createElement("div");
        nameSection.className = "nameSection";
        nameSection.innerHTML = slotData.type;

        let hideArrow = document.createElement("div");
        hideArrow.className = "hideArrowSection";
        $(hideArrow).click(function () {
            // TODO
        });

        nameSection.appendChild(hideArrow);
        newSection.appendChild(nameSection);
        parent.appendChild(newSection);
        return newSection;
    }
}