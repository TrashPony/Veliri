function FlightBullet() {
    for (let i in game.artilleryBulletLayer.children) {
        let bullet = game.artilleryBulletLayer.children[i];
        if (bullet.typeBullet === "rocket") {
            FlightRockets(bullet)
        }
        if (bullet.typeBullet === "laser") {
            FlightLaser(bullet);
        }
    }
}