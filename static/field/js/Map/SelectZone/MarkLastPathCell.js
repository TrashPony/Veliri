function MarkLastPathCell(unit, cellState) {
    unit.lastCell = cellState;

    let xy = GetXYCenterHex(cellState.q, cellState.r);

    let mark = game.add.sprite(xy.x, xy.y, 'MarkMoveLastCell'); // создаем метку
    mark.anchor.setTo(0.5);
    mark.scale.set(.32);
    mark.alpha = 0.8;
    mark.z = 1;

    game.map.OneLayerMap[cellState.q][cellState.r].moveMark = mark;
}

function DeleteMarkLastPathCell(cellState) {
    if (cellState) {
        if (game.map.OneLayerMap[cellState.q][cellState.r].moveMark) {
            game.map.OneLayerMap[cellState.q][cellState.r].moveMark.destroy()
        }
    }
}