function GetWidth(tileWidth, gameMap) { //получаем ширину окна игры
    var width;

    if (window.innerWidth < tileWidth * gameMap.QSize) {
        width = window.innerWidth;
    } else {
        width = tileWidth * 150
    }

    return width
}

function GetHeight(tileWidth, gameMap) { //получаем высоту окна игры
    var height;

    if (window.innerHeight < tileWidth * gameMap.RSize) {
        height = window.innerHeight;
    } else {
        height = tileWidth * gameMap.RSize;
    }

    return height
}