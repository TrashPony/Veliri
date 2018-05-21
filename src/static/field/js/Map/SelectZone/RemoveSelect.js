function RemoveSelect() {
    RemoveSelectLine();
    RemoveSelectCoordinate();
}

function RemoveSelectLine() {
    while (game.SelectLineLayer.children.length > 0) {
        var lineSprite = game.SelectLineLayer.children.shift();
        lineSprite.destroy();
    }
}

function RemoveSelectCoordinate() {
    while (game.SelectLayer.children.length > 0) {
        var lineSprite = game.SelectLayer.children.shift();
        lineSprite.destroy();
    }
}

