function CreateMap() {

    let hexagonWidth = 100;
    let hexagonHeight = 111;

    let verticalOffset = hexagonHeight * 3 / 4;
    let horizontalOffset = hexagonWidth;
    let startX;
    let startY;
    let startXInit = hexagonWidth / 2;
    let startYInit = hexagonHeight / 2;

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
                CreateObjects(coordinate, startX, startY);
            }

            if (coordinate.effects != null && coordinate.effects.length > 0) {
                MarkZoneEffect(coordinate, startX, startY)
            }
        }
    }
}

