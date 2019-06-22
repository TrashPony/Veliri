function OnlyInventory() {
    let inventory = document.createElement("div");
    inventory.id = "Inventory";
    inventory.style.position = "absolute";
    inventory.style.bottom = "70px";
    inventory.style.right = "15px";
    document.body.appendChild(inventory);

    CreateInventory();

    document.getElementById("utilButton").remove();
    document.getElementById("destroyButton").style.left = "78px";
    document.getElementById("InventoryHead").style.margin = "2px 0px 3px 6px";
    inventory.style.marginLeft = "0";

    let buttons = CreateControlButtons("0", "23px", "-3px", "-3px");
    buttons.close.style.width = "22px";
    buttons.close.onclick = function () {
        setState(inventory.id, $(inventory).position().left, $(inventory).position().top, $(inventory).height(), $(inventory).width(), false);
    };

    buttons.move.style.width = "22px";
    buttons.move.onmousedown = function (event) {
        moveWindow(event, 'Inventory')
    };

    $(inventory).resizable({
        minHeight: 133,
        minWidth: 163,
        handles: "se",
        maxHeight: 400,
        maxWidth: 600,
        stop: function (e, ui) {
            setState(this.id, $(this).position().left, $(this).position().top, $(this).height(), $(this).width(), true);
        }
    });

    inventory.appendChild(buttons.close);
    inventory.appendChild(buttons.move);
    inventory.style.zIndex = "11";

    openWindow(inventory.id, inventory);
}