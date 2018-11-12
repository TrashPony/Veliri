let marketSocket;

function ConnectMarket() {
    marketSocket = new WebSocket("ws://" + window.location.host + "/wsMarket");
    console.log("Websocket market - status: " + marketSocket.readyState);

    marketSocket.onopen = function() {
        console.log("Connection market opened..." + this.readyState);
        this.send(JSON.stringify({
            event: "openMarket"
        }));
    };

    marketSocket.onmessage = function(msg) {
        FillingInventory(msg.data);
    };

    marketSocket.onerror = function(msg) {
        console.log("Error market occured sending..." + msg.data);
    };

    marketSocket.onclose = function(msg) {
        console.log("Disconnected market - status " + this.readyState);
    };
}