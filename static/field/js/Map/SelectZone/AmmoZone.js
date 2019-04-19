function AmmoZone(coordinates) {
    //название спрайта для заполнения ammoCoordinate
    for (let q in coordinates) {
        for (let r in coordinates[q]) {
            let xy = GetXYCenterHex(q, r);
            let sprite = game.floorSelectLineLayer.create(xy.x, xy.y, 'ammoCoordinate');
            sprite.anchor.setTo(0.5);
            sprite.scale.set(0.1);
        }
    }
}