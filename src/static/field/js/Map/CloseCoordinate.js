function CloseCoordinates(coordinates) {
    while (coordinates.length > 0) {
        var coordinate = coordinates.shift();
        game.map.OneLayerMap[coordinate.x][coordinate.y].fogSprite.hide = false;

        var unit = GetGameUnitXY(coordinate.x, coordinate.y);

        if (unit) {
            delete game.units[unit.x][unit.y];
            unit.sprite.destroy();
            unit.shadow.destroy();
        }
    }
}

function CloseCoordinate(jsonMessage) {
    /*var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    GameMap.OneLayerMap[x][y].sprite.tint = 0x757575;*/
}