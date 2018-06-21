function CreateMap() {
    console.log(game.map.OneLayerMap);

    for (var x = 0; x < game.map.XSize; x++) {
        for (var y = 0; y < game.map.YSize; y++) {
            var coordinate = game.map.OneLayerMap[x][y];
            // todo построение карты из существующих координат а не по размеру

            CreateTerrain(coordinate);

            if (coordinate.level === 3) {
                var style = { font: "16px Arial", fill: "#ffa92b", align: "center" };
                game.add.text(x * game.tileSize + 30, y * game.tileSize + 30,
                    "x:" + x + " y:" + y + "\nl:" + coordinate.level, style);
            }

            if (coordinate.level === 4) {
                var style = { font: "16px Arial", fill: "#ff3f41", align: "center" };
                game.add.text(x * game.tileSize + 30, y * game.tileSize + 30,
                    "x:" + x + " y:" + y + "\nl:" + coordinate.level, style);
            }


            if (coordinate.texture_object !== "") {
                CreateObjects(coordinate);
            }

            if (coordinate.effects != null && coordinate.effects.length > 0) {
                MarkZoneEffect(coordinate)
            }
        }
    }
}