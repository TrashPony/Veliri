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
}

function RemoveSelectRangeCoordinate() {
    while (game.SelectRangeLayer.children.length > 0) {
        let sprite = game.SelectRangeLayer.children.shift();
        sprite.destroy();
    }
}



