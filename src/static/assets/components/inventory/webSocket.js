let inventorySocket;

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