function CreateAnomalies(anomalies) {
    if (!game.anomalies) {
        game.anomalies = game.add.graphics(0, 0);
        game.geoDataLayer.add(game.anomalies);
    }
    game.anomalies.clear();

    for (let i = 0; i < anomalies.length; i++) {
        if (anomalies[i]) {
            game.anomalies.beginFill(0x0098ff, 0.3);
            game.anomalies.drawCircle(anomalies[i].x, anomalies[i].y, anomalies[i].radius * 2);
        }
    }
}