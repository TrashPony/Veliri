function mouse_over(sprite) {
    sprite.alpha = 0.5;
    toolTip()
}


function mouse_out(sprite) {
    sprite.alpha = 1;

    var floatTipStyle = document.getElementById("floatTip").style;
    floatTipStyle.display = "none"; // Прячем слой

    var targetCell = document.getElementsByClassName("aim mouse");
    while (targetCell.length > 0) {
        targetCell[0].remove();
    }
}

function moveTip(e) {
    var floatTipStyle = document.getElementById("floatTip").style;
    var w = 250; // Ширина слоя
    var x = e.pageX; // Координата X курсора
    var y = e.pageY; // Координата Y курсора

    if ((x + w + 10) < document.body.clientWidth) {
        // Показывать слой справа от курсора
        floatTipStyle.left = x + 'px';
    } else {
        // Показывать слой слева от курсора
        floatTipStyle.left = x - w + 'px';
    }
    // Положение от верхнего края окна браузера
    floatTipStyle.top = y + 20 + 'px';
}

function toolTip(jsonMessage) {
    var floatTipStyle = document.getElementById("floatTip").style;
    /*document.getElementById("tipUnit").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).type_unit + "</font>";
    document.getElementById("tipOwned").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).user_owned + "</font>";
    document.getElementById("tipHP").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).hp + "</font>";
    document.getElementById("tipAction").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).unit_action + "</font>";
    document.getElementById("tipTarget").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).target + "</font>";
    document.getElementById("tipDamage").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).damage + "</font>";
    document.getElementById("tipMove").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).move_speed + "</font>";
    document.getElementById("tipInit").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).init + "</font>";
    document.getElementById("tipRangeAttack").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).range_attack + "</font>";
    document.getElementById("tipRangeView").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).range_view + "</font>";
    document.getElementById("tipArea").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).area_attack + "</font>";
    document.getElementById("tipTypeAttack").innerHTML = "<font class='Value'>" + JSON.parse(jsonMessage).type_attack + "</font>";*/
    floatTipStyle.display = "block"; // Показываем слой
}