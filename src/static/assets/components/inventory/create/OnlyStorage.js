function OnlyStorage() {
    let storage = document.createElement("div");
    storage.id = "storage";
    storage.style.position = "absolute";
    storage.style.top = "70px";
    storage.style.right = "15px";
    storage.style.width = "187px";
    document.body.appendChild(storage);

    CreateStorage();
    ConnectMarket();

    $(storage).resizable({
        alsoResize: "#inventoryStorage",
        minHeight: 105,
        minWidth: 187,
        handles: "se",
    });

    let buttons = CreateControlButtons("0", "61px", "-3px", "29px");

    $('#storage .InventoryHead').css("margin", "-3px 0px 3px 0px");

    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'storage')
    });

    $(buttons.close).mousedown(function () {
        storage.remove();
    });

    storage.appendChild(buttons.move);
    storage.appendChild(buttons.hide);
    storage.appendChild(buttons.close);
}