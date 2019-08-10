function FillRecycler(jsonData) {
    let itemsPool = document.getElementById("itemsPool");
    let previewPool = document.getElementById("previewPool");

    if (!itemsPool) return;

    document.getElementById('UserRecyclePercent').innerHTML = 'Потери: ' + jsonData.user_recycle_skill + '%';
    document.getElementById('fillBackPercent').style.width = jsonData.user_recycle_skill + '%';

    $("#itemsPool .InventoryCell").remove();
    $("#itemsPool .RecycleSection").remove();
    $("#previewPool .InventoryCell").remove();
    $("#previewPool .RecycleSection").remove();

    for (let source in jsonData.recycle_slots) {
        for (let i in jsonData.recycle_slots[source]) {

            let itemSlot = jsonData.recycle_slots[source][i];

            let cell = document.createElement("div");
            CreateInventoryCell(cell, itemSlot.slot, i, "recycler", onclick);

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

            if (!itemSlot.recycled) {
                cell.style.border = "1px solid red";
                cell.innerHTML += "<div class='noAllowCell'></div>"
            }

            cell.onclick = function () {
                lobby.send(JSON.stringify({
                    event: "RemoveItemFromProcessor",
                    recycler_slot: Number(i),
                    item_source: itemSlot.source,
                }));
            };

            let tax = document.createElement('div');
            tax.className = 'itemTax';
            tax.innerHTML = `
            <span style="float: left">Налог:</span><br>
            <span style="float: right">${itemSlot.tax_percent}%</span>
        `;

            if (itemSlot.tax_percent > 0) cell.appendChild(tax);

            let section = CheckRecycleSection(itemSlot.slot, itemsPool);
            section.appendChild(cell);
        }
    }

    for (let i in jsonData.preview_recycle_slots) {

        if (jsonData.preview_recycle_slots[i].quantity === 0) continue;

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