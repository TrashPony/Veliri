function RemoveSelect() {
    RemoveSelectLine();
    RemoveSelectCoordinate();
    RemoveTargetLine();
    TipEquipOff();
    RemoveUnitMarks();
    RemoveSelectRangeCoordinate();

    if (document.getElementById("UnitSubMenu")) {
        document.getElementById("UnitSubMenu").remove()
    }

    if (document.getElementById("effectDetailZonePanel")) {
        document.getElementById("effectDetailZonePanel").remove();
    }
}

function RemoveSelectLine() {
    while (game.SelectLineLayer.children.length > 0) {
        let lineSprite = game.SelectLineLayer.children.shift();
        lineSprite.destroy();
    }
}

function RemoveTargetLine() {
    while (game.SelectTargetLineLayer.children.length > 0) {
        let lineSprite = game.SelectTargetLineLayer.children.shift();
        lineSprite.destroy();
    }
}

function RemoveSelectCoordinate() {
    while (game.SelectLayer.children.length > 0) {
        let sprite = game.SelectLayer.children.shift();
        sprite.destroy();
    }

    for (let q in game.map.OneLayerMap) {
        for (let r in game.map.OneLayerMap[q]) {
            game.map.OneLayerMap[q][r].targetSelectSprite = null;

            if (game.map.OneLayerMap[q][r].fastTargetSprite) {
                game.map.OneLayerMap[q][r].fastTargetSprite.destroy();
                game.map.OneLayerMap[q][r].fastTargetSprite = null;
            }
        }
    }
}

function RemoveSelectRangeCoordinate() {
    while (game.SelectRangeLayer.children.length > 0) {
        let sprite = game.SelectRangeLayer.children.shift();
        sprite.destroy();
    }
}



