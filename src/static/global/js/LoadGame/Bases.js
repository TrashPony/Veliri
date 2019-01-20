function CreateBase(bases) {

    game.bases = bases;

    for (let i in bases) {
        if (bases.hasOwnProperty(i)) {
            if (game.map.OneLayerMap.hasOwnProperty(bases[i].q) && game.map.OneLayerMap.hasOwnProperty(bases[i].r)) {
                let coordinate = game.map.OneLayerMap[bases[i].q][bases[i].r];

                if (!coordinate.objectSprite) {
                    for (let j = 0; j < game.mapPoints.length; j++) {
                        if (game.mapPoints[j].q === bases[i].q && game.mapPoints[j].r === bases[i].r) {
                            CreateTerrain(coordinate, game.mapPoints[j].x, game.mapPoints[j].y, game.mapPoints[j].q, game.mapPoints[j].r);
                            CreateObjects(coordinate, game.mapPoints[j].x, game.mapPoints[j].y);
                            break
                        }
                    }
                }

                coordinate.objectSprite.inputEnabled = true;
                coordinate.objectSprite.input.pixelPerfectOver = true;
                coordinate.objectSprite.input.pixelPerfectClick = true;
                coordinate.base = true;

                coordinate.objectSprite.events.onInputDown.add(function () {
                    if (game.input.activePointer.leftButton.isDown) {
                        game.squad.toBase = {
                            baseID: bases[i].id,
                            into: true,
                            x: coordinate.objectSprite.x,
                            y: coordinate.objectSprite.y
                        }
                    }
                });
            }
        }
    }
}