var typeUnit;
var phase;
var unitInfo;
var move = null;

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
        console.log("Disconnected - status " + this.readyState);
        location.href = "../../login"
    };

}

function createUnit(type) {
    typeUnit = type;
    sock.send(JSON.stringify({
        event: "SelectCoordinateCreate"
    }));
}

function reply_click(clicked_id) {
    var xy = clicked_id.split(":");

    var x = xy[0];
    var y = xy[1];

    if(phase === "move" && move !== null) {
        var unit = move.split(":");
        var unit_x = unit[0];
        var unit_y = unit[1];
        sock.send(JSON.stringify({
            event: "MoveUnit",
            x: Number(unit_x),
            y: Number(unit_y),
            to_x: Number(x),
            to_y: Number(y)
        }));
    } else {
        move = null;
    }

    if(phase === "Init" && typeUnit !== null && typeUnit !== undefined) {
        sock.send(JSON.stringify({
            event: "CreateUnit",
            type_unit: typeUnit,
            id_game: Number(idGame),
            x: Number(x),
            y: Number(y)
        }));
    } else {
        typeUnit = null;
    }
}

function mouse_over(unit_id) {
    var xy = unit_id.split(":");

    var x = xy[0];
    var y = xy[1];

    sock.send(JSON.stringify({
        event: "MouseOver",
        id_game: Number(idGame),
        x: Number(x),
        y: Number(y)
    }));
}

function mouse_out() {
    unitInfo = document.getElementById("unitInfo");
    unitInfo.innerHTML = "";
}

function SelectUnit(id) {
    if (move !== null) {
        DelMoveCell();
    }

    var xy = id.split(":");
    var x = xy[0];
    var y = xy[1];
    if (phase === "move") {
        move = id;
    }
    sock.send(JSON.stringify({
        event: "SelectUnit",
        x: Number(x),
        y: Number(y)
    }));
}

function Ready(){
    sock.send(JSON.stringify({
        event: "Ready",
        id_game: Number(idGame)
    }));
}