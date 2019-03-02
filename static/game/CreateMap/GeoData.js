function CreateGeoData(geoData) {

    if (!game.geoData) {
        game.geoData = game.add.graphics(0, 0);
        game.geoDataLayer.add(game.geoData);
    }
    game.geoData.clear();

    for (let i = 0; i < geoData.length; i++) {
        if (geoData[i]) {
            game.geoData.beginFill(0xFF0000, 0.3);
            game.geoData.drawCircle(geoData[i].x, geoData[i].y, geoData[i].radius * 2);
        }
    }
}