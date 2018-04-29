function moveTip(e) {

    var tipWeapon = document.getElementById("tipWeapon").style;
    var tipChassis = document.getElementById("tipChassis").style;
    var tipTower = document.getElementById("tipTower").style;
    var tipBody = document.getElementById("tipBody").style;
    var tipRadar = document.getElementById("tipRadar").style;


    var w = 250; // Ширина слоя
    var x = e.pageX; // Координата X курсора
    var y = e.pageY; // Координата Y курсора

    if ((x + w + 10) < document.body.clientWidth) {
        // Показывать слой справа от курсора
        tipChassis.left = x + 'px';
        tipWeapon.left = x + 'px';
        tipTower.left = x + 'px';
        tipRadar.left = x + 'px';
        tipBody.left = x + 'px';
    } else {
        // Показывать слой слева от курсора
        tipChassis.left = x - w + 'px';
        tipWeapon.left = x - w + 'px';
        tipTower.left = x - w + 'px';
        tipRadar.left = x - w + 'px';
        tipBody.left = x - w + 'px';
    }
    // Положение от верхнего края окна браузера
    tipChassis.top = y + 20 + 'px';
    tipWeapon.top = y + 20 + 'px';
    tipTower.top = y + 20 + 'px';
    tipRadar.top = y + 20 + 'px';
    tipBody.top = y + 20 + 'px';
}