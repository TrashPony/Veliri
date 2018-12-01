function CreateMap() {

    let verticalOffset = game.hexagonHeight * 3 / 4;
    let horizontalOffset = game.hexagonWidth;
    let startX;
    let startY;
    let startXInit = game.hexagonWidth / 2;
    let startYInit = game.hexagonHeight / 2;

    for (let r = 0; r < game.map.RSize; r++) {

        if (r % 2 !== 0) {
            startX = 2 * startXInit;
        } else {
            startX = startXInit;
        }

        startY = startYInit + (r * verticalOffset);

        for (let q = 0; q < game.map.QSize; q++) {

            let coordinate = game.map.OneLayerMap[q][r];

            CreateTerrain(coordinate, startX, startY, q, r);

            startX += horizontalOffset;

            if (coordinate.texture_object !== "") {
                CreateObjects(coordinate);
            }

            if (coordinate.effects != null && coordinate.effects.length > 0) {
                MarkZoneEffect(coordinate)
            }
        }
    }
}

