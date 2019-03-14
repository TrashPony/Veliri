function CreateAnomalies(anomalies) {
    if (!game.anomalies) {
        game.anomalies = game.add.graphics(0, 0);
        game.geoDataLayer.add(game.anomalies);
    }
    game.geoData.clear();

    for (let i = 0; i < anomalies.length; i++) {
        if (anomalies[i]) {
            game.geoData.beginFill(0x0098ff, 0.3);
            game.geoData.drawCircle(anomalies[i].x, anomalies[i].y, anomalies[i].radius * 2);
        }
    }
}