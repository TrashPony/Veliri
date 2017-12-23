var typeUnit;
var phase;
var move = null;
var target = null;

function ConnectField() {
    sock = new WebSocket("ws://" + window.location.host + "/wsField");
    console.log("Websocket - status: " + sock.readyState);

    sock.onopen = function(msg) {
        console.log("CONNECTION opened..." + this.readyState);
        InitGame();
    };
    sock.onmessage = function(msg) {
        //console.log("message: " + msg.data);
        ReadResponse(msg.data);
    };
    sock.onerror = function(msg) {
        console.log("Error occured sending..." + msg.data);
    };
    sock.onclose = function(msg) {
        // 1006 ошибка при выключение сервера 1001 - F5
        console.log("Disconnected - status " + this.readyState);
        if (msg.code !== 1001) {
            location.href = "../../login"
        }
    };

}

function Ready(){
    if (move !== null) {
        DelMoveCell();
    }

    var targetCell = document.getElementsByClassName("aim");
    while (targetCell.length > 0) {
        targetCell[0].remove();
    }

    sock.send(JSON.stringify({
        event: "Ready",
        id_game: Number(idGame)
    }));
}