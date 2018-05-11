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