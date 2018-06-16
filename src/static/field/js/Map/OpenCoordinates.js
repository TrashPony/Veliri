function OpenCoordinate(coordinate) {
    var openFog = game.map.OneLayerMap[coordinate.x][coordinate.y].fogSprite; //.hide = true;
    game.add.tween(openFog).to({alpha: 0}, 1500, Phaser.Easing.Linear.None, true, 0);
}

function OpenCoordinates(coordinates) {
    while (coordinates.length > 0) {
        var coordinate = coordinates.shift();
        var openFog = game.map.OneLayerMap[coordinate.x][coordinate.y].fogSprite; //.hide = true;
        game.add.tween(openFog).to({alpha: 0}, 1500, Phaser.Easing.Linear.None, true, 0);
    }
}