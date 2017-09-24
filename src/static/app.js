try {
    var sock = new WebSocket('ws://' + window.location.host + '/ws');
    console.log("Websocket - status: " + sock.readyState);
    sock.onopen = function(m) { 
        console.log("CONNECTION opened..." + this.readyState);
    }

    sock.onmessage = function(m) { 
        console.log("message: " + m.data);
    }

    sock.onerror = function(m) {
        console.log("Error occured sending..." + m.data);
    }

    sock.onclose = function(m) { 
        console.log("Disconnected - status " + this.readyState);
    }

    function Send() {
        sock.send(JSON.stringify({
          email: "test@test",
          username: "test",
          message: "test"
        }));
    }

} catch(exception) {
    console.log(exception);
}