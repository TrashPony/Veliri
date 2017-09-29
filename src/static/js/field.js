var fieldUnit = 100;
var userName;
var idUnit;
var create = false;
var typeUnit;


function ConnectField() {
    sock = new WebSocket("ws://" + window.location.host + "/wsField");
    console.log("Websocket - status: " + sock.readyState);

    sock.onopen = function(msg) {
        console.log("CONNECTION opened..." + this.readyState);
    };
    sock.onmessage = function(msg) {
        console.log("message: " + msg.data);
        Response(msg.data);
    };
    sock.onerror = function(msg) {
        console.log("Error occured sending..." + msg.data);
    };
    sock.onclose = function(msg) {
        console.log("Disconnected - status " + this.readyState);
        //location.href = "http://642e0559eb9c.sn.mynetname.net:8080/login"
    }
}
/////////////////////////////////////////////////////////////////////Интерфейс////////////////////////////////////////////////
function SizeMap(params) {
    var div = document.getElementsByClassName("fieldUnit");
    if (params === 1) fieldUnit = fieldUnit + 30;
    if (fieldUnit > 45) {
        if (params === 2) fieldUnit = fieldUnit - 30;
    }

    for (var i = 0; 0 < div.length; i++) {
        if (params === 1) {
            div[i].style.height = fieldUnit + "px";
            div[i].style.width = fieldUnit + "px";
        }

        if (params === 2) {
            div[i].style.height = fieldUnit + "px";
            div[i].style.width = fieldUnit + "px";
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
/////////////////////////////////////////////////////////////////////RESPONSE////////////////////////////////////////////////
function Response(jsonMessage) {
    var event = JSON.parse(jsonMessage.body).event;
}
function CreateField() {
    Field(10,10)
}
/////////////////////////////////////////////////////////////////////CREATE FIELD/////////////////////////////////////////////////////////////////////
function Field(xSize,ySize) {
    var main = document.getElementById("main");
    main.style.boxShadow = "25px 25px 20px rgba(0,0,0,0.5)";

    for (var y = 0; y < ySize; y++) {
        for (var x = 0; x < xSize; x++) {
            var div = document.createElement('div');
                div.className = "fieldUnit";
                div.id = x + ":" + y;
                div.innerHTML = x + ":" + y;
                div.onclick = function () {
                    reply_click(this.id);
                };
                main.appendChild(div);
        }
        var nline = document.createElement('div');
        nline.className = "nline";
        nline.innerHTML = "";
        main.appendChild(nline);
    }
}
/////////////////////////////////////////////////////////////////////CREATE UNIT/////////////////////////////////////////////////////////////////////
function createUnit(type) {
    typeUnit = type;
    create = true;
}
/////////////////////////////////////////////////////////////////////
function reply_click(clicked_id) {

    var x = clicked_id[1];
    var y = clicked_id[3];

    if(create){
        var cell = document.getElementById(clicked_id);
        //sendCreateUnit(typeUnit, userName, x, y);
        if(typeUnit === "tank") cell.className = "fieldUnit tank";
        if(typeUnit === "scout") cell.className = "fieldUnit scout";
        if(typeUnit === "arta") cell.className = "fieldUnit arta";
        create = false;
        typeUnit = null;

    } else {
        //sendSelectEvent(x,y)
    }
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
