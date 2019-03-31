function MarkGameZone(coordinates) {

    // удаляем старые и рисуем новые
    while (game.GameZone && game.GameZone.children.length > 0) {
        let lineSprite = game.GameZone.children.shift();
        lineSprite.destroy();
    }

    for (let q in coordinates) {
        for (let r in coordinates[q]) {
            let xy = GetXYCenterHex(q, r);
            MarkZone(xy, coordinates, q, r, 'GameZone', false, game.GameZone, "gameZone", game.GameZone, true);
        }
    }
}