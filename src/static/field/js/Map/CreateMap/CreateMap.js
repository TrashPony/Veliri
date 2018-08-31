function CreateMap() {

    let hexagonWidth = 80;
    let hexagonHeight = 100;

    let verticalOffset = hexagonHeight * 3 / 4;
    let horizontalOffset = hexagonWidth;
    let startX;
    let startY;
    let startXInit = hexagonWidth / 2;
    let startYInit = hexagonHeight / 2;

    for (let r = 0; r < game.map.QSize; r++) {

        if (r % 2 !== 0) {
            startX = 2 * startXInit;
        } else {
            startX = startXInit;
        }

        startY = startYInit + (r * verticalOffset);

        for (let q = 0; q < game.map.RSize; q++) {

            let coordinate = game.map.OneLayerMap[q][r];

            let floorSprite = game.floorLayer.create(startX, startY, "hexagon");
            let fogSprite = game.fogOfWar.create(startX, startY, 'FogOfWar');

            floorSprite.inputEnabled = true; // включаем ивенты на спрайт
            floorSprite.events.onInputOut.add(TipOff, floorSprite);
            floorSprite.events.onInputDown.add(RemoveSelect);
            floorSprite.z = 0;

            let label = game.add.text(20, 15, q + "," + r);
            floorSprite.addChild(label);

            coordinate.sprite = floorSprite;
            coordinate.fogSprite = fogSprite;

            startX += horizontalOffset;
        }
    }


    /*for (let q = 0; q < game.map.QSize/2; q++) {
        for (let r = 0; r < game.map.RSize; r++) {

            # преобразование кубических в осевые координаты
                q = x
                r = z
            # преобразование осевых в кубические координаты
                x = q
                z = r
                y = -x-z

                console.log("x: " + q + " z: " + r);
                console.log("y: " + Number(-q - r));



            //let coordinate = game.map.OneLayerMap[q][r];
            //CreateTerrain(coordinate, q, r);

            /*if (coordinate.level === 3) {
                let style = { font: "16px Arial", fill: "#ffa92b", align: "center" };
                game.add.text(x * game.tileSize + 30, y * game.tileSize + 30,
                    "x:" + x + " y:" + y + "\nl:" + coordinate.level, style);
            }

            if (coordinate.level === 4) {
                let style = { font: "16px Arial", fill: "#ff3f41", align: "center" };
                game.add.text(x * game.tileSize + 30, y * game.tileSize + 30,
                    "x:" + x + " y:" + y + "\nl:" + coordinate.level, style);
            }

            if (coordinate.texture_object !== "") {
                CreateObjects(coordinate);
            }

            if (coordinate.effects != null && coordinate.effects.length > 0) {
                MarkZoneEffect(coordinate)
            }*/
}

