function ReadInfoMouseOver(jsonMessage) {
    if (JSON.parse(jsonMessage).target !== "") {
        var xy = JSON.parse(jsonMessage).target.split(":");
        var x = xy[0];
        var y = xy[1];
        var idTarget = x + ":" + y;
        var targetCell = document.getElementById(idTarget);

        var div = document.createElement('div');
        div.className = "aim mouse";
        targetCell.appendChild(div);
    }
    toolTip(jsonMessage);
}

function mouse_over(cell) {
    var xy = cell.id.split(":");

    var x = xy[0];
    var y = xy[1];

    field.send(JSON.stringify({
        event: "MouseOver",
        id_game: Number(idGame),
        x: Number(x),
        y: Number(y)
    }));
}

function mouse_out() {
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
    if (jsonMessage) {
        document.getElementById("tipUnit").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).type_unit + "</spen>";
        document.getElementById("tipOwned").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).user_owned + "</spen>";
        document.getElementById("tipHP").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).hp + "</spen>";
        document.getElementById("tipAction").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).unit_action + "</spen>";
        document.getElementById("tipTarget").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).target + "</spen>";
        document.getElementById("tipDamage").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).damage + "</spen>";
        document.getElementById("tipMove").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).move_speed + "</spen>";
        document.getElementById("tipInit").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).init + "</spen>";
        document.getElementById("tipRangeAttack").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).range_attack + "</spen>";
        document.getElementById("tipRangeView").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).range_view + "</spen>";
        document.getElementById("tipArea").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).area_attack + "</spen>";
        document.getElementById("tipTypeAttack").innerHTML = "<spen class='Value'>" + JSON.parse(jsonMessage).type_attack + "</spen>";
        floatTipStyle.display = "block"; // Показываем слой
    }
}
