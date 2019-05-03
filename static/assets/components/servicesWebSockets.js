let inventorySocket;
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
        if (JSON.parse(msg.data).event === "UpdateStorage") {
            UpdateStorage(JSON.parse(msg.data).inventory);
        } else {
            FillingInventory(msg.data);
        }
    };

    inventorySocket.onerror = function (msg) {
        console.log("Error inventory occured sending..." + msg.data);
    };

    inventorySocket.onclose = function (msg) {
        console.log("Disconnected inventory - status " + this.readyState);
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