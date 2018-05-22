function SelectTargetCoordinateCreate(jsonMessage) {

    var targetCoordinates = JSON.parse(jsonMessage).targets;

    for (var x in targetCoordinates) {
        if (targetCoordinates.hasOwnProperty(x)) {
            for (var y in targetCoordinates[x]) {
                if (targetCoordinates[x].hasOwnProperty(y)) {
                    var cellSprite = game.map.OneLayerMap[targetCoordinates[x][y].x][targetCoordinates[x][y].y].sprite;

                    if (game.Phase === "move") {
                        MarkZone(cellSprite, targetCoordinates, x, y, 'Target', false, game.SelectTargetLineLayer);
                    }

                    if (game.Phase === "target") {
                        MarkZone(cellSprite, targetCoordinates, x, y, 'Target', true, game.SelectTargetLineLayer);
                        // todo
                    }
                }
            }
        }
    }
}