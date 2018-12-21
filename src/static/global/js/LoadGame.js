let game;

function LoadGame(jsonData) {
    game = CreateGame(jsonData.map);
    game.typeService = "global";

    setTimeout(function () { // todo костыль связаной с прогрузкой карты )
        CreateSquad(jsonData.squad);
        game.input.onDown.add(initMove, game);
        CreateBase(jsonData.bases)
    }, 1500);
}

function CreateBase(bases) {
    for (let i in bases) {
        if (bases.hasOwnProperty(i)) {
            if (game.map.OneLayerMap.hasOwnProperty(bases[i].q) && game.map.OneLayerMap.hasOwnProperty(bases[i].r)) {
                let coordinate = game.map.OneLayerMap[bases[i].q][bases[i].r];

                coordinate.objectSprite.inputEnabled = true;
                coordinate.objectSprite.input.pixelPerfectOver = true;
                coordinate.objectSprite.input.pixelPerfectClick = true;

                coordinate.objectSprite.events.onInputDown.add(function () {
                    game.squad.toBase = {
                        baseID: bases[i].id,
                        into: true,
                        x: coordinate.sprite.world.x,
                        y: coordinate.sprite.world.y
                    }
                });
            }
        }
    }
}