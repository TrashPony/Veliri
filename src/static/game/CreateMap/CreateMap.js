function CreateMap() {
    return new Promise((resolve) => {
        let verticalOffset = game.hexagonHeight * 3 / 4;
        let horizontalOffset = game.hexagonWidth;
        let startX;
        let startY;
        let startXInit = game.hexagonWidth / 2;
        let startYInit = game.hexagonHeight / 2;

        game.mapPoints = []; // карта точек координат для динамического обнавления карты в методе Update

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

                if (coordinate.texture_object !== "") {
                    CreateObjects(coordinate);
                }

                if (coordinate.animate_sprite_sheets !== "") {
                    CreateAnimate(coordinate);
                }

                if (coordinate.dynamic_object) {
                    CreateDynamicObjects(coordinate.dynamic_object, q, r, true, coordinate);
                }

                if (coordinate.effects != null && coordinate.effects.length > 0) {
                    MarkZoneEffect(coordinate);
                }

                game.mapPoints.push({x: startX, y: startY, q: q, r: r}); // x y - пиксельная координата положения, q r гексовая сеть
                startX += horizontalOffset;
            }
        }
        resolve()
    });
}

