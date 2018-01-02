var webSocket;

function ConnectWS() {
    webSocket = new WebSocket("ws://" + window.location.host + "/wsGlobal");
    console.log("Websocket global - status: " + webSocket.readyState);

    webSocket.onopen = function() {
        console.log("CONNECTION global opened..." + this.readyState);
        Game();
    };

    webSocket.onerror = function(msg) {
        console.log("Error global occured sending..." + msg.data);
    };

    webSocket.onclose = function(msg) {
        // 1006 ошибка при выключение сервера или отказа, 1001 - F5
        console.log("Disconnected global - status " + this.readyState);
        if (msg.code !== 1001) {
            location.href = "../../login";
        }
    };
}