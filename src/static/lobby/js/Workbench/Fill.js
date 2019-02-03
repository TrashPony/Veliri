function FillWorkbench(jsonData) {
    let bpBlock = document.getElementById("bluePrints");
    for (let i in jsonData.storage.slots) {
        if (jsonData.storage.slots[i].type === "blueprints") {
            let blueRow = document.createElement("div");
            blueRow.className = "blueRow";
            blueRow.innerHTML = "" +
                "<div class='nameBP'>" + jsonData.storage.slots[i].item.name + "</div>" +
                "<div class='countBP'>x" + jsonData.storage.slots[i].quantity + "</div>";
            bpBlock.appendChild(blueRow);

            $(blueRow).click(function () {
                lobby.send(JSON.stringify({
                    event: "SelectBP",
                    storage_slot: Number(i),
                    count: 1,
                }));
            });
        }
    }
}

function SelectBP(jsonData) {
    document.getElementById("bpName").innerHTML = jsonData.blue_print.name;
    document.getElementById("bpIcon").style.backgroundImage = "url(/assets/blueprints/" + jsonData.blue_print.name + ".png)";
    document.getElementById("bpCraftTime").innerHTML = jsonData.blue_print.craft_time + "s";

    let itemPreview = document.getElementById("itemPreview");

    if (jsonData.blue_print.item_type === "resource" || jsonData.blue_print.item_type === "recycle") {
        itemPreview.style.backgroundImage = "url(/assets/resource/" + jsonData.bp_item.name + ".png)";
    } else if (jsonData.blue_print.item_type === "boxes") {
        itemPreview.style.backgroundImage = "url(/assets/" + jsonData.blue_print.item_type + "/" + jsonData.bp_item.name + ".png)";
    } else if (jsonData.blue_print.item_type === "detail") {
        itemPreview.style.backgroundImage = "url(/assets/resource/detail/" + jsonData.bp_item.name + ".png)";
    } else if (jsonData.blue_print.item_type === "blueprints") {
        itemPreview.style.backgroundImage = "url(/assets/blueprints/" + jsonData.bp_item.name + ".png)";
    } else {
        itemPreview.style.backgroundImage = "url(/assets/units/" + jsonData.blue_print.item_type + "/" + jsonData.bp_item.name + ".png)";
    }
    itemPreview.innerHTML = "<span>x" + jsonData.blue_print.count + "</span>";

    fillNeedItems(jsonData.preview_recycle_slots);

    let bpCountWork = document.getElementById("bpCountWork");
    bpCountWork.value = jsonData.count;
    bpCountWork.max = jsonData.max_count;
    bpCountWork.oninput = function () {
        lobby.send(JSON.stringify({
            event: "SelectBP",
            storage_slot: Number(jsonData.storage_slot),
            count: Number(this.value)
        }));
    }
}

function fillNeedItems(items) {
    let needItems = document.getElementById("needItems");
    $(needItems).empty();

    for (let i in items) {
        let cell = document.createElement("div");

        let item = Object.assign({}, items[i]);

        CreateInventoryCell(cell, item, i, "", onclick);
        let section = CheckRecycleSection(item, needItems);

        if (!item.find) {
            cell.style.border = "1px solid red";
            cell.innerHTML += "<div class='noAllowCell'></div>"
        }

        section.appendChild(cell);
    }
}