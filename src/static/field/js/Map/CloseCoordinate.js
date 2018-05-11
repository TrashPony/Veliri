function CloseCoordinates(coordinates) {
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

function CloseCoordinate(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    GameMap.OneLayerMap[x][y].sprite.tint = 0x757575;
}