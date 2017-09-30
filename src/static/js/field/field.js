var SizeUnit = 100;


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
/////////////////////////////////////////////////////////////////////Интерфейс////////////////////////////////////////////////
function SizeMap(params) {
    var div = document.getElementsByClassName("fieldUnit");
    if (params === 1) SizeUnit = SizeUnit + 30;
    if (SizeUnit > 45) {
        if (params === 2) SizeUnit = SizeUnit - 30;
    }

    for (var i = 0; 0 < div.length; i++) {
        if (params === 1) {
            div[i].style.height = SizeUnit + "px";
            div[i].style.width = SizeUnit + "px";
        }

        if (params === 2) {
            div[i].style.height = SizeUnit + "px";
            div[i].style.width = SizeUnit + "px";
        }

    }
}
function Rotate(params) {
    var div = document.getElementById('main');
    if(params === 0) {
        div.style.transition = "5s all";
        div.style.boxShadow = "25px 25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotateX(13deg) translate(0px, -250px) rotate(0deg)";
    }
    if(params === 90) {
        div.style.transition = "5s all";
        div.style.boxShadow = "25px -25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotateX(13deg) translate(0px, -250px) rotate(90deg)";
    }
    if(params === 180) {
        div.style.transition = "5s all";
        div.style.boxShadow = "-25px -25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotateX(13deg) translate(0px, -250px) rotate(180deg)";
    }
    if(params === 270) {
        div.style.transition = "5s all";
        div.style.boxShadow = "-25px 25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotateX(13deg) translate(0px, -250px) rotate(270deg)";
    }
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

    var cell = document.getElementById(clicked_id);
    //sendCreateUnit(typeUnit, userName, x, y);
    if (typeUnit === "tank") cell.className = "fieldUnit tank";
    if (typeUnit === "scout") cell.className = "fieldUnit scout";
    if (typeUnit === "arta") cell.className = "fieldUnit arta";
    typeUnit = null;
}

/////////////////////////////////////////////////////////////////////GAME PROTOCOL/////////////////////////////////////////////////////////////////////

function sendCreateUnit(type, userName, x, y){
    stompClient.send("/app/ControllerLobby", {}, JSON.stringify({'event': "CreateUnit",
                                                         'typeUnit': type,
                                                         'userName': "tost",
                                                                'x': x,
                                                                'y': y
                                                          }));
}

function sendSelectEvent(x,y) {
    stompClient.send("/app/ControllerLobby", {}, JSON.stringify({'event': "SelectUnit",
                                                         'userName': "tost",
                                                           'idUnit': idUnit,
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
