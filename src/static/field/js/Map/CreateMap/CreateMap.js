function CreateMap() {


    for (let q = 0; q < game.map.QSize; q++) {
        for (let r = 0; r < game.map.RSize; r++) {
            if (game.map.QSize % 2 === 0 || q + 1 < game.map.QSize / 2 || r % 2 === 0) {
                /*
                # преобразование кубических в осевые координаты
                    q = x
                    r = z
                # преобразование осевых в кубические координаты
                    x = q
                    z = r
                    y = -x-z

                    console.log("x: " + q + " z: " + r);
                    console.log("y: " + Number(-q - r));
                */


                let coordinate = game.map.OneLayerMap[q][r];
                CreateTerrain(coordinate, q, r);

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
        }
    }
}