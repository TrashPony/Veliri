function GetWidth(tileWidth, gameMap) { //получаем ширину окна игры
    var width;

    if (window.innerWidth < tileWidth * gameMap.XSize) {
        width = window.innerWidth;
    } else {
        width = tileWidth * gameMap.XSize
    }

    return width
}

function GetHeight(tileWidth, gameMap) { //получаем высоту окна игры
    var height;

    if (window.innerHeight < tileWidth * gameMap.YSize) {
        height = window.innerHeight;
    } else {
        height = tileWidth * gameMap.YSize;
    }

    return height
}