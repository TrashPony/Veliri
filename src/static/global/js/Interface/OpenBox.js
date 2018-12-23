function OpenBox(inventory, boxID) {
    console.log(inventory, boxID);

    if (game.squad.toBox) {
        game.squad.toBox.to = false
    }

    if (document.getElementById("openBox" + boxID)) {
        document.getElementById("openBox" + boxID).remove();
    }

    let openBox = document.createElement("div");
    openBox.id = "openBox" + boxID;
    openBox.className = "openBox";
    openBox.style.top = "30%";
    openBox.style.left = "calc(50% - 200px)";

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

    let storageCell = document.createElement("div");
    storageCell.className = "storageCell";
    fillInventory(storageCell, inventory);

    openBox.appendChild(storageCell);
    document.body.appendChild(openBox);
}

function fillInventory(parent, inventory) {
    for (let i in inventory.slots) {
        if (inventory.slots.hasOwnProperty(i)) {
            let slot = document.createElement("div");
            slot.className = "InventoryCell";
            slot.number = i;
            slot.style.backgroundImage = "url(/assets/units/" + inventory.slots[i].type + "/" + inventory.slots[i].item.name + ".png)";
            slot.innerHTML = "<span class='QuantityItems'>" + inventory.slots[i].quantity + "</span>";

            parent.appendChild(slot);
        }
    }
}