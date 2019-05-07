function OnlyStorage() {
    let storage = document.createElement("div");
    storage.id = "storage";
    storage.style.width = "157px";
    storage.style.padding = '0px 3px 0 0px';
    storage.style.margin = '4px 0px 0px';
    storage.style.float = 'left';
    storage.style.height = '107px';

    let inventory = document.createElement("div");
    inventory.id = "Inventory";
    inventory.style.width = "157px";
    inventory.style.padding = '0px 3px 0 0px';
    inventory.style.margin = '0';
    inventory.style.float = 'left';
    inventory.style.height = '110px';

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
        minHeight: 225,
        minWidth: 162,
        handles: "se",
        resize() {
            let storage = $('#storage');
            let inventoryStorage = $('#inventoryStorage');
            let inventory = $('#Inventory');
            let inventoryStorageInventory = $('#inventoryStorageInventory');

            if (storage.height() <= 51) {
                inventory.css("height", $(this).height() - 58);
                inventoryStorageInventory.css("height", $(this).height() - 110);
                storage.css("height", 50);
                inventoryStorage.css("height", 0);
            }

            inventory.css("width", $(this).width() - 5);
        }
    });


    document.getElementById('utilButton').remove();
    document.getElementById('destroyButton').remove();
    document.getElementById('inventoryStorage').style.height = '58px';
    document.getElementById('inventoryStorageInventory').style.height = '58px';
    document.getElementById('sizeInventoryInfo').style.margin = '8px 0 -4px 1px';

    $('#storage .InventoryHead').css("margin", "1px 0px 3px");
    $('#Inventory .InventoryHead').css("margin", "1px 0px 3px");
    $('#Inventory .sortPanel').css("display", "none");


    let buttons = CreateControlButtons("1px", "32px", "-3px", "29px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'wrapperInventoryAndStorage')
    });

    $(buttons.close).mousedown(function () {
        wrapper.remove();
    });

    inventory.appendChild(buttons.move);
    inventory.appendChild(buttons.close);
}