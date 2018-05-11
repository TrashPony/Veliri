function OpenCoordinate(coordinate) {
    GameMap.OneLayerMap[coordinate.x][coordinate.y].sprite.tint = 0xffffff * 2;
}

function OpenCoordinates(coordinates) {
    while (coordinates.length > 0) {
        var coordinate = coordinates.shift();
        GameMap.OneLayerMap[coordinate.x][coordinate.y].sprite.tint = 0xffffff * 2;
    }
}