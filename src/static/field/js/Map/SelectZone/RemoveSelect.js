function RemoveSelect() {
    RemoveSelectLine();
    RemoveSelectCoordinate();
    RemoveTargetLine();
    TipEquipOff();
    RemoveUnitMarks();

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
        let lineSprite = game.SelectLayer.children.shift();
        lineSprite.destroy();
    }
}



