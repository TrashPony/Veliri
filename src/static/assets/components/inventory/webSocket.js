var inventory;

function ConnectInventory() {
    inventory = new WebSocket("ws://" + window.location.host + "/wsInventory");
    console.log("Websocket inventory - status: " + inventory.readyState);

    inventory.onopen = function() {
        console.log("Connection inventory opened..." + this.readyState);
    };

    inventory.onmessage = function(msg) {
        console.log(msg);
        NewInventoryMessage(msg.data);
    };

    inventory.onerror = function(msg) {
        console.log("Error inventory occured sending..." + msg.data);
    };

    inventory.onclose = function(msg) {
        console.log("Disconnected inventory - status " + this.readyState);
    };
}

function NewInventoryMessage(data) {
    console.log(data);
}