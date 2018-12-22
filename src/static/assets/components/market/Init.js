let initMarket = false;

function InitMarketMenu(noMask) {
    let promise = new Promise((resolve) => {
        if (!initMarket) {
            if (typeof webSocketInit === 'undefined' || webSocketInit === null) {
                includeJS("../assets/components/servicesWebSockets.js");
            }

            includeJS("../assets/components/market/ToggleTab.js");

            includeJS("../assets/components/market/create/Create.js");
            includeJS("../assets/components/market/create/CreateTables.js");

            includeJS("../assets/components/market/Filling/Filling.js");
            includeJS("../assets/components/market/Filling/fillSellTable.js");
            includeJS("../assets/components/market/Filling/fillBuyTable.js");
            includeJS("../assets/components/market/Filling/myOrders.js");

            includeJS("../assets/components/market/Assortment/Assortment.js");
            includeJS("../assets/components/market/Assortment/fillEquip.js");
            includeJS("../assets/components/market/Assortment/fillCabs.js");
            includeJS("../assets/components/market/Assortment/fillWeapon.js");
            includeJS("../assets/components/market/Assortment/fillAmmo.js");
            includeJS("../assets/components/market/Assortment/fillRes.js");
            includeJS("../assets/components/market/Assortment/CreateFilter.js");

            includeCSS("../assets/components/market/css/main.css");
            includeCSS("../assets/components/market/css/leftBar.css");
            includeCSS("../assets/components/market/css/orderTables.css");
            includeCSS("../assets/components/market/css/marketTopMenu.css");
            includeCSS("../assets/components/market/css/marketRow.css");
        }
        resolve();
    });

    promise.then(
        () => {
            initMarket = true;
            setTimeout(function () {
                ConnectMarket();
                CreateMarketMenu(noMask);
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