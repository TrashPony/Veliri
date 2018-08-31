function LoadOpenCoordinate() {
    console.log(game.user.watch);
    for (let q in game.user.watch) {
        if (game.user.watch.hasOwnProperty(q)) {
            for (let r in game.user.watch[q]) {
                if (game.user.watch[q].hasOwnProperty(r)) {
                    OpenCoordinate(game.user.watch[q][r].q, game.user.watch[q][r].r);
                }
            }
        }
    }
}