function InitMarketMenu() {
    let promise = new Promise((resolve) => {
        includeJS("../assets/components/market/webSocket.js");

        includeCSS("../assets/components/market/css/main.css");
        return resolve();
    });
    //todo чето я хз, промис не работает
    promise.then(
        () => {
            setTimeout(function () {
                ConnectMarket();
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