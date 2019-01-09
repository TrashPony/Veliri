let initInventory = false;
function InitInventoryMenu(closeFunc, option) {
    /* осталь надежду вся сюда входящий, это дерьмо уже не отрефачить */
    let promise = new Promise((resolve) => {
        if (!initInventory) {
            if (typeof webSocketInit === 'undefined' || webSocketInit === null) {
                includeJS("../assets/components/servicesWebSockets.js");
            }
            includeJS("../assets/components/inventory/create/Create.js");
            includeJS("../assets/components/inventory/create/Constructor.js");
            includeJS("../assets/components/inventory/create/Inventory.js");
            includeJS("../assets/components/inventory/create/MotherShipParams.js");
            includeJS("../assets/components/inventory/create/Squad.js");
            includeJS("../assets/components/inventory/create/UserStatus.js");
            includeJS("../assets/components/inventory/create/Storage.js");
            includeJS("../assets/components/inventory/create/Button.js");

            includeJS("../assets/components/inventory/create/paramsPanel/AttackInfo.js");
            includeJS("../assets/components/inventory/create/paramsPanel/DefendInfo.js");
            includeJS("../assets/components/inventory/create/paramsPanel/NavInfo.js");

            includeJS("../assets/components/inventory/selected/SelectInventoryItem.js");
            includeJS("../assets/components/inventory/selected/DeselectInventoryItem.js");
            includeJS("../assets/components/inventory/selected/SelectEquip.js");

            includeJS("../assets/components/inventory/filling/Filling.js");
            includeJS("../assets/components/inventory/filling/ConstructorTable.js");
            includeJS("../assets/components/inventory/filling/InventoryTable.js");
            includeJS("../assets/components/inventory/filling/SquadTable.js");
            includeJS("../assets/components/inventory/filling/MarkConstructorsCell.js");
            includeJS("../assets/components/inventory/filling/DeactivateCell.js");
            includeJS("../assets/components/inventory/filling/UnitPanel.js");
            includeJS("../assets/components/inventory/filling/HealBar.js");
            includeJS("../assets/components/inventory/filling/UpdateWeaponIcon.js");
            includeJS("../assets/components/inventory/filling/Storage.js");

            includeJS("../assets/components/inventory/tip/ClickTip.js");
            includeJS("../assets/components/inventory/tip/SelectItem.js");
            includeJS("../assets/components/inventory/tip/CreatePlaceBoxDialog.js");

            includeJS("../assets/components/inventory/set/SetAmmo.js");
            includeJS("../assets/components/inventory/set/SetBody.js");
            includeJS("../assets/components/inventory/set/SetEquip.js");
            includeJS("../assets/components/inventory/set/SetWeapon.js");
            includeJS("../assets/components/inventory/set/SetThorium.js");

            includeJS("../assets/components/inventory/repair/CreateRepairMenu.js");
            includeJS("../assets/components/inventory/repair/InventoryRepair.js");
            includeJS("../assets/components/inventory/repair/EquipsRepair.js");

            includeJS("../assets/components/inventory/inventory/RecycleItems.js");
            includeJS("../assets/components/inventory/inventory/SortingItems.js");
            includeJS("../assets/components/inventory/inventory/ThrowItems.js");
            includeJS("../assets/components/inventory/inventory/checkConfirmMenu.js");
            includeJS("../assets/components/inventory/inventory/SelectItems.js");
            includeJS("../assets/components/inventory/inventory/BlockInterface.js");

            includeCSS("../assets/components/inventory/css/constructorMS.css");
            includeCSS("../assets/components/inventory/css/constructorUnit.css");
            includeCSS("../assets/components/inventory/css/equipMSBox.css");
            includeCSS("../assets/components/inventory/css/equipUnitBox.css");
            includeCSS("../assets/components/inventory/css/inventoryCells.css");
            includeCSS("../assets/components/inventory/css/tip.css");
            includeCSS("../assets/components/inventory/css/repair.css");
            includeCSS("../assets/components/inventory/css/inventory.css");
            includeCSS("../assets/components/inventory/css/userStatus.css");
            includeCSS("../assets/components/inventory/css/weaponType.css");
            includeCSS("../assets/components/inventory/css/storage.css");
            includeCSS("../assets/components/inventory/css/marketDialog.css");
            includeCSS("../assets/components/inventory/css/passBlock.css");
        }
        return resolve();
    });
    //todo чето я хз, промис не работает
    promise.then(
        () => {
            initInventory = true;
            setTimeout(function () {
                CreateInventoryMenu(closeFunc, option);
                ConnectInventory();
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