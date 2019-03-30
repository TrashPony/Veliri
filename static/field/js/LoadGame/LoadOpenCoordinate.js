function LoadOpenCoordinate() {
    // преобразуем матрицу координат в масив
    let watchCoordinates = [];
    for (let q in game.user.watch) {
        if (game.user.watch.hasOwnProperty(q)) {
            for (let r in game.user.watch[q]) {
                if (game.user.watch[q].hasOwnProperty(r)) {
                    watchCoordinates.push((game.user.watch[q][r]));
                }
            }
        }
    }
    OpenCoordinates(watchCoordinates);
    CreateAllFogOfWar();
}