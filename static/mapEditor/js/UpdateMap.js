function UpdateMap(newMap, game, bases) {
    removeSubMenus();

    game.floorLayer.forEach(function (c) {
        c.kill();
    });
    game.floorObjectLayer.forEach(function (c) {
        c.kill();
    });
    game.floorOverObjectLayer.forEach(function (c) {
        c.kill();
    });
    game.redactorButton.forEach(function (c) {
        c.kill();
    });
    game.redactorMetaText.forEach(function (c) {
        c.kill();
    });
    game.SelectLayer.forEach(function (c) {
        c.kill();
    });

    game.map = newMap;
    game.bases = bases;

    //game.world.setBounds(0, 0, (game.hexagonWidth + 5) * game.map.QSize, 185 * game.map.RSize/2); //размеры карты

    if (bases) {
        CreateLabelBase(bases);
    }
    CreateMiniMap();
    CreateMap();
    CreateGeoData(game.map.geo_data);
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

    if (document.getElementById("changeType")) {
        document.getElementById("changeType").remove()
    }

    if (document.getElementById("rotateBlock")) {
        document.getElementById("rotateBlock").remove()
    }
}