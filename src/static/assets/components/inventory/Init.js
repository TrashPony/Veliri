function InitInventoryMenu(closeFunc) {
    /* осталь надежду вся сюда входящий, это дерьмо уже не отрефачить */
    let promise = new Promise((resolve) => {
        includeJS("../assets/components/inventory/webSocket.js");
        includeJS("../assets/components/inventory/Create.js");

        includeJS("../assets/components/inventory/selected/SelectInventoryItem.js");
        includeJS("../assets/components/inventory/selected/DeselectInventoryItem.js");
        includeJS("../assets/components/inventory/selected/SelectEquip.js");

        includeJS("../assets/components/inventory/filling/Filling.js");
        includeJS("../assets/components/inventory/filling/ConstructorTable.js");
        includeJS("../assets/components/inventory/filling/InventoryTable.js");
        includeJS("../assets/components/inventory/filling/SquadTable.js");
        includeJS("../assets/components/inventory/filling/MarkConstructorsCell.js");
        includeJS("../assets/components/inventory/filling/DeactivateCell.js");
        includeJS("../assets/components/inventory/filling/PowerPanel.js");
        includeJS("../assets/components/inventory/filling/HealBar.js");

        includeJS("../assets/components/inventory/tip/ClickTip.js");
        includeJS("../assets/components/inventory/tip/SelectItem.js");

        includeJS("../assets/components/inventory/set/SetAmmo.js");
        includeJS("../assets/components/inventory/set/SetBody.js");
        includeJS("../assets/components/inventory/set/SetEquip.js");
        includeJS("../assets/components/inventory/set/SetWeapon.js");

        includeJS("../assets/components/inventory/repair/CreateRepairMenu.js");

        includeJS("../assets/components/inventory/inventory/RecycleItems.js");
        includeJS("../assets/components/inventory/inventory/SortingItems.js");
        includeJS("../assets/components/inventory/inventory/ThrowItems.js");
        includeJS("../assets/components/inventory/inventory/checkConfirmMenu.js");

        includeCSS("../assets/components/inventory/css/constructorMS.css");
        includeCSS("../assets/components/inventory/css/constructorUnit.css");
        includeCSS("../assets/components/inventory/css/equipMSBox.css");
        includeCSS("../assets/components/inventory/css/equipUnitBox.css");
        includeCSS("../assets/components/inventory/css/inventoryCells.css");
        includeCSS("../assets/components/inventory/css/tip.css");
        includeCSS("../assets/components/inventory/css/repair.css");
        includeCSS("../assets/components/inventory/css/inventory.css");


        return resolve();
    });
    //todo чето я хз, промис не работает
    promise.then(
        () => {
            setTimeout(function () {
                ConnectInventory();
                CreateInventoryMenu(closeFunc);
            }, 400);
        }
    );
}

function includeJS(url) {
    let script = document.createElement('script');
    script.type = "text/javascript";
    script.src = url;
    document.getElementsByTagName('head')[0].appendChild(script);
}

function includeCSS(url) {
    let css = document.createElement('link');
    css.type = "text/css";
    css.rel = "stylesheet";
    css.href = url;
    document.getElementsByTagName('head')[0].appendChild(css);
}