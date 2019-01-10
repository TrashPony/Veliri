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
                if (draggable.data("selectedItems") !== undefined) {
                    global.send(JSON.stringify({
                        event: "placeItemsToBox",
                        box_id: Number(boxID),
                        slots: draggable.data("selectedItems").slotsNumbers
                    }));
                    $(draggable).removeData("selectedItems");
                } else {
                    global.send(JSON.stringify({
                        event: "placeItemToBox",
                        box_id: Number(boxID),
                        slot: Number(draggable.data("slotData").number)
                    }))
                }
            } else if (draggable.data("slotData").parent.split(':')[0] === "box") {

                let toBoxID = draggable.data("slotData").parent.split(':')[1];

                if (draggable.data("selectedItems") !== undefined) {
                    global.send(JSON.stringify({
                        event: "boxToBoxItems",
                        to_box_id: Number(boxID),
                        slots: draggable.data("selectedItems").slotsNumbers,
                        box_id: Number(toBoxID),
                    }));
                    $(draggable).removeData("selectedItems");
                } else {
                    global.send(JSON.stringify({
                        event: "boxToBoxItem",
                        to_box_id: Number(boxID),
                        slot: Number(draggable.data("slotData").number),
                        box_id: Number(toBoxID),
                    }));
                }
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
    $(storageCell).selectable();

    fillInventory(storageCell, inventory, boxID);

    openBox.appendChild(storageCell);
    document.body.appendChild(openBox);
}

function fillInventory(parent, inventory, boxID) {
    for (let i in inventory.slots) {
        if (inventory.slots.hasOwnProperty(i) && inventory.slots[i].item) {
            let cell = document.createElement("div");
            CreateInventoryCell(cell, inventory.slots[i], i, "box:" + boxID, onclick);

            cell.onclick = function () {
                global.send(JSON.stringify({
                    event: "getItemFromBox",
                    box_id: Number(boxID),
                    slot: Number(i)
                }))
            };
            parent.appendChild(cell);
        }
    }
}

function sizeBox(sizeInfo, capacitySize, inventory) {
    let size = 0;

    for (let i in inventory.slots) {
        if (inventory.slots.hasOwnProperty(i) && inventory.slots[i].item) {
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