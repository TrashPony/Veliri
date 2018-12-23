function CreateBase(bases) {
    for (let i in bases) {
        if (bases.hasOwnProperty(i)) {
            if (game.map.OneLayerMap.hasOwnProperty(bases[i].q) && game.map.OneLayerMap.hasOwnProperty(bases[i].r)) {
                let coordinate = game.map.OneLayerMap[bases[i].q][bases[i].r];

                coordinate.objectSprite.inputEnabled = true;
                coordinate.objectSprite.input.pixelPerfectOver = true;
                coordinate.objectSprite.input.pixelPerfectClick = true;

                coordinate.objectSprite.events.onInputDown.add(function () {
                    if (game.input.activePointer.leftButton.isDown) {
                        game.squad.toBase = {
                            baseID: bases[i].id,
                            into: true,
                            x: coordinate.sprite.x,
                            y: coordinate.sprite.y
                        }
                    }
                });
            }
        }
    }
}