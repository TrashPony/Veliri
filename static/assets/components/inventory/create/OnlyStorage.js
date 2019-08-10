function OnlyStorage() {
    let storage = document.createElement("div");
    storage.id = "storage";
    storage.style.width = "172px";
    storage.style.padding = '0px 3px 0 0px';
    storage.style.margin = '4px 0px 0px';
    storage.style.float = 'left';
    storage.style.height = '109px';

    let inventory = document.createElement("div");
    inventory.id = "Inventory";
    inventory.style.width = "172px";
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

    $(wrapper).data({
        resize: function (event, ui, el) {
            let storage = $('#storage');
            let inventoryStorage = $('#inventoryStorage');

            let inventory = $('#Inventory');
            let inventoryStorageInventory = $('#inventoryStorageInventory');

            inventory.css("height", el.height() / 2 - 4);
            storage.css("height", el.height() / 2 - 4);

            inventoryStorageInventory.css("height", inventory.height() - 51);
            inventoryStorage.css("height", storage.height() - 51);

            inventory.css("width", el.width() - 5);
            storage.css("width", el.width() - 5);
        }
    });

    $(wrapper).resizable({
        minHeight: 225,
        minWidth: 177,
        handles: "se",
        resize: function (e, ui) {
            $(this).data("resize")(event, ui, $(this))
        },
        stop: function (e, ui) {
            setState(this.id, $(this).position().left, $(this).position().top, $(this).height(), $(this).width(), true);
        }
    });

    document.getElementById('utilButton').remove();
    document.getElementById('destroyButton').remove();
    document.getElementById('inventoryStorage').style.height = '58px';
    document.getElementById('inventoryStorageInventory').style.height = '58px';

    $('#storage .InventoryHead').css("margin", "1px 0px 3px");
    $('#Inventory .InventoryHead').css("margin", "1px 0px 3px");
    $('#Inventory .sortPanel').css("display", "none");


    let buttons = CreateControlButtons("1px", "32px", "-3px", "29px");
    $(buttons.move).mousedown(function (event) {
        moveWindow(event, 'wrapperInventoryAndStorage')
    });

    $(buttons.close).mousedown(function () {
        setState(wrapper.id, $(wrapper).position().left, $(wrapper).position().top, $(wrapper).height(), $(wrapper).width(), false);
    });

    inventory.appendChild(buttons.move);
    inventory.appendChild(buttons.close);

    openWindow(wrapper.id, wrapper);
}