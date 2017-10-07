var typeUnit;
var phase;
function ConnectField() {
    sock = new WebSocket("ws://" + window.location.host + "/wsField");
    console.log("Websocket - status: " + sock.readyState);

    sock.onopen = function(msg) {
        console.log("CONNECTION opened..." + this.readyState);
        InitGame();
    };
    sock.onmessage = function(msg) {
        console.log("message: " + msg.data);
        ReadResponse(msg.data);
    };
    sock.onerror = function(msg) {
        console.log("Error occured sending..." + msg.data);
    };
    sock.onclose = function(msg) {
        console.log("Disconnected - status " + this.readyState);
        location.href = "../../login"
    };

}

/////////////////////////////////////////////////////////////////////CREATE UNIT/////////////////////////////////////////////////////////////////////
function createUnit(type) {
    typeUnit = type;
}
/////////////////////////////////////////////////////////////////////
function reply_click(clicked_id) {
    var xy = clicked_id.split(":");

    var x = xy[0];
    var y = xy[1];

    if(phase === "Init" && typeUnit !== null) {
        sendCreateUnit(typeUnit, x, y);
    } else {
        typeUnit = null;
    }
}

/////////////////////////////////////////////////////////////////////GAME PROTOCOL/////////////////////////////////////////////////////////////////////

function sendCreateUnit(type, x, y){
    sock.send(JSON.stringify({
        event: "CreateUnit",
        type_unit: type,
        id_game: idGame,
        x: x,
        y: y
    }));
}
function sendReady(){
    sock.send(JSON.stringify({
        event: "Ready",
        id_game: idGame
    }));
}

function sendSelectEvent(x,y) {
    stompClient.send("/app/ControllerLobby", {}, JSON.stringify({'event': "SelectUnit",
                                                         'userName': "tost",
                                                                'x': x,
                                                                'y': y
                                                            }));
}

function sendMoveEvent(x,y) {
    stompClient.send("/app/ControllerLobby", {}, JSON.stringify({'event': "MoveUnit",
                                                         'userName': "Tost",
                                                           'idUnit': 999,
                                                                'x': x,
                                                                'y': y
                                                            }));
}

function sendTargetEvent(x,y) {
    stompClient.send("/app/ControllerLobby", {}, JSON.stringify({'event': "targetUnit",
                                                         'userName': "Tost",
                                                           'idUnit': 999,
                                                         'idTarget': 888,
                                                                'x': x,
                                                                'y': y
                                                            }));
}

function sleepFor( sleepDuration ){
    var now = new Date().getTime();
    while(new Date().getTime() < now + sleepDuration){ /* do nothing */ }
}
