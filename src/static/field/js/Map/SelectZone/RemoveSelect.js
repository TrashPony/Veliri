function RemoveSelect() {
    RemoveSelectLine();
    RemoveSelectCoordinate();
    RemoveTargetLine();

    if (document.getElementById("UnitSubMenu")) {
        document.getElementById("UnitSubMenu").remove()
    }
}

function RemoveSelectLine() {
    while (game.SelectLineLayer.children.length > 0) {
        var lineSprite = game.SelectLineLayer.children.shift();
        lineSprite.destroy();
    }
}

function RemoveTargetLine() {
    while (game.SelectTargetLineLayer.children.length > 0) {
        var lineSprite = game.SelectTargetLineLayer.children.shift();
        lineSprite.destroy();
    }
}

function RemoveSelectCoordinate() {
    while (game.SelectLayer.children.length > 0) {
        var lineSprite = game.SelectLayer.children.shift();
        lineSprite.destroy();
    }
}



