function FlightBullet() {
    for(let i in game.bulletLayer.children) {
        let bullet = game.bulletLayer.children[i];
        if (bullet.typeBullet === "rocket") {
            FlightRockets(bullet)
        }
    }
}