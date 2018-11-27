function UpdateMap(newMap, game) {
    
    game.floorLayer.forEach(function (c) { c.kill(); });
    game.floorObjectLayer.forEach(function (c) { c.kill(); });
    game.redactorButton.forEach(function (c) { c.kill(); });
    game.redactorMetaText.forEach(function (c) { c.kill(); });

    game.map = newMap;

    CreateMap();

    addButtons(game.map.OneLayerMap);
}