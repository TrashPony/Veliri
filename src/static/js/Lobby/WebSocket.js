function ConnectLobby() {
    sock = new WebSocket("ws://" + window.location.host + "/wsLobby");
    console.log("Websocket - status: " + sock.readyState);

    var date = new Date(0);
    document.cookie = "idGame=; path=/; expires=" + date.toUTCString();

    sock.onopen = function(msg) {
        console.log("CONNECTION opened..." + this.readyState);
        InitLobby();
        sendGameSelection();
        sendDontEndGamesList();
    };
    sock.onmessage = function(msg) {
        console.log("message: " + msg.data);
        ReaderLobby(msg.data);
    };
    sock.onerror = function(msg) {
        console.log("Error occured sending..." + msg.data);
    };
    sock.onclose = function(msg) {
        console.log("Disconnected - status " + this.readyState);
        if(!toField && msg.code !== 1001) {
            location.href = "../../login";
        }
    };
}