let lobby;
let game = {typeService: "lobby"};

function ConnectLobby() {
    lobby = new WebSocket("ws://" + window.location.host + "/wsLobby");

    let date = new Date(0);
    document.cookie = "idGame=; path=/; expires=" + date.toUTCString();

    lobby.onopen = function (msg) {
    };

    lobby.onmessage = function (msg) {
        ReaderLobby(msg.data);
    };
    lobby.onerror = function (msg) {
        console.log("Error lobby occured sending..." + msg.data);
    };
    lobby.onclose = function (msg) {
        console.log("Disconnected lobby - status " + this.readyState);
        if (msg.code !== 1001) {
            location.href = "../../login";
        }
    };
}