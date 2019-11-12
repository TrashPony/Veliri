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

    for (let i in game.objects) {
        let obj = game.objects[i];
        if (obj.geo_data && obj.geo_data.length > 0) {
            for (let i = 0; i < obj.geo_data.length; i++) {
                game.geoData.drawCircle(
                    obj.geo_data[i].x,
                    obj.geo_data[i].y,
                    obj.geo_data[i].radius * 2);
            }
        }
    }
}