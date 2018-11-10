function CloseCoordinates(coordinates) {
    while (coordinates.length > 0) {
        let coordinate = coordinates.shift();

        let closeFog = game.map.OneLayerMap[coordinate.q][coordinate.r].fogSprite;
        game.add.tween(closeFog).to({alpha: 1}, 1500, Phaser.Easing.Linear.None, true, 0);

        let unit = GetGameUnitXY(coordinate.q, coordinate.r);

        if (unit) {
            UnitHide(unit);
        }
    }
}