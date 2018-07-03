function InitInventoryMenu() {

    includeJS("../assets/components/inventory/webSocket.js");
    includeJS("../assets/components/inventory/Create.js");
    includeJS("../assets/components/inventory/filling.js");
    includeJS("../assets/components/inventory/SelectInventoryItem.js");
    includeJS("../assets/components/inventory/Tip.js");

    includeCSS("../assets/components/inventory/css/constructor.css");
    includeCSS("../assets/components/inventory/css/equipBox.css");
    includeCSS("../assets/components/inventory/css/inventoryCells.css");
    includeCSS("../assets/components/inventory/css/tip.css");

    setTimeout(function () {
        ConnectInventory();
        CreateInventoryMenu();
    }, 400);
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