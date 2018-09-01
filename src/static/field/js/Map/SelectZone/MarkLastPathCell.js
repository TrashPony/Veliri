function MarkLastPathCell(unit, cellState) {
    unit.lastCell = cellState;

    let q = cellState.q;
    let r = cellState.r;

    let mark = game.add.sprite(0, 0, 'MarkMoveLastCell'); // создаем метку
    mark.scale.set(.32);
    mark.alpha = 0.8;
    mark.z = 1;

    if (game.map.OneLayerMap[q][r].sprite) {
        game.map.OneLayerMap[q][r].sprite.addChild(mark);
    }
}

function DeleteMarkLastPathCell(cellState) {
    if (cellState) {
        let q = cellState.q;
        let r = cellState.r;
        let mark = game.map.OneLayerMap[q][r].sprite.getChildAt(0);
        mark.destroy();
    }
}