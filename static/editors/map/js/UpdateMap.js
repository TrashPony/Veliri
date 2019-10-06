function UpdateMap(newMap, game) {

    removeSubMenus();

    let clear = function (group) {
        while (group && group.children.length > 0) {
            let sprite = group.children.shift();
            sprite.destroy();
        }
    };

    game.map = newMap;

    clear(game.floorLayer);
    clear(game.SelectLayer);
    clear(game.floorObjectSelectLineLayer);
    clear(game.floorSelectLineLayer);
    clear(game.floorObjectLayer);
    clear(game.floorOverObjectLayer);
    clear(game.redactorButton);
    clear(game.redactorMetaText);
    clear(game.unitLayer);
    clear(game.SelectRangeLayer);
    clear(game.SelectLineLayer);
    clear(game.SelectTargetLineLayer);
    clear(game.effectsLayer);
    clear(game.artilleryBulletLayer);
    clear(game.weaponEffectsLayer);
    clear(game.flyObjectsLayer);
    clear(game.cloudsLayer);
    clear(game.fogOfWar);
    clear(game.geoDataLayer);
    clear(game.icon);

    CreateMiniMap();
    CreateMap();
}

function clear(group) {
    while (group && group.children.length > 0) {
        let sprite = group.children.shift();
        sprite.destroy();
    }
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