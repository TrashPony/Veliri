var typeUnit;
var phase;
var move = null;
var target = null;
var field;

function ConnectField() {
    field = new WebSocket("ws://" + window.location.host + "/wsField");
    console.log("Websocket field - status: " + field.readyState);

    field.onopen = function() {
        console.log("CONNECTION field opened..." + this.readyState);
        InitGame();
    };

    field.onmessage = function(msg) {
        ReadResponse(msg.data);
    };

    field.onerror = function(msg) {
        console.log("Error field occured sending..." + msg.data);
    };

    field.onclose = function(msg) {
        // 1006 - ошибка при выключение сервера или отказа, 1001 - F5
        console.log("Disconnected field - status " + this.readyState);
        if (msg.code !== 1001) {
            location.href = "../../../login";
        }
    };
}