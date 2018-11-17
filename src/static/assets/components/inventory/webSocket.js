let inventorySocket;
let storageSocket;


function ConnectInventory() {
    inventorySocket = new WebSocket("ws://" + window.location.host + "/wsInventory");
    console.log("Websocket inventory - status: " + inventorySocket.readyState);

    inventorySocket.onopen = function() {
        console.log("Connection inventory opened..." + this.readyState);
        this.send(JSON.stringify({
            event: "openInventory"
        }));
    };

    inventorySocket.onmessage = function(msg) {
        FillingInventory(msg.data);
    };

    inventorySocket.onerror = function(msg) {
        console.log("Error inventory occured sending..." + msg.data);
    };

    inventorySocket.onclose = function(msg) {
        console.log("Disconnected inventory - status " + this.readyState);
    };
}

function ConnectStorage() {
    storageSocket = new WebSocket("ws://" + window.location.host + "/wsStorage");
    console.log("Storage - status: " + storageSocket.readyState);

    storageSocket.onopen = function() {
        console.log("Connection storage opened..." + this.readyState);
        this.send(JSON.stringify({
            event: "openStorage"
        }));
    };

    storageSocket.onmessage = function(msg) {
        UpdateStorage(JSON.parse(msg.data));
    };

    storageSocket.onerror = function(msg) {
        console.log("Error storage occured sending..." + msg.data);
    };

    storageSocket.onclose = function(msg) {
        console.log("Disconnected storage - status " + this.readyState);
    };
}