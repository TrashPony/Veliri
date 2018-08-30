function LoadOpenCoordinate() {
    console.log(game.user.watch);
    for (let x in game.user.watch) {
        if (game.user.watch.hasOwnProperty(x)) {
            for (let y in game.user.watch[x]) {
                if (game.user.watch[x].hasOwnProperty(y)) {
                    OpenCoordinate(game.user.watch[x][y].x, game.user.watch[x][y].z);
                }
            }
        }
    }
}