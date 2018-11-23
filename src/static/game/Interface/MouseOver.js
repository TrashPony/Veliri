let stylePositionParams = {};

function moveMouse(e) {

    let w = 150; // Ширина слоя
    let x = e.pageX; // Координата X курсора
    let y = e.pageY; // Координата Y курсора

    if ((x + w + 20) < document.body.clientWidth) { // если слой выходит за пределый жкрана делает сноску в другой стороне
        // Показывать слой справа от курсора
        stylePositionParams.left = 20 + x;
    } else {
        // Показывать слой слева от курсора
        stylePositionParams.left = x - w - 20;
    }
    // Положение от верхнего края окна браузера
    stylePositionParams.top = y + 10;

    if (game && game.typeService === "battle") {
        updatePositionTipEquip();
    }
    updatePositionTipEffect();
}
