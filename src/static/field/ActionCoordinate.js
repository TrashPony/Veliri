var moveCell = [];

function DelMoveCoordinate() {
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

function DelUnit(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    cells[x + ":" + y].tint = 0x757575;
}

function OpenCoordinate(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    cells[x + ":" + y].tint = 0xffffff * 2;
}



function SelectCoordinateCreate(jsonMessage) {

}