function CloseCoordinates(coordinates) {
    while (coordinates.length > 0) {
        var coordinate = coordinates.shift();

        var closeFog = game.map.OneLayerMap[coordinate.x][coordinate.y].fogSprite;
        game.add.tween(closeFog).to({alpha: 1}, 1500, Phaser.Easing.Linear.None, true, 0);

        var unit = GetGameUnitXY(coordinate.x, coordinate.y);

        if (unit) {
            unit.destroy = true;
        }
    }
}