function OpenCoordinate(q, r) {
    let openFog = game.map.OneLayerMap[q][r].fogSprite; //.hide = true;
    game.map.OneLayerMap[q][r].open = true;
    if (openFog) {
        game.add.tween(openFog).to({alpha: 0}, 1500, Phaser.Easing.Linear.None, true, 0);
    }
}

function OpenCoordinates(coordinates) {
    while (coordinates.length > 0) {
        let coordinate = coordinates.shift();

        game.map.OneLayerMap[coordinate.q][coordinate.r].open = true;
        let openFog = game.map.OneLayerMap[coordinate.q][coordinate.r].fogSprite; //.hide = true;
        game.add.tween(openFog).to({alpha: 0}, 1500, Phaser.Easing.Linear.None, true, 0);
    }
}