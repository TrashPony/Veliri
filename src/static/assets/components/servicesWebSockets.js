let inventorySocket;
let storageSocket;
let marketSocket;

let webSocketInit = true;

function ConnectInventory() {
    inventorySocket = new WebSocket("ws://" + window.location.host + "/wsInventory");
    console.log("Websocket inventory - status: " + inventorySocket.readyState);

    inventorySocket.onopen = function () {
        console.log("Connection inventory opened..." + this.readyState);
        this.send(JSON.stringify({
            event: "openInventory"
        }));
    };

    inventorySocket.onmessage = function (msg) {
        FillingInventory(msg.data);
    };

    inventorySocket.onerror = function (msg) {
        console.log("Error inventory occured sending..." + msg.data);
    };

    inventorySocket.onclose = function (msg) {
        console.log("Disconnected inventory - status " + this.readyState);
    };
}

function ConnectStorage() {
    storageSocket = new WebSocket("ws://" + window.location.host + "/wsStorage");
    console.log("Storage - status: " + storageSocket.readyState);

    storageSocket.onopen = function () {
        console.log("Connection storage opened..." + this.readyState);
        this.send(JSON.stringify({
            event: "openStorage"
        }));
    };

    storageSocket.onmessage = function (msg) {
        UpdateStorage(JSON.parse(msg.data));
    };

    storageSocket.onerror = function (msg) {
        console.log("Error storage occured sending..." + msg.data);
    };

    storageSocket.onclose = function (msg) {
        console.log("Disconnected storage - status " + this.readyState);
    };
}

function ConnectMarket() {
    marketSocket = new WebSocket("ws://" + window.location.host + "/wsMarket");
    console.log("Market - status: " + marketSocket.readyState);

    marketSocket.onopen = function () {
        console.log("Connection market opened..." + this.readyState);
        this.send(JSON.stringify({
            event: "openMarket"
        }));
    };

    marketSocket.onmessage = function (msg) {
        if (JSON.parse(msg.data).error) {
            alert(JSON.parse(msg.data).error);
        } else if (document.getElementById("marketBox")) {
            if (JSON.parse(msg.data).event === "getMyOrders") {
                FillMyOrders(JSON.parse(msg.data).orders, JSON.parse(msg.data).base_name)
            } else {
                Filling(JSON.parse(msg.data));
            }
        }
    };

    marketSocket.onerror = function (msg) {
        console.log("Error market occured sending..." + msg.data);
    };

    marketSocket.onclose = function (msg) {
        console.log("Disconnected market - status " + this.readyState);
    };
}