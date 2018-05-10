var moveCell = [];

function DelMoveCoordinate() {
    move = null;

    var buttonSkip = document.getElementById("SkipButton");
    buttonSkip.className = "button noActive";
    buttonSkip.onclick = null;

    for (var i = 0; i < moveCell.length; i++) {
        moveCell[i].tint = 0xffffff * 2;
        delete moveCell[i];
    }

    moveCell = [];
}

function DelUnit(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    GameMap.OneLayerMap[x][y].sprite.tint = 0x757575;
}

function OpenCoordinate(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    GameMap.OneLayerMap[x][y].sprite.tint = 0xffffff * 2;
}

function OpenCoordinates(coordinates) {
    while (coordinates.length > 0) {
        var coordinate = coordinates.shift();
        GameMap.OneLayerMap[coordinate.x][coordinate.y].sprite.tint = 0xffffff * 2;
    }
}

function DeleteCoordinates(coordinates) {
    while (coordinates.length > 0) {
        var coordinate = coordinates.shift();
        var id = coordinate.x + ":" + coordinate.y;

        GameMap.OneLayerMap[coordinate.x][coordinate.y].sprite.tint = 0x757575;

        if (units.hasOwnProperty(id)) {
            var unit = units[id];
            delete units[id];
            unit.destroy() // убиваем юнита
        }
    }
}


function SelectCoordinateCreate(jsonMessage) {

}