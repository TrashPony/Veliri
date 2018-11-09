/*

 tank 360 0вой градус орудия
 97 горизонталь
 40 вертикаль

*/

function PositionAttachSprite(angle, a) {
    // взятие координат угла элипса на изометричной окружности по радиусу
    let b = a / 2;

    let psi = angle * Math.PI / 180.0;
    let fi = Math.atan2(a * Math.sin(psi), b * Math.cos(psi));
    let x = a * Math.cos(fi);
    let y = b * Math.sin(fi);

    return {x: x, y: y};
}