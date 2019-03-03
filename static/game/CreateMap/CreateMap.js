function CreateMap() {
    return new Promise((resolve) => {
        let verticalOffset = game.hexagonHeight * 3 / 4;
        let horizontalOffset = game.hexagonWidth;
        let startX;
        let startY;
        let startXInit = game.hexagonWidth / 2;
        let startYInit = game.hexagonHeight / 2;

        game.mapPoints = []; // карта точек координат для динамического обнавления карты в методе Update

        game.bmdTerrain.clear();

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
                    CreateObjects(coordinate, startX, startY);
                }

                if (coordinate.animate_sprite_sheets !== "") {
                    CreateAnimate(coordinate, startX, startY);
                }

                if (coordinate.dynamic_object) {
                    CreateDynamicObjects(coordinate.dynamic_object, q, r, true, coordinate);
                }

                if (coordinate.effects != null && coordinate.effects.length > 0) {
                    MarkZoneEffect(coordinate, startX, startY);
                }

                game.mapPoints.push({
                    x: startX,
                    y: startY,
                    q: q,
                    r: r,
                    textureOverFlore: coordinate.texture_over_flore
                }); // x y - пиксельная координата положения, q r гексовая сеть
                startX += horizontalOffset;
            }
        }

        CreateTexture();
        CreateBeams();
        resolve()
    });
}

function CreateBeams() {
    for (let i = 0; i < game.map.beams.length; i++) {
        CreateBeamLaser(
            game.map.beams[i].x_start,
            game.map.beams[i].y_start,
            game.map.beams[i].x_end,
            game.map.beams[i].y_end,
            game.map.beams[i].color,
        );
    }
}

function CreateTexture() {
    for (let i in game.mapPoints) {
        if (game.mapPoints[i].textureOverFlore !== '') {
            let bmd = game.make.bitmapData(512, 512);
            bmd.alphaMask(game.mapPoints[i].textureOverFlore, 'brush');
            game.bmdTerrain.draw(bmd, game.mapPoints[i].x - 256, game.mapPoints[i].y - 256);
            bmd.destroy();
        }
    }
}