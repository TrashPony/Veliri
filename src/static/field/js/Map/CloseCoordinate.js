function CloseCoordinates(coordinates) {
    while (coordinates.length > 0) {
        var coordinate = coordinates.shift();
        game.map.OneLayerMap[coordinate.x][coordinate.y].fogSprite.hide = false;

        var unit = GetGameUnitXY(coordinate.x, coordinate.y);

        if (unit) {
            unit.destroy = true;
        }
    }
}