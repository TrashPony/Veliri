function OpenCoordinate(coordinate) {
    game.map.OneLayerMap[coordinate.x][coordinate.y].fogSprite.hide = true;
}

function OpenCoordinates(coordinates) {
    while (coordinates.length > 0) {
        var coordinate = coordinates.shift();
        game.map.OneLayerMap[coordinate.x][coordinate.y].fogSprite.hide = true;
    }
}