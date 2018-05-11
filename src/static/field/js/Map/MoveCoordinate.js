var moveCell = [];

function RemoveSelectMoveCoordinate() {
    move = null;

    var buttonSkip = document.getElementById("SkipButton");
    buttonSkip.className = "button noActive";
    buttonSkip.onclick = null;

    for (var i = 0; i < moveCell.length; i++) {
        moveCell[i].tint = 0xffffff * 2;
        delete moveCell[i];
    }

    moveCell = [];
}

function SelectMoveCoordinate(x,y) {
    var cell = GameMap.OneLayerMap[x][y].sprite;
    cell.tint = 0xb5b5ff;

    if (cell) {
        moveCell.push(cell); // кладем выделеные ячейки в масив что бы потом удалить
    }
}