function OpenBox(inventory, boxID, capacitySize, error) {

    if (game.squad.toBox) {
        game.squad.toBox.to = false
    }

    if (error) {
        if (document.getElementById("boxPass")) {
            document.getElementById("boxPass").remove();
        }
        let func = function () {
            global.send(JSON.stringify({
                event: "openBox",
                box_id: boxID,
                box_password: Number(document.getElementById("passPlaceBox").value),
            }));
        };
        PassBlock(func, "boxPass");
        return;
    }

    let openBox;
    if (document.getElementById("openBox" + boxID)) {
        openBox = document.getElementById("openBox" + boxID);
        $(openBox).empty();
    } else {
        openBox = document.createElement("div");
        openBox.id = "openBox" + boxID;
        openBox.className = "openBox";
        openBox.style.top = "30%";
        openBox.style.left = "calc(50% - 200px)";
    }

    $(openBox).droppable({
        drop: function (event, ui) {
            let draggable = ui.draggable;
            if (draggable.data("slotData").parent === "squadInventory") {
                global.send(JSON.stringify({
                    event: "placeItemToBox",
                    box_id: Number(boxID),
                    slot: Number(draggable.data("slotData").number)
                }))
            }
        }
    });

    let head = document.createElement("span");
    head.innerHTML = "Ящик " + boxID;
    openBox.appendChild(head);

    let buttons = CreateControlButtons("5px", "31px", "-3px", "");
    buttons.close.onclick = function () {
        openBox.remove();
    };
    openBox.appendChild(buttons.close);
    buttons.move.onmousedown = function (event) {
        moveWindow(event, "openBox" + boxID)
    };
    openBox.appendChild(buttons.move);

    let sizeInfo = document.createElement("div");
    sizeInfo.className = "sizeInventoryInfo";
    sizeBox(sizeInfo, capacitySize, inventory);

    openBox.appendChild(sizeInfo);

    let storageCell = document.createElement("div");
    storageCell.className = "storageCell";
    fillInventory(storageCell, inventory, boxID);

    openBox.appendChild(storageCell);
    document.body.appendChild(openBox);
}

function fillInventory(parent, inventory, boxID) {
    for (let i in inventory.slots) {
        if (inventory.slots.hasOwnProperty(i)) {

            let slot = document.createElement("div");
            slot.className = "InventoryCell";
            slot.slotData = JSON.stringify(inventory.slots[i]);
            slot.number = i;

            if (JSON.parse(slot.slotData).type === "resource" || JSON.parse(slot.slotData).type === "recycle") {
                slot.style.backgroundImage = "url(/assets/resource/" + JSON.parse(slot.slotData).item.name + ".png)";
            } else if (JSON.parse(slot.slotData).type === "boxes") {
                slot.style.backgroundImage = "url(/assets/" + JSON.parse(slot.slotData).type + "/" + JSON.parse(slot.slotData).item.name + ".png)";
            } else {
                slot.style.backgroundImage = "url(/assets/units/" + JSON.parse(slot.slotData).type + "/" + JSON.parse(slot.slotData).item.name + ".png)";
            }

            slot.innerHTML = "<span class='QuantityItems'>" + inventory.slots[i].quantity + "</span>";

            slot.onclick = function () {
                global.send(JSON.stringify({
                    event: "getItemFromBox",
                    box_id: Number(boxID),
                    slot: Number(i)
                }))
            };

            $(slot).data("slotData", {parent: "box:" + boxID, data: inventory.slots[i], number: i});
            $(slot).draggable({
                revert: "invalid",
                zIndex: 100,
                helper: 'clone'
            });
            parent.appendChild(slot);
        }
    }
}

function sizeBox(sizeInfo, capacitySize, inventory) {
    let size = 0;

    for (let i in inventory.slots) {
        if (inventory.slots.hasOwnProperty(i)) {
            size += inventory.slots[i].size
        }
    }

    let percentFill = 100 / (capacitySize / size);

    let textColor = "";
    if (size > capacitySize) {
        textColor = "#b9281d"
    } else {
        textColor = "#decbcb"
    }

    sizeInfo.innerHTML = "<div class='realSize' style='width:" + percentFill + "%'>" +
        "<span>" + size.toFixed(1) + " / " + capacitySize + "</span>" +
        "</div>";
    sizeInfo.style.color = textColor;
}