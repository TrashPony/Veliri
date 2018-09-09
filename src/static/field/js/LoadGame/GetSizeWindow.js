function GetWidth(hexagonWidth, gameMap) { //получаем ширину окна игры
    let width;

    if (window.innerWidth < hexagonWidth * gameMap.QSize) {
        width = window.innerWidth;
    } else {
        width = hexagonWidth * gameMap.QSize
    }

    return width
}

function GetHeight(hexagonHeight, gameMap) { //получаем высоту окна игры
    let height;

    if (window.innerHeight < hexagonHeight * gameMap.RSize) {
        height = window.innerHeight;
    } else {
        height = hexagonHeight * gameMap.RSize;
    }

    return height - 35
}