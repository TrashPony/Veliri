function MarkLastPathCell(unit, cellState) {
    unit.lastCell = cellState;

    var x = cellState.x;
    var y = cellState.y;

    var mark = game.add.sprite(0, 0, 'MarkMoveLastCell'); // создаем метку
    mark.scale.set(.32);
    mark.alpha = 0.8;
    mark.z = 1;

    if (game.map.OneLayerMap[x][y].sprite) {
        game.map.OneLayerMap[x][y].sprite.addChild(mark);
    }
}

function DeleteMarkLastPathCell(cellState) {
    if (cellState) {
        var x = cellState.x;
        var y = cellState.y;
        var mark = game.map.OneLayerMap[x][y].sprite.getChildAt(0);
        mark.destroy();
    }
}