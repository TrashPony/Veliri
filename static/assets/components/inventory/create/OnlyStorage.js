function OnlyStorage() {
    let storage = document.createElement("div");
    storage.id = "storage";
    storage.style.width = "157px";
    storage.style.padding = '0';
    storage.style.margin = '4px 0px 0px';

    let inventory = document.createElement("div");
    inventory.id = "Inventory";
    inventory.style.width = "157px";
    inventory.style.padding = '0';
    inventory.style.margin = '0';
    inventory.style.float = 'left';

    let wrapper = document.createElement('div');
    wrapper.id = "wrapperInventoryAndStorage";
    wrapper.appendChild(inventory);
    wrapper.appendChild(storage);
    document.body.appendChild(wrapper);

    CreateInventory();
    CreateStorage();
    ConnectMarket();

    $(wrapper).resizable({
        alsoResize: "#inventoryStorage, #storage",
        minHeight: 248,
        minWidth: 159,
        handles: "se",
    });

    let buttons = CreateControlButtons("1px", "32px", "-3px", "29px");

    document.getElementById('utilButton').remove();
    document.getElementById('destroyButton').remove();
    document.getElementById('inventoryStorage').style.height = '58px';
    document.getElementById('sizeInventoryInfo').style.margin = '8px 0 -4px 1px';


    storage.style.height = '107px';

    $('#storage .InventoryHead').css("margin", "1px 0px 3px");
    $('#Inventory .InventoryHead').css("margin", "1px 0px 3px");

    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'wrapperInventoryAndStorage')
    });

    $(buttons.close).mousedown(function () {
        wrapper.remove();
    });

    inventory.appendChild(buttons.move);
    inventory.appendChild(buttons.close);
}