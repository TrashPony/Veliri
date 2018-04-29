function SendEventAddOrDelDetail() {
    // берем все выбраные части и шлем на свервер для обсчета статы

    lobby.send(JSON.stringify({
        event: "UnitConstructor",
        chassis: Number(checkDetailID(document.getElementById("chassisElement").detail)),
        weapon: Number(checkDetailID(document.getElementById("weaponElement").detail)),
        tower: Number(checkDetailID(document.getElementById("towerElement").detail)),
        body: Number(checkDetailID(document.getElementById("bodyElement").detail)),
        radar: Number(checkDetailID(document.getElementById("radarElement").detail))
    }));
}

function checkDetailID(detail) {
    if (detail){
        return detail.id;
    } else {
        return 0;
    }
}