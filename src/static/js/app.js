try {
    var sock = new WebSocket("ws://" + window.location.host + "/ws");
    console.log("Websocket - status: " + sock.readyState);

    sock.onopen = function(msg) {
        console.log("CONNECTION opened..." + this.readyState);
    }

    sock.onmessage = function(msg) {
        console.log("message: " + msg.data);
    }

    sock.onerror = function(msg) {
        console.log("Error occured sending..." + msg.data);
    }

    sock.onclose = function(msg) {
        console.log("Disconnected - status " + this.readyState);
        location.href = "http://642e0559eb9c.sn.mynetname.net:8080/login"
    }

    function Send() {
        sock.send(JSON.stringify({
          target : "lobby",
          username: "test",
          message: "test"
        }));
    }

} catch(exception) {
    console.log(exception);
}