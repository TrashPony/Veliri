function SendEventAddOrDelDetail() {
    // берем все выбраные части и шлем на свервер для обсчета статы
    var unitConstructor = document.getElementById("unitConstructor");

    lobby.send(JSON.stringify({
        event: "UnitConstructor",
        chassis: Number(checkDetailID(document.getElementById("chassisElement").detail)),
        weapon: Number(checkDetailID(document.getElementById("weaponElement").detail)),
        tower: Number(checkDetailID(document.getElementById("towerElement").detail)),
        body: Number(checkDetailID(document.getElementById("bodyElement").detail)),
        radar: Number(checkDetailID(document.getElementById("radarElement").detail)),
        slot: Number(unitConstructor.unitSlot)
    }));
}

function checkDetailID(detail) {
    if (detail) {
        return detail.id;
    } else {
        return 0;
    }
}

function SendEventSelectUnit() {

    var unitConstructor = document.getElementById("unitConstructor");

    var event;

    if (unitConstructor.unit === undefined || unitConstructor.unit === null) {
        event = "AddUnit"
    } else {
        event = "ReplaceUnit"
    }

    var chassisElement = checkDetailID(document.getElementById("chassisElement").detail);
    var weaponElement = checkDetailID(document.getElementById("weaponElement").detail);
    var towerElement = checkDetailID(document.getElementById("towerElement").detail);
    var bodyElement = checkDetailID(document.getElementById("bodyElement").detail);
    var radarElement = checkDetailID(document.getElementById("radarElement").detail);

    if (chassisElement === 0 && weaponElement === 0 && towerElement === 0 && bodyElement === 0 && radarElement === 0) {
        lobby.send(JSON.stringify({
            event: "RemoveUnit",
            slot: Number(unitConstructor.unitSlot)
        }));
    } else {
        lobby.send(JSON.stringify({
            event: event,
            chassis: Number(chassisElement),
            weapon: Number(weaponElement),
            tower: Number(towerElement),
            body: Number(bodyElement),
            radar: Number(radarElement),
            slot: Number(unitConstructor.unitSlot)
        }));
    }

    BackToLobby();
}