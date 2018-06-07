var stylePositionParams = {};

function moveMouse(e) {

    var w = 0; // Ширина слоя
    var x = e.pageX; // Координата X курсора
    var y = e.pageY; // Координата Y курсора

    if ((x + w + 100) < document.body.clientWidth) {
        // Показывать слой справа от курсора
        stylePositionParams.left = 20 + x + 'px';
    } else {
        // Показывать слой слева от курсора
        stylePositionParams.left = 20 + x - w + 'px';
    }
    // Положение от верхнего края окна браузера
    stylePositionParams.top = y + 'px';
}
