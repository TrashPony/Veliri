function LoadOpenCoordinate(watch) {
    for (var x in watch) {
        if (watch.hasOwnProperty(x)) {
            for (var y in watch[x]) {
                if (watch[x].hasOwnProperty(y)) {
                    OpenCoordinate(watch[x][y]);
                }
            }
        }
    }
}