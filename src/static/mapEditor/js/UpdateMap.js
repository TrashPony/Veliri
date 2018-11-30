function UpdateMap(newMap, game) {
    removeSubMenus();
    game.floorLayer.forEach(function (c) { c.kill(); });
    game.floorObjectLayer.forEach(function (c) { c.kill(); });
    game.redactorButton.forEach(function (c) { c.kill(); });
    game.redactorMetaText.forEach(function (c) { c.kill(); });
    game.SelectLayer.forEach(function (c) { c.kill(); });

    game.map = newMap;

    CreateMap();

    addButtons(game.map.OneLayerMap);
}

function removeSubMenus() {
    if (document.getElementById("menuBlock")) {
        document.getElementById("menuBlock").remove();
    }

    if (document.getElementById("typeTip")) {
        document.getElementById("typeTip").remove()
    }

    if (document.getElementById("notification")) {
        document.getElementById("notification").remove()
    }
}