function CreateGeoData(geoData) {
    if (game.geoData) {
        game.geoData.clear();
        game.geoData.destroy();
    }

    game.geoData = game.add.graphics(0, 0);
    game.geoDataLayer.add(game.geoData);
    game.geoData.beginFill(0xFF0000, 0.3);

    for (let i = 0; i < geoData.length; i++) {
        if (geoData[i]) {
            game.geoData.drawCircle(geoData[i].x, geoData[i].y, geoData[i].radius * 2);
        }
    }

    for (let x in game.map.OneLayerMap) {
        for (let y in game.map.OneLayerMap[x]) {
            let coordinate = game.map.OneLayerMap[x][y];

            if (coordinate.geo_data && coordinate.geo_data.length > 0) {
                for (let i = 0; i < coordinate.geo_data.length; i++) {
                    game.geoData.drawCircle(
                        coordinate.geo_data[i].x,
                        coordinate.geo_data[i].y,
                        coordinate.geo_data[i].radius * 2);
                }
            }
        }
    }
}