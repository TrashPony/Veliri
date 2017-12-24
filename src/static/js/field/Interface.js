
var SizeUnit = 70;
var SizeText = 18;

function SizeMap(params) {

    var divs = document.getElementsByClassName("fieldUnit");

    if (params === 1) {
        SizeUnit = SizeUnit + 3;
        SizeText = SizeText + 0.6;
    }

    if (params === 2) {
        SizeUnit = SizeUnit - 3;
        SizeText = SizeText - 0.6;
    }

    for (var i in divs) {
        var div = document.getElementById(divs[i].id);
        if (div) {
            if (params === 1) {
                div.style.height = SizeUnit + "px";
                div.style.width = SizeUnit + "px";
                div.style.fontSize = SizeText + "px";
            }

            if (params === 2) {
                div.style.height = SizeUnit + "px";
                div.style.width = SizeUnit + "px";
                div.style.fontSize = SizeText + "px";
            }
        }
    }
}

function Wheel(e) {

    var delta = e.deltaY || e.detail || e.wheelDelta;
    // отмасштабируем при помощи CSS
    if (delta > 0) {
        SizeMap(1);
    } else {
        SizeMap(2);
    }
    // отменим прокрутку
    e.preventDefault();
}

function Rotate(params) {
    var div = document.getElementById('main');
    if(params === 0) {
        div.style.transition = "5s all";
        div.style.boxShadow = "25px 25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotate(0deg)";
    }
    if(params === 90) {
        div.style.transition = "5s all";
        div.style.boxShadow = "25px -25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotate(90deg)";
    }
    if(params === 180) {
        div.style.transition = "5s all";
        div.style.boxShadow = "-25px -25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotate(180deg)";
    }
    if(params === 270) {
        div.style.transition = "5s all";
        div.style.boxShadow = "-25px 25px 20px  rgba(0,0,0,0.5)";
        div.style.transform = "rotate(270deg)";
    }
}


