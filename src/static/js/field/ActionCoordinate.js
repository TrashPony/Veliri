var SelectCell = [];

function DelMoveCoordinate() {
    move = null;

    var buttonSkip = document.getElementById("SkipButton");
    buttonSkip.className = "button noActive";
    buttonSkip.onclick = null;

    for (var i = 0; i < SelectCell.length; i++) {
        if (SelectCell[i].type === "fieldUnit open") {
            OpenCoordinate(SelectCell[i].id)
        }
        if (SelectCell[i].type === "fieldUnit") {
            DelUnit(SelectCell[i].id)
        }
        delete SelectCell[i];
    }
    SelectCell = [];
}

function DelUnit(id) {
    var Cell = document.getElementById(id);
    Cell.className = "fieldUnit";
    Cell.innerHTML = id;
    Cell.style.color = "#FBFDFF";
    Cell.style.borderColor = "#404040";
    Cell.style.filter = "brightness(100%)";

    Cell.onclick = function () {
        SelectTarget(this.id);
    };
}

function OpenCoordinate(id) {
    var classUnit = "fieldUnit open";
    if (move != null) {
        classUnit = "fieldUnit move"
    }

    var Cell = document.getElementById(id);
    Cell.className = classUnit;
    Cell.id = id;
    Cell.innerHTML = id;
    Cell.style.color = "#fbfdff";
    Cell.style.borderColor = "#404040";
    Cell.style.filter = "brightness(100%)";

    Cell.onclick = function () {
        SelectTarget(this.id);
    };
    Cell.onmouseover = function () {
        mouse_over(this.id);
    };
    Cell.onmouseout = function () {
        mouse_out()
    };
}

function EmptyCoordinate(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var coor_id = x + ":" + y;
    var cell = document.getElementById(coor_id);
    if (cell) {
        cell.className = "fieldUnit open";
    }
}

function SelectCoordinateCreate(jsonMessage) {
    var x = JSON.parse(jsonMessage).x;
    var y = JSON.parse(jsonMessage).y;
    var coor_id = x + ":" + y;
    var cell = document.getElementById(coor_id);
    if (cell) {
        cell.className = "fieldUnit create";
    }
}