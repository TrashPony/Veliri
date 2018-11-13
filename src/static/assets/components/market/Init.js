function InitMarketMenu() {
    let promise = new Promise((resolve) => {
        includeJS("../assets/components/market/webSocket.js");
        includeJS("../assets/components/market/create/Create.js");
        includeJS("../assets/components/market/create/CreateTables.js");

        includeCSS("../assets/components/market/css/main.css");
        resolve();
    });

    promise.then(
        () => {
            setTimeout(function () {
                ConnectMarket();
                CreateMarketMenu();
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