function InitInventoryMenu() {

    includeJS("../assets/components/inventory/webSocket.js");
    includeJS("../assets/components/inventory/Create.js");
    includeJS("../assets/components/inventory/filling.js");

    includeCSS("../assets/components/inventory/css/constructor.css");
    includeCSS("../assets/components/inventory/css/equipBox.css");
    includeCSS("../assets/components/inventory/css/inventoryCells.css");

    setTimeout(function () {
        ConnectInventory();
        CreateInventoryMenu();
    }, 400);

    if (inventorySocket && inventorySocket.readyState === 1) {
        inventorySocket.send(JSON.stringify({
            event: "openInventory"
        }));
    }
}

function includeJS(url) {
    var script = document.createElement('script');
    script.type = "text/javascript";
    script.charset = "charset='utf-8'";
    script.src = url;
    document.getElementsByTagName('head')[0].appendChild(script);
}

function includeCSS(url) {
    var css = document.createElement('link');
    css.type = "text/css";
    css.rel = "stylesheet";
    css.href = url;
    document.getElementsByTagName('head')[0].appendChild(css);
}