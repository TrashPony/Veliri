function InitInventoryMenu() {

    includeJS("../assets/components/inventory/webSocket.js");
    includeJS("../assets/components/inventory/Create.js");
    includeJS("../assets/components/inventory/SelectInventoryItem.js");
    includeJS("../assets/components/inventory/Tip.js");

    includeJS("../assets/components/inventory/RemoveTip.js");
    includeJS("../assets/components/inventory/EquipSelect.js");

    includeJS("../assets/components/inventory/filling/Filling.js");
    includeJS("../assets/components/inventory/filling/ConstructorTable.js");
    includeJS("../assets/components/inventory/filling/InventoryTable.js");
    includeJS("../assets/components/inventory/filling/SquadTable.js");

    includeCSS("../assets/components/inventory/css/constructorMS.css");
    includeCSS("../assets/components/inventory/css/constructorUnit.css");
    includeCSS("../assets/components/inventory/css/equipMSBox.css");
    includeCSS("../assets/components/inventory/css/equipUnitBox.css");
    includeCSS("../assets/components/inventory/css/inventoryCells.css");
    includeCSS("../assets/components/inventory/css/tip.css");

    setTimeout(function () {
        ConnectInventory();
        CreateInventoryMenu();
    }, 400);
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