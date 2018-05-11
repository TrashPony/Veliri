function LoadOpenCoordinate() {
    for (var x in game.user.watch) {
        if (game.user.watch.hasOwnProperty(x)) {
            for (var y in game.user.watch[x]) {
                if (game.user.watch[x].hasOwnProperty(y)) {
                    OpenCoordinate(game.user.watch[x][y]);
                }
            }
        }
    }
}