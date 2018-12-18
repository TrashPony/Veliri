let global;

function openSockets() {
    ConnectGlobal();
}

function ConnectGlobal() {
    global = new WebSocket("ws://" + window.location.host + "/wsGlobal");
    console.log("Websocket global - status: " + global.readyState);

    global.onopen = function() {
        console.log("CONNECTION global opened..." + this.readyState);
        global.send(JSON.stringify({
            event: "InitGame"
        }));
    };

    global.onmessage = function(msg) {
        ReadResponse(JSON.parse(msg.data));
    };

    global.onerror = function(msg) {
        console.log("Error global occured sending..." + msg.data);
    };

    global.onclose = function(msg) {
        // 1006 - ошибка при выключение сервера или отказа, 1001 - F5
        console.log("Disconnected global - status " + this.readyState);
        if (msg.code !== 1001) {
            location.href = "../../../login";
        }
    };
}