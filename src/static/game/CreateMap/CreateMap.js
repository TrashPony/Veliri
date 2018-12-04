function CreateMap() {

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
            game.mapPoints.push({x: startX, y: startY, q: q, r:r}); // x y - пиксельная координата положения, q r гексовая сеть
            startX += horizontalOffset;
        }
    }
}

